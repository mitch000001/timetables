package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	runtimeDebug "runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest/auth"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2/google"
)

var funcMap = template.FuncMap{
	"printDate":         printDate,
	"printTimeframe":    printTimeframe,
	"printFiscalPeriod": printFiscalPeriod,
	"startsWith":        strings.HasPrefix,
}

var layoutPattern = filepath.Join(mustString(os.Getwd()), "templates", "layout.html.tmpl")
var partialTemplatePattern = filepath.Join(mustString(os.Getwd()), "templates", "_*.tmpl")
var layout = template.Must(template.ParseGlob(layoutPattern)).Funcs(funcMap)
var partialTmpl *template.Template = template.Must(layout.ParseGlob(partialTemplatePattern))

func printDate(date harvest.ShortDate) string {
	return date.Format("02.01.2006")
}

func printFiscalPeriod(fp *FiscalPeriod) string {
	return fmt.Sprintf("%s - %s", fp.StartDate.Format("02.01.2006"), fp.EndDate.Format("02.01.2006"))
}

func printTimeframe(tf harvest.Timeframe) string {
	return fmt.Sprintf("%s - %s", tf.StartDate.Format("02.01.2006"), tf.EndDate.Format("02.01.2006"))
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
	hostEnv := os.Getenv("HOST")
	if hostEnv != "" {
		host = hostEnv
	}
	port := os.Getenv("PORT")
	if port != "" {
		httpPort = port
	}
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
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
		RedirectURL:  host + "/google_oauth2redirect",
	}

	cache = &InMemoryCache{}
	workerQueue = newWorker(5)
	sessions = make(SessionManager)

	// TODO(mw): find a more readable way to compose handler
	http.HandleFunc("/", logHandler(htmlHandler(getHandler(authHandler(errorHandler(indexHandler()))))))
	http.HandleFunc("/login", logHandler(htmlHandler(loginHandler())))
	http.HandleFunc("/logout", logHandler(htmlHandler(getHandler(authHandler(errorHandler(logoutHandler()))))))
	http.HandleFunc("/google_login", logHandler(htmlHandler(googleLoginHandler(googleOauth2Config))))
	http.HandleFunc("/google_oauth2redirect", logHandler(htmlHandler(getHandler(googleRedirectHandler(googleOauth2Config)))))
	http.HandleFunc("/harvest_connect", logHandler(htmlHandler(getHandler(authHandler(errorHandler(harvestConnectHandler()))))))
	http.HandleFunc("/harvest_oauth", logHandler(htmlHandler(postHandler(authHandler(errorHandler(harvestOauthHandler()))))))
	http.HandleFunc("/harvest_oauth2redirect", logHandler(htmlHandler(getHandler(authHandler(errorHandler(harvestOauthRedirectHandler(harvestOauth2Config)))))))
	http.HandleFunc("/timeframes", logHandler(htmlHandler(authHandler(errorHandler(timeframesHandler())))))
	http.HandleFunc("/timeframes/new", logHandler(htmlHandler(getHandler(authHandler(errorHandler(timeframesNewHandler()))))))
	http.HandleFunc("/timeframes/show", logHandler(htmlHandler(getHandler(authHandler(errorHandler(harvestHandler(timeframeTableHandler())))))))
	http.HandleFunc("/plan_years", logHandler(htmlHandler(authHandler(errorHandler(planYearsHandler())))))
	http.HandleFunc("/plan_years/new", logHandler(htmlHandler(authHandler(errorHandler(planYearsNewHandler())))))
	http.HandleFunc("/plan_years/show/", logHandler(htmlHandler(authHandler(errorHandler(planYearsShowHandler())))))
	http.HandleFunc("/plan_items", logHandler(htmlHandler(authHandler(errorHandler(harvestHandler(planItemshandler()))))))
	http.HandleFunc("/plan_items/new", logHandler(htmlHandler(authHandler(errorHandler(harvestHandler(planItemNewHandler()))))))
	http.HandleFunc("/plan_items/show", logHandler(htmlHandler(getHandler(authHandler(errorHandler(planItemsShowHandler()))))))
	http.HandleFunc("/users/", logHandler(htmlHandler(authHandler(errorHandler(usersShowHandler())))))
	http.HandleFunc("/500", logHandler(htmlHandler(getHandler(authHandler(internalServerError())))))

	log.Printf("Listening on address %s\n", hostAddress)
	debug.Printf("Running in debug mode\n")
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}

var fiscalYear *FiscalYear

