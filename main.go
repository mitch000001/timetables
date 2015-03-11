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
var googleOauth2Config *oauth2.Config
var harvestOauth2Config *oauth2.Config

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

	harvestOauth2Config = &oauth2.Config{
		ClientID:     harvestClientId,
		ClientSecret: harvestClientSecret,
		RedirectURL:  "http://127.0.0.1:" + port + "/harvest_oauth2redirect",
	}

	googleOauth2Config = &oauth2.Config{
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

	http.HandleFunc("/", logHandler(htmlHandler(getHandler(authHandler(indexHandler(client))))))
	http.HandleFunc("/login", logHandler(htmlHandler(loginHandler())))
	http.HandleFunc("/logout", logHandler(htmlHandler(getHandler(authHandler(logoutHandler())))))
	http.HandleFunc("/google_login", logHandler(htmlHandler(googleLoginHandler(googleOauth2Config))))
	http.HandleFunc("/google_oauth2redirect", logHandler(htmlHandler(getHandler(googleRedirectHandler(googleOauth2Config)))))
	http.HandleFunc("/harvest_connect", logHandler(htmlHandler(getHandler(authHandler(harvestConnectHandler())))))
	http.HandleFunc("/harvest_oauth", logHandler(htmlHandler(postHandler(authHandler(harvestOauthHandler())))))
	http.HandleFunc("/harvest_oauth2redirect", logHandler(htmlHandler(getHandler(authHandler(harvestOauthRedirectHandler(harvestOauth2Config))))))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var indexTemplate = template.Must(template.Must(rootTemplate.Clone()).Parse(`{{define "content"}}{{template "table" .}}{{end}}`))

func indexHandler(client *harvest.Harvest) authHandlerFunc {
	startDate := harvest.Date(2015, 01, 01, time.Local)
	endDate := harvest.Date(2015, 01, 25, time.Local)
	rowTimeframe := harvest.Timeframe{StartDate: startDate, EndDate: endDate}
	calcFn := func() {
		invalidateTableCacheForTimeframe(rowTimeframe, client)
	}
	calcFn()
	// TODO(2015-02-22): remove brute force syncing with delta receiving via updated_since param
	go schedule(15*time.Second, calcFn)
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		page := pageForSession(s)
		page.Set("table", cache.Get("table"))
		var buf bytes.Buffer
		err := indexTemplate.Execute(&buf, page)
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
			http.SetCookie(w, &http.Cookie{Name: "timetable", Value: s.id, Expires: time.Now().Add(5 * 24 * time.Hour)})
			http.Error(w, "NOT IMPLEMENTED", http.StatusInternalServerError)
			return
		}
	}
}

func logoutHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		sessions.Remove(s)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func authHandler(fn authHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("timetable")
		if err != nil {
			r.Header.Set("X-Referer", r.URL.String())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		sessionId := cookie.Value
		session := sessions.Find(sessionId)
		if session == nil {
			r.Header.Set("X-Referer", r.URL.String())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		fn(w, r, session)
	}
}

type harvestHandlerFunc func(http.ResponseWriter, *http.Request, *session, *harvest.Harvest)

func harvestHandler(fn harvestHandlerFunc) authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		client, err := s.getHarvestClient()
		if err != nil {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		fn(w, r, s, client)
	}
}

func googleLoginHandler(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := newSession()
		s.location = r.Header.Get("X-Referer")
		sessions.Add(s)
		url := config.AuthCodeURL(s.id, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}

func googleRedirectHandler(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		state := params.Get("state")
		if state == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session := sessions.Find(state)
		if session == nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		code := params.Get("code")
		if code == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session.googleToken = token
		id := token.Extra("id_token")
		idToken, err := decode(id.(string))
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session.idToken = idToken
		http.SetCookie(w, &http.Cookie{Name: "timetable", Value: session.id, Expires: time.Now().Add(5 * 24 * time.Hour)})
		http.Redirect(w, r, session.location, http.StatusFound)
	}
}

