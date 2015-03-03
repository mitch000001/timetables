package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/go-harvest/harvest/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"
)

var funcMap = template.FuncMap{
	"printDate": printDate,
}

var rootTemplatePattern = filepath.Join(mustString(os.Getwd()), "index.html.tmpl")
var partialTemplatePattern = filepath.Join(mustString(os.Getwd()), "_*.tmpl")
var rootTemplate = template.Must(template.ParseGlob(rootTemplatePattern)).Funcs(funcMap)
var partialTmpl *template.Template = template.Must(rootTemplate.ParseGlob(partialTemplatePattern))

func printDate(date harvest.ShortDate) string {
	return date.Format("02.01.2006")
}

func mustString(str string, err error) string {
	if err != nil {
		panic(err)
	}
	return str
}

var cache Cache

func main() {
	subdomain := os.Getenv("HARVEST_SUBDOMAIN")
	username := os.Getenv("HARVEST_USERNAME")
	password := os.Getenv("HARVEST_PASSWORD")
	harvestClientId := os.Getenv("HARVEST_CLIENTID")
	harvestClientSecret := os.Getenv("HARVEST_CLIENTSECRET")
	googleClientId := os.Getenv("GOOGLE_CLIENTID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENTSECRET")
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	harvestOauthEndpoint := auth.NewOauth2EndpointForSubdomain(subdomain)

	harvestConfig := oauth2.Config{
		ClientID:     harvestClientId,
		ClientSecret: harvestClientSecret,
		Endpoint:     harvestOauthEndpoint,
		RedirectURL:  "http://localhost:" + port,
	}

	googleConfig := oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"openid", "email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://127.0.0.1:" + port + "/google_oauth2redirect",
	}

	clientProvider := auth.NewBasicAuthClientProvider(&auth.BasicAuthConfig{username, password})

	client, err := harvest.New(subdomain, clientProvider.Client)
	if err != nil {
		fmt.Printf("There was an error creating the client:\n")
		fmt.Printf("%T: %v\n", err, err)
		os.Exit(1)
	}
	cache = &InMemoryCache{}
	sessions = make(sessionMap)

	http.HandleFunc("/", logHandler(htmlHandler(getHandler(indexHandler(client)))))
	http.HandleFunc("/login", logHandler(htmlHandler(loginHandler())))
	http.HandleFunc("/google_login", logHandler(htmlHandler(googleLoginHandler(googleConfig))))
	http.HandleFunc("/google_oauth2redirect", logHandler(htmlHandler(googleRedirectHandler(googleConfig))))
	http.HandleFunc("/harvest_oauth", logHandler(htmlHandler(harvestOauthHandler(harvestConfig))))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var indexTemplate = template.Must(template.Must(rootTemplate.Clone()).Parse(`{{define "content"}}{{template "table" .}}{{end}}`))

func indexHandler(client *harvest.Harvest) http.HandlerFunc {
	startDate := harvest.Date(2015, 01, 01, time.Local)
	endDate := harvest.Date(2015, 01, 25, time.Local)
	rowTimeframe := harvest.Timeframe{StartDate: startDate, EndDate: endDate}
	calcFn := func() {
		invalidateTableCacheForTimeframe(rowTimeframe, client)
	}
	calcFn()
	// TODO(2015-02-22): remove brute force syncing with delta receiving via updated_since param
	go schedule(15*time.Second, calcFn)
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO(2015-02-20): reasonable caching
		var buf bytes.Buffer
		err := indexTemplate.Execute(&buf, cache.Get("table"))
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
		} else {
			io.Copy(w, &buf)
		}
	}
}

var loginTemplate = template.Must(template.Must(rootTemplate.Clone()).Parse(`{{define "content"}}{{template "login" .}}{{end}}`))

func loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var buf bytes.Buffer
			err := loginTemplate.Execute(&buf, nil)
			if err != nil {
				fmt.Fprintf(w, "%T: %v\n", err, err)
				return
			}
			io.Copy(w, &buf)
			return
		}
		if r.Method == "POST" {
			s := newSession()
			sessions.Add(s)
			http.Error(w, "NOT IMPLEMENTED", http.StatusInternalServerError)
			return
		}
	}
}

