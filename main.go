package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest/auth"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2/google"
)

var funcMap = template.FuncMap{
	"printDate": printDate,
}

var layoutPattern = filepath.Join(mustString(os.Getwd()), "templates", "layout.html.tmpl")
var partialTemplatePattern = filepath.Join(mustString(os.Getwd()), "templates", "_*.tmpl")
var layout = template.Must(template.ParseGlob(layoutPattern)).Funcs(funcMap)
var partialTmpl *template.Template = template.Must(layout.ParseGlob(partialTemplatePattern))

func printDate(date harvest.ShortDate) string {
	return date.Format("02.01.2006")
}

func mustString(str string, err error) string {
	if err != nil {
		panic(err)
	}
	return str
}

var debug *log.Logger
var debugMode bool
var cache Cache
var googleOauth2Config *oauth2.Config
var harvestOauth2Config *oauth2.Config
var workerQueue *worker
var httpAddr string
var httpPort string
var host string

func init() {
	flag.BoolVar(&debugMode, "debug", false, "-debug")
	flag.StringVar(&httpAddr, "http.addr", "127.0.0.1", "-http.addr=localhost")
	flag.StringVar(&httpPort, "http.port", "4000", "-http.port=4000")
}

func main() {
	flag.Parse()
	debug = newDebugLogger(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	hostAddress := strings.TrimLeft(strings.TrimSuffix(httpAddr, ":"), "https://") + ":" + strings.TrimPrefix(httpPort, ":")
	host = "http://" + hostAddress
	harvestClientId := os.Getenv("HARVEST_CLIENTID")
	harvestClientSecret := os.Getenv("HARVEST_CLIENTSECRET")
	googleClientId := os.Getenv("GOOGLE_CLIENTID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENTSECRET")

	harvestOauth2Config = &oauth2.Config{
		ClientID:     harvestClientId,
		ClientSecret: harvestClientSecret,
		RedirectURL:  host + "/harvest_oauth2redirect",
	}

	googleOauth2Config = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"openid", "email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  host + "/google_oauth2redirect",
	}

	cache = &InMemoryCache{}
	workerQueue = newWorker(5)
	sessions = make(sessionMap)

	// TODO(mw): find a more readable way to compose handler
	http.HandleFunc("/", logHandler(htmlHandler(getHandler(authHandler(indexHandler())))))
	http.HandleFunc("/login", logHandler(htmlHandler(loginHandler())))
	http.HandleFunc("/logout", logHandler(htmlHandler(getHandler(authHandler(logoutHandler())))))
	http.HandleFunc("/google_login", logHandler(htmlHandler(googleLoginHandler(googleOauth2Config))))
	http.HandleFunc("/google_oauth2redirect", logHandler(htmlHandler(getHandler(googleRedirectHandler(googleOauth2Config)))))
	http.HandleFunc("/harvest_connect", logHandler(htmlHandler(getHandler(authHandler(harvestConnectHandler())))))
	http.HandleFunc("/harvest_oauth", logHandler(htmlHandler(postHandler(authHandler(harvestOauthHandler())))))
	http.HandleFunc("/harvest_oauth2redirect", logHandler(htmlHandler(getHandler(authHandler(harvestOauthRedirectHandler(harvestOauth2Config))))))
	http.HandleFunc("/timeframe", logHandler(htmlHandler(getHandler(authHandler(harvestHandler(timeframeHandler()))))))
	http.HandleFunc("/timeframes", logHandler(htmlHandler(authHandler(timeframesHandler()))))

	log.Printf("Listening on address %s\n", hostAddress)
	debug.Printf("Running in debug mode\n")
	log.Fatal(http.ListenAndServe(hostAddress, nil))
}

var indexTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "index" .}}{{end}}`))
var fiscalYear *FiscalYear

func indexHandler() authHandlerFunc {
	if fiscalYear == nil {
		fiscalYear = &FiscalYear{Year: time.Now().Year()}
	}
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		if r.URL.Path != "/" {
			s.AddError(fmt.Errorf("Die eingegebene Seite existiert nicht: '%s'", r.URL.Path))
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		page := pageForSession(s)
		page.Set("CurrentTimeframe", fiscalYear.CurrentFiscalPeriod())
		pastFiscalPeriods := fiscalYear.PastFiscalPeriods()
		pastTimeframes := make([]map[string]interface{}, len(pastFiscalPeriods))
		for i, fp := range pastFiscalPeriods {
			pastTimeframes[i] = map[string]interface{}{
				"Link":      template.URL(fp.Timeframe.ToQuery().Encode()),
				"StartDate": fp.StartDate,
				"EndDate":   fp.EndDate,
			}
		}
		page.Set("PastTimeframes", pastTimeframes)
		var buf bytes.Buffer
		err := indexTemplate.Execute(&buf, page)
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
		} else {
			io.Copy(w, &buf)
		}
	}
}

var timeframesTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "create-timeframes" .}}{{end}}`))

func timeframesHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		debug.Printf("Request: %+#v\n", r)
		if r.Method == "GET" {
			page := pageForSession(s)
			var buf bytes.Buffer
			err := timeframesTemplate.Execute(&buf, page)
			if err != nil {
				fmt.Fprintf(w, "%T: %v\n", err, err)
			} else {
				io.Copy(w, &buf)
			}
			return
		}
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				s.AddError(err)
				http.Redirect(w, r, "/timeframes", http.StatusFound)
				return
			}
			var multiErr multiError
			params := r.Form
			startDate, err := time.Parse("2006-01-02", params.Get("start-date"))
			if err != nil {
				multiErr.Add(err)
			}
			endDate, err := time.Parse("2006-01-02", params.Get("end-date"))
			if err != nil {
				multiErr.Add(err)
			}
			businessDays, err := strconv.Atoi(params.Get("business-days"))
			if err != nil {
				multiErr.Add(err)
			}
			fiscalPeriod := NewFiscalPeriod(startDate, endDate, businessDays)
			err = fiscalYear.Add(fiscalPeriod)
			if err != nil {
				multiErr.Add(err)
			}
			if len(multiErr) > 0 {
				for _, err := range multiErr {
					s.AddError(err)
				}
				http.Redirect(w, r, "/timeframes", http.StatusFound)
				return
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
}

var tableTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "table" .}}{{end}}`))

func timeframeHandler() harvestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session, c *harvest.Harvest) {
		params := r.URL.Query()
		tf, err := harvest.TimeframeFromQuery(params)
		if err != nil {
			debug.Printf("Error fetching timeframe from params: sessionId=%s\tparams=%+#v\terror=%T:%v\n", s.id, params, err, err)
			s.AddDebugError(err)
			// TODO(mw): What to do if the timeframe is not correct?
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		calcFn := func() {
			invalidateTableCacheForTimeframe(tf, c)
		}
		workerQueue.addJob(calcFn)
		page := pageForSession(s)
		page.Set("table", cache.Get(fmt.Sprintf("table:timeframe=%s", tf)))
		var buf bytes.Buffer
		err = tableTemplate.Execute(&buf, page)
		if err != nil {
			fmt.Fprintf(w, "%T: %v\n", err, err)
		} else {
			io.Copy(w, &buf)
		}
	}
}

var loginTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "login" .}}{{end}}`))

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
			http.SetCookie(w, newSessionCookie(s.id))
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
			debug.Printf("No cookie found: %v\n", err)
			r.Header.Set("X-Referer", r.URL.String())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if expired := cookie.Expires.After(time.Now()); expired {
			debug.Printf("Cookie expired: %+#v\n", cookie.Expires)
			r.Header.Set("X-Referer", r.URL.String())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		sessionId := cookie.Value
		session := sessions.Find(sessionId)
		if session == nil {
			debug.Printf("No session found for sessionId '%s'\n", sessionId)
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
		client, err := s.GetHarvestClient()
		if err != nil {
			debug.Printf("no client found: sessionId='%s', error=%T:%v\n", s.id, err, err)
			s.location = r.URL.String()
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		fn(w, r, s, client)
	}
}

var harvestConnectTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "harvest_connect" .}}{{end}}`))

func harvestConnectHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		page := pageForSession(s)
		var buf bytes.Buffer
		err := harvestConnectTemplate.Execute(&buf, page)
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
			s.AddError(err)
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		params := r.Form
		subdomain := params.Get("subdomain")
		if subdomain == "" {
			s.AddError(fmt.Errorf("Subdomain muss gefüllt sein"))
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
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

func harvestOauthRedirectHandler(harvestConfig *oauth2.Config) authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		params := r.URL.Query()
		state := params.Get("state")
		if state == "" {
			s.AddError(fmt.Errorf("State was not set in harvest oauth redirect"))
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		session := sessions.Find(state)
		if session == nil {
			s.AddError(fmt.Errorf("Sie müssen eingeloggt sein"))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		code := params.Get("code")
		if code == "" {
			s.AddError(fmt.Errorf("Die Antwort von Harvest war fehlerhaft."))
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		config := s.harvestOauth2Config
		if config == nil {
			s.AddError(fmt.Errorf("Keine oauth config für diese session gefunden."))
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			s.AddError(err)
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		session.harvestToken = token
		http.Redirect(w, r, session.location, http.StatusFound)
		return
	}
}

type authHandlerFunc func(http.ResponseWriter, *http.Request, *session)

func pageForSession(s *session) *pageObject {
	p := make(pageObject)
	p["session"] = s
	if s.idToken != nil {
		p["email"] = s.idToken.Email
	}
	sessErrors := s.GetErrors()
	if len(sessErrors) > 0 {
		p.AddErrors(sessErrors)
		s.ResetErrors()
	}
	return &p
}

type pageObject map[string]interface{}

func (p *pageObject) LoggedIn() bool {
	s, ok := (*p)["session"]
	if !ok {
		return false
	}
	return s.(*session).LoggedIn()
}

func (p *pageObject) Debug() bool {
	return debugMode
}

func (p *pageObject) CurrentUser() string {
	return (*p)["email"].(string)
}

func (p *pageObject) Errors() []error {
	errs, ok := (*p)["errors"]
	if !ok {
		return nil
	} else {
		return errs.([]error)
	}
}

func (p *pageObject) AddError(err error) {
	errs, ok := (*p)["errors"]
	var errors []error
	if !ok || errs == nil {
		errors = make([]error, 0)
	} else {
		errors = errs.([]error)
	}
	errors = append(errors, err)
	(*p)["errors"] = errors
}

func (p *pageObject) AddErrors(errs []error) {
	pErrs, ok := (*p)["errors"]
	var errors []error
	if !ok || errs == nil {
		errors = make([]error, 0)
	} else {
		errors = pErrs.([]error)
	}
	errors = append(errors, errs...)
	(*p)["errors"] = errors
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
	errors              []error
}

func (s *session) LoggedIn() bool {
	return s.idToken != nil
}

func (s *session) GetHarvestClient() (*harvest.Harvest, error) {
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

func (s *session) AddError(err error) {
	if s.errors == nil {
		s.errors = make([]error, 0)
	}
	s.errors = append(s.errors, err)
}

func (s *session) AddDebugError(err error) {
	if debugMode {
		s.AddError(err)
	}
}

func (s *session) GetErrors() []error {
	return s.errors
}

func (s *session) ResetErrors() {
	s.errors = make([]error, 0)
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

type debugWriter struct {
	debugMode *bool
	io.Writer
}

func (d *debugWriter) Write(p []byte) (int, error) {
	if *d.debugMode {
		return d.Writer.Write(p)
	}
	return 0, nil
}

func newDebugLogger(out io.Writer, prefix string, flag int) *log.Logger {
	debugOut := &debugWriter{Writer: out, debugMode: &debugMode}
	return log.New(debugOut, prefix, flag)
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

type worker struct {
	queue chan func()
}

func newWorker(size int) *worker {
	var queue chan func()
	if size <= 0 {
		queue = make(chan func())
	} else {
		queue = make(chan func(), size)
	}
	w := &worker{queue: queue}
	w.run()
	return w
}

func (w *worker) run() {
	go func() {
		for {
			select {
			case fn := <-w.queue:
				go fn()
			}
		}
	}()
}

func (w *worker) addJob(fn func()) {
	w.queue <- fn
}

func newUserHours(user *harvest.User, timeframe harvest.Timeframe, billable bool) *userHours {
	return &userHours{
		user:      user,
		timeframe: timeframe,
		billable:  billable,
	}
}

type userHours struct {
	user      *harvest.User
	timeframe harvest.Timeframe
	billable  bool
	hours     float64
}

func (u *userHours) getHours() float64 {
	return u.hours
}

type multiError []error

func (m *multiError) init() {
	if m == nil {
		*m = make(multiError, 0)
	}
}

func (m *multiError) Add(err error) {
	if m == nil {
		*m = make([]error, 0)
	}
	*m = append(*m, err)
}

func (m multiError) Error() string {
	var msgs []string
	for _, e := range m {
		msgs = append(msgs, fmt.Sprintf("%T %v", e, e))
	}
	return fmt.Sprintf("Errors: [%s]", strings.Join(msgs, ","))
}