func indexHandler() authHandlerFunc {
	if fiscalYear == nil {
		fiscalYear = &FiscalYear{Year: time.Now().Year()}
	}
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		if r.URL.Path != "/" {
			s.AddError(fmt.Errorf("Die eingegebene Seite existiert nicht: '%s'", r.URL.Path))
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		timeframesHandler()(w, r, s)
	}
}

func timeframesNewHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		page := PageForSession(s)
		renderTemplate(w, "timeframes-new", page)
		return
	}
}

func timeframesHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		if r.Method == "GET" {
			page := PageForSession(s)
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
			renderTemplate(w, "timeframes", page)
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

func timeframeTableHandler() harvestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session, c *harvest.Harvest) {
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
		page := PageForSession(s)
		page.Set("Table", cache.Get(fmt.Sprintf("table:timeframe=%s", tf)))
		renderTemplate(w, "timeframe-table", page)
	}
}

var htmlReplacer = strings.NewReplacer("\n", "<br>", "\r", "<br>")

func internalServerError() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		page := PageForSession(s)
		page.Set("Request", r)
		page.Set("Stack", template.HTML(s.Stack))
		renderTemplate(w, "internal-server-error", page)
	}
}

func loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			renderTemplate(w, "login", &PageObject{})
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
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
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
		session.URL = r.URL
		fn(w, r, session)
	}
}

type harvestHandlerFunc func(http.ResponseWriter, *http.Request, *Session, *harvest.Harvest)

func harvestHandler(fn harvestHandlerFunc) authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		client, err := s.GetHarvestClient()
		if err != nil {
			debug.Printf("no client found: sessionId='%s', error=%T:%v\n", s.id, err, err)
			s.location = r.URL.String()
			http.Redirect(w, r, "/harvest_connect", http.StatusFound)
			return
		}
		if !s.User.Admin {
			s.AddError(fmt.Errorf("Der verknüpfte Account hat keine Administrationsrechte. Die Anwendung kann deshalb nur eingeschränkt genutzt werden."))
		}
		fn(w, r, s, client)
	}
}

func harvestConnectHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		page := PageForSession(s)
		renderTemplate(w, "harvest-connect", page)
		return
	}
}

func harvestOauthHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
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

		harvestOauth2Config := &HarvestOauth2Config{
			Config:    oauth2ConfigForEndpoint(harvestOauthEndpoint),
			Subdomain: subdomain,
		}
		s.User.SetHarvestOauthConfig(harvestOauth2Config)

		url := harvestOauth2Config.AuthCodeURL(s.id, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}

type HarvestOauth2Config struct {
	*oauth2.Config
	Subdomain string
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
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
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
		config := s.User.HarvestOauth2Config()
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
		session.User.SetHarvestToken(token)
		defer http.Redirect(w, r, session.location, http.StatusFound)

		client, err := session.GetHarvestClient()
		if err != nil {
			session.AddError(err)
			return
		}
		account, err := client.Account()
		if err != nil {
			session.AddError(err)
			return
		}
		session.User.SetHarvestAccount(account)
		return
	}
}

type authHandlerFunc func(http.ResponseWriter, *http.Request, *Session)

var sessions SessionManager

var contentTemplateString = `{{define "content"}}{{template "%s" .}}{{end}}`

func renderTemplate(w http.ResponseWriter, tmpl string, page *PageObject) {
	formattedTemplateString := fmt.Sprintf(contentTemplateString, tmpl)
	contentTemplate := template.Must(template.Must(layout.Clone()).Parse(formattedTemplateString))
	var buf bytes.Buffer
	var err error
	// TODO(mw): this is a really dirty hack to use this function with no pageObject
	if page == nil {
		err = contentTemplate.Execute(&buf, nil)
	} else {
		err = contentTemplate.Execute(&buf, page)
	}
	if err != nil {
		debug.Printf("Template error(%T): %v\n", err, err)
		http.Error(w, fmt.Sprintf("%T: %v\n", err, err), http.StatusInternalServerError)
	} else {
		io.Copy(w, &buf)
	}

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

func errorHandler(fn authHandlerFunc) authHandlerFunc {
	return func(w http.ResponseWriter, request *http.Request, s *Session) {
		defer func() {
			if r := recover(); r != nil {
				debug.Printf("Recovered: %+#v\n", r)
				if err, ok := r.(error); ok {
					s.AddError(err)
				} else {
					s.AddError(fmt.Errorf("%+#v", r))
				}
				stack := runtimeDebug.Stack()
				debug.Println(string(stack))
				s.Stack = htmlReplacer.Replace(string(stack))
				http.Redirect(w, request, "/500", http.StatusFound)
				return
			}
		}()
		fn(w, request, s)
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