func authHandler(fn authHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("timetable")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		sessionId := cookie.Value
		session := sessions.Find(sessionId)
		if session == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		fn(w, r, session)
	}
}

func googleLoginHandler(config oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := newSession()
		sessions.Add(s)
		url := config.AuthCodeURL(s.id, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}

func googleRedirectHandler(config oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		state := params.Get("state")
		if state == "" {
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}
		session := sessions.Find(state)
		if session == nil {
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}
		code := params.Get("code")
		if code == "" {
			http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
			//http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%+#v\n", token)
		id := token.Extra("id_token")
		idToken, err := decode(id.(string))
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
			//http.Redirect(w, r, "/login", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%+#v\n", idToken)
		session.email = idToken.Email
	}
}

func decode(payload string) (*googleIdToken, error) {
	// decode returned id token to get expiry
	s := strings.Split(payload, ".")
	if len(s) < 2 {
		// TODO(jbd): Provide more context about the error.
		return nil, errors.New("jws: invalid token received")
	}
	decoded, err := base64Decode(s[1])
	if err != nil {
		return nil, err
	}
	c := &googleIdToken{}
	err = json.NewDecoder(bytes.NewBuffer(decoded)).Decode(c)
	return c, err
}
func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}

type googleIdToken struct {
	jws.ClaimSet
	AccessTokenHash     string `json:"at_hash"`
	EmailVerified       bool   `json:"email_verified"`
	AuthorizedPresenter string `json:"azp"`
	Email               string `json:"email"`
}

func harvestOauthHandler(config oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "NOT IMPLEMENTED", http.StatusInternalServerError)
	}
}

type authHandlerFunc func(http.ResponseWriter, *http.Request, *session)

var sessions sessionMap

type sessionMap map[string]*session

func (s *sessionMap) init() {
	if s == nil {
		*s = make(map[string]*session)
	}
}

func (sm *sessionMap) Add(s *session) {
	sm.init()
	(*sm)[s.id] = s
}

func (sm *sessionMap) Find(sessionId string) *session {
	return (*sm)[sessionId]
}

type session struct {
	email     string
	subdomain string
	id        string
}

func newSession() *session {
	b := make([]byte, 30)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	id := fmt.Sprintf("%x", sha256.Sum256(b))
	return &session{id: id}
}

func getHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
			return
		}
		fn(w, r)
	}
}

func postHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
			return
		}
		fn(w, r)
	}
}

func htmlHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fn(w, r)
	}
}

func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

func schedule(d time.Duration, fn func()) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			fn()
		}
	}
}

type table struct {
	Timeframe harvest.Timeframe
	Rows      []row
}

type row struct {
	User                   *harvest.User
	Hours                  float64
	Days                   float64
	CumulatedHours         float64
	CumulatedDays          float64
	BillableHours          float64
	BillableDays           float64
	BillingDegree          float64
	CumulatedBillableHours float64
	CumulatedBillableDays  float64
	CumulatedBillingDegree float64
}

func invalidateTableCacheForTimeframe(timeframe harvest.Timeframe, client *harvest.Harvest) {
	log.Printf("Invalidating table cache for timeframe %s\n", timeframe)
	start := time.Now()
	var t *table
	cacheValue := cache.Get("table")
	if cacheValue == nil {
		t = &table{}
	} else {
		t = cacheValue.(*table)
	}
	err := populateTable(t, timeframe, client)
	if err != nil {
		log.Printf("%T: %v\n", err, err)
	}
	cache.Store("table", t)
	log.Printf("Table cache invalidated, took %s", time.Since(start))
}

