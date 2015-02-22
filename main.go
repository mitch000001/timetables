package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

var indexTemplate = template.Must(template.Must(rootTemplate.Clone()).Parse(`{{define "content"}}{{template "table" .}}{{end}}`))

func printDate(date harvest.ShortDate) string {
	return date.Format("02.01.2006")
}

func mustString(str string, err error) string {
	if err != nil {
		panic(err)
	}
	return str
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
	http.HandleFunc("/", logHandler(htmlHandler(indexHandler(client))))
	log.Fatal(http.ListenAndServe(":4000", nil))
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

func indexHandler(client *harvest.Harvest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO(2015-02-20): reasonable caching
		startDate := harvest.Date(2015, 01, 01, time.Local)
		endDate := harvest.Date(2015, 01, 25, time.Local)
		rowTimeframe := harvest.Timeframe{StartDate: startDate, EndDate: endDate}
		var t table
		err := populateTable(&t, rowTimeframe, client)
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
			return
		}
		var buf bytes.Buffer
		err = indexTemplate.Execute(&buf, t)
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
		} else {
			io.Copy(w, &buf)
		}
	}
}

func populateTable(t *table, timeframe harvest.Timeframe, client *harvest.Harvest) error {
	var users []*harvest.User
	err := client.Users.All(&users, nil)
	if err != nil {
		return err
	}
	params := harvest.Params{}
	params.ForTimeframe(timeframe)
	billable := params.Clone().OnlyBillable(true)
	cumulated := harvest.Params{}
	cumulated.ForTimeframe(harvest.Timeframe{harvest.Date(2015, 01, 01, time.Local), timeframe.EndDate})
	cumulatedBillable := cumulated.Clone().OnlyBillable(true)
	var rows []row
	var multiErr multiError
	for _, user := range users {
		var entries []*harvest.DayEntry
		err = client.Users.DayEntries(user).All(&entries, url.Values(params))
		if err != nil {
			multiErr.Add(err)
			continue
		}
		hours := 0.0
		for _, entry := range entries {
			hours += entry.Hours
		}
		err = client.Users.DayEntries(user).All(&entries, url.Values(billable))
		if err != nil {
			multiErr.Add(err)
			continue
		}
		billableHours := 0.0
		for _, entry := range entries {
			billableHours += entry.Hours
		}
		// TODO: don't fetch all data since new years eve, use cached values
		err = client.Users.DayEntries(user).All(&entries, url.Values(cumulated))
		if err != nil {
			multiErr.Add(err)
			continue
		}
		cumulatedHours := 0.0
		for _, entry := range entries {
			cumulatedHours += entry.Hours
		}
		// TODO: don't fetch all data since new years eve, use cached values
		err = client.Users.DayEntries(user).All(&entries, url.Values(cumulatedBillable))
		if err != nil {
			multiErr.Add(err)
			continue
		}
		cumulatedBillableHours := 0.0
		for _, entry := range entries {
			cumulatedBillableHours += entry.Hours
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
	*t = table{
		Timeframe: timeframe,
		Rows:      rows,
	}
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