var harvestConnectTemplate = template.Must(template.Must(rootTemplate.Clone()).Parse(`{{define "content"}}{{template "harvest_connect" .}}{{end}}`))

func harvestConnectHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		var buf bytes.Buffer
		err := harvestConnectTemplate.Execute(&buf, nil)
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
			return
		}
		io.Copy(w, &buf)
		return
	}
}

func harvestOauthHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		err := r.ParseForm()
		if err != nil {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		params := r.Form
		subdomain := params.Get("subdomain")
		if subdomain == "" {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		// TODO(mw): move into harvest package and extract as utility function
		harvestUrl := "https://" + subdomain + ".harvestapp.com"
		harvestOauthEndpoint := auth.NewOauth2EndpointForSubdomain(harvestUrl)

		s.harvestSubdomain = subdomain
		s.harvestOauth2Config = oauth2ConfigForEndpoint(harvestOauthEndpoint)

		url := s.harvestOauth2Config.AuthCodeURL(s.id, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}

func oauth2ConfigForEndpoint(endpoint oauth2.Endpoint) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     harvestOauth2Config.ClientID,
		ClientSecret: harvestOauth2Config.ClientSecret,
		RedirectURL:  harvestOauth2Config.RedirectURL,
		Endpoint:     endpoint,
	}

}

func harvestOauthRedirectHandler(config *oauth2.Config) authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		params := r.URL.Query()
		state := params.Get("state")
		if state == "" {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		session := sessions.Find(state)
		if session == nil {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		code := params.Get("code")
		if code == "" {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Redirect(w, r, "/harvest_connect", http.StatusTemporaryRedirect)
			return
		}
		session.harvestToken = token
		http.Redirect(w, r, session.location, http.StatusFound)
		return
	}
}

func decode(payload string) (*googleIdToken, error) {
	// decode returned id token to get expiry
	s := strings.Split(payload, ".")
	if len(s) < 2 {
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

type authHandlerFunc func(http.ResponseWriter, *http.Request, *session)

func pageForSession(s *session) *pageObject {
	p := make(pageObject)
	p["session"] = s
	if s.idToken != nil {
		p["email"] = s.idToken.Email
	}
	return &p
}

type pageObject map[string]interface{}

func (p *pageObject) LoggedIn() bool {
	s, ok := (*p)["session"]
	if !ok {
		return false
	}
	return s.(*session).loggedIn()
}

func (p *pageObject) CurrentUser() string {
	return (*p)["email"].(string)
}

func (p *pageObject) Set(key string, value interface{}) {
	(*p)[key] = value
}

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

func (sm *sessionMap) Remove(s *session) {
	delete(*sm, s.id)
}

type session struct {
	location            string
	googleToken         *oauth2.Token
	idToken             *googleIdToken
	harvestOauth2Config *oauth2.Config
	harvestSubdomain    string
	harvestToken        *oauth2.Token
	id                  string
}

func (s *session) loggedIn() bool {
	return s.idToken != nil
}

func (s *session) getHarvestClient() (*harvest.Harvest, error) {
	if s.harvestOauth2Config == nil {
		return nil, fmt.Errorf("Missing harvest oauth config")
	}
	if s.harvestToken == nil {
		return nil, fmt.Errorf("Missing harvest token")
	}
	if s.harvestSubdomain == "" {
		return nil, fmt.Errorf("Missing harvest subdomain")
	}
	// TODO(mw): validate that the token is valid and if not, exchange a new token!
	client, err := harvest.New(s.harvestSubdomain, func() harvest.HttpClient { return s.harvestOauth2Config.Client(oauth2.NoContext, s.harvestToken) })
	if err != nil {
		return nil, fmt.Errorf("Error while creating new harvest client: %T(%v)", err, err)
	}
	return client, nil
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