func populateTable(t *table, timeframe harvest.Timeframe, client *harvest.Harvest) error {
	var users []*harvest.User
	err := client.Users.All(&users, nil)
	if err != nil {
		return err
	}
	cumulationTimeframe := harvest.Timeframe{harvest.Date(2015, 01, 01, time.Local), timeframe.EndDate}
	var rows []row
	var multiErr multiError
	for _, user := range users {
		var hours float64
		var billableHours float64
		var cumulatedHours float64
		var cumulatedBillableHours float64
		hours, err = getHoursForUserAndTimeframe(user, timeframe, false, client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		billableHours, err = getHoursForUserAndTimeframe(user, timeframe, true, client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		// TODO: don't fetch all data since new years eve, use cached values
		cumulatedHours, err = getHoursForUserAndTimeframe(user, cumulationTimeframe, false, client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		// TODO: don't fetch all data since new years eve, use cached values
		cumulatedBillableHours, err = getHoursForUserAndTimeframe(user, cumulationTimeframe, true, client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		r := row{
			User:                   user,
			Hours:                  hours,
			Days:                   hours / 8,
			CumulatedHours:         cumulatedHours,
			CumulatedDays:          cumulatedHours / 8,
			BillableHours:          billableHours,
			BillableDays:           billableHours / 8,
			CumulatedBillableHours: cumulatedBillableHours,
			CumulatedBillableDays:  cumulatedBillableHours / 8,
			BillingDegree:          (billableHours / hours) * 100,
			CumulatedBillingDegree: (cumulatedBillableHours / cumulatedHours) * 100,
		}
		rows = append(rows, r)
	}
	if len(multiErr) != 0 {
		return multiErr
	}
	if t == nil {
		*t = table{}
	}
	t.Timeframe = timeframe
	t.Rows = rows
	return nil
}

func getHoursForUserAndTimeframe(user *harvest.User, timeframe harvest.Timeframe, billable bool, client *harvest.Harvest) (float64, error) {
	key := fmt.Sprintf("user=%d&timeframe=%s&billable=%t", user.Id(), timeframe, billable)
	dayEntries := cache.Get(key)
	var entries []*harvest.DayEntry
	var lastUpdate time.Time
	var cachedEntries []*harvest.DayEntry
	if dayEntries != nil {
		lastUpdate = time.Now()
		cachedEntries = dayEntries.([]*harvest.DayEntry)
	}
	params := harvest.Params{}
	params.ForTimeframe(timeframe).UpdatedSince(lastUpdate)
	if billable {
		params.OnlyBillable(billable)
	}
	err := client.Users.DayEntries(user).All(&entries, params.Values())
	if err != nil {
		return -1.0, err
	}
	newEntries := make(map[int]*harvest.DayEntry)
	for _, entry := range entries {
		newEntries[entry.ID] = entry
	}
	replacements := make(map[int]int)
	hours := 0.0
	for i, entry := range cachedEntries {
		if newEntry, ok := newEntries[entry.ID]; ok {
			replacements[entry.ID] = i
			hours += newEntry.Hours
		} else {
			hours += entry.Hours
		}
	}
	for id, i := range replacements {
		cachedEntries[i] = newEntries[id]
		delete(newEntries, id)
	}
	if len(newEntries) != 0 {
		for _, entry := range newEntries {
			cachedEntries = append(cachedEntries, entry)
			hours += entry.Hours
		}
	}
	cache.Store(key, cachedEntries)
	return hours, nil
}

type Cache interface {
	Get(key string) interface{}
	GetType(key string) reflect.Type
	Store(key string, data interface{}) error
}

type InMemoryCache struct {
	store map[string]interface{}
	mutex sync.RWMutex
}

func (i *InMemoryCache) init() {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.store == nil {
		i.store = make(map[string]interface{})
	}
}

func (i *InMemoryCache) Get(key string) interface{} {
	i.init()
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	return i.store[key]
}

func (i *InMemoryCache) GetType(key string) reflect.Type {
	i.init()
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	return reflect.TypeOf(i.store[key])
}

func (i *InMemoryCache) Store(key string, data interface{}) error {
	i.init()
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.store[key] = data
	return nil
}

type multiError []error

func (m *multiError) init() {
	if m == nil {
		*m = make(multiError, 0)
	}
}

func (m multiError) Add(err error) {
	m = append(m, err)
}

func (m multiError) Error() string {
	var msgs []string
	for _, e := range m {
		msgs = append(msgs, fmt.Sprintf("%T %v", e, e))
	}
	return fmt.Sprintf("Errors: [%s]", strings.Join(msgs, ","))
}
