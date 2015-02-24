package main

import (
	"bytes"
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

	clientProvider := auth.NewBasicAuthClientProvider(&auth.BasicAuthConfig{username, password})

	client, err := harvest.New(subdomain, clientProvider.Client)
	if err != nil {
		fmt.Printf("There was an error creating the client:\n")
		fmt.Printf("%T: %v\n", err, err)
		os.Exit(1)
	}
	cache = &InMemoryCache{}
	http.HandleFunc("/", logHandler(htmlHandler(indexHandler(client))))
	log.Fatal(http.ListenAndServe(":4000", nil))
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
			} else {
				io.Copy(w, &buf)
			}
		}
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
