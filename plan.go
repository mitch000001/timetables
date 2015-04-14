package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"time"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
)

type PlanYears map[int]*PlanYear

func (p *PlanYears) FindByYear(year int) *PlanYear {
	return (*p)[year]
}

type PlanYear struct {
	*FiscalYear
	PlanItems                 PlanItems
	PlanUsers                 PlanUsers
	AverageDaysOfIllness      float64
	AverageDaysOfChildrenCare float64
	DefaultVacationInterest   float64
}

type PlanItems []*PlanItem

func (p *PlanItems) FindByFiscalPeriod(fp *FiscalPeriod) *PlanItem {
	for _, planItem := range *p {
		if reflect.DeepEqual(planItem.FiscalPeriod.Timeframe, fp.Timeframe) {
			return planItem
		}
	}
	return nil
}

type PlanItem struct {
	FiscalPeriod        *FiscalPeriod
	PlanUserDataEntries []*PlanUserDataEntry
}

type PlanUser struct {
	*harvest.User
	BillingDegree             float64
	WorkingDegree             float64
	VacationInterest          float64
	RemainingVacationInterest float64
}

var PlanUserRepository = make(PlanUsers)

type PlanUsers map[int]*PlanUser

func (p *PlanUsers) FindByHarvestUser(harvestUser *harvest.User) *PlanUser {
	user, ok := (*p)[harvestUser.Id()]
	if ok {
		return user
	}
	return nil
}

func (p *PlanUsers) AddUser(planUser *PlanUser) bool {
	harvestUserId := planUser.Id()
	_, ok := (*p)[harvestUserId]
	if ok {
		return !ok
	}
	(*p)[harvestUserId] = planUser
	return true
}

type PlanUserDataEntry struct {
	FiscalPeriod              *FiscalPeriod
	User                      *PlanUser
	BillingDegree             float64
	WorkingDegree             float64
	VacationInterest          float64
	RemainingVacationInterest float64
	DaysOfIllness             float64
}

func (p *PlanUserDataEntry) BusinessDays() float64 {
	return float64(p.FiscalPeriod.BusinessDays) * p.WorkingDegree
}

func (p *PlanUserDataEntry) CumulatedBusinessDays() float64 {
	return float64(p.FiscalPeriod.BusinessDays) * p.WorkingDegree
}

func (p *PlanUserDataEntry) BillableDays() float64 {
	return (p.BusinessDays() - p.VacationInterest - p.RemainingVacationInterest - p.DaysOfIllness) * p.WorkingDegree
}

func (p *PlanUserDataEntry) CumulatedBillableDays() float64 {
	return p.BillableDays()
}

func (p *PlanUserDataEntry) EffectiveBillingDegree() float64 {
	return p.BillableDays() / p.BusinessDays()
}

func (p *PlanUserDataEntry) CumulatedEffectiveBillingDegree() float64 {
	return p.EffectiveBillingDegree()
}

func NewPlanItemFromForm(form url.Values, users []*harvest.User) (*PlanItem, []error) {
	var errors []error
	var planUserDataItems []*PlanUserDataEntry
	timeframe := form.Get("timeframe")
	timeframeQuery, err := url.ParseQuery(timeframe)
	if err != nil {
		errors = append(errors, err)
	}
	fiscalPeriod, err := FiscalPeriodFromQuery(timeframeQuery)
	if err != nil {
		errors = append(errors, err)
	}
	parser := NewFormParser(form)
	for _, user := range users {
		planUser := PlanUserRepository.FindByHarvestUser(user)
		if planUser == nil {
			planUser = &PlanUser{User: user}
			PlanUserRepository.AddUser(planUser)
			//errors = append(errors, fmt.Errorf("Keine Plandaten für Mitarbeiter %s gefunden", user.FirstName))
			//continue
		}
		parser.ResetErrors()
		planUserData, errs := NewPlanUserFromForm(parser, planUser)
		if len(errs) > 0 {
			errors = append(errors, errs...)
		}
		if planUserData != nil {
			planUserData.FiscalPeriod = fiscalPeriod
			planUserDataItems = append(planUserDataItems, planUserData)
		}
	}
	if len(errors) > 0 {
		return nil, errors
	}
	return &PlanItem{PlanUserDataEntries: planUserDataItems, FiscalPeriod: fiscalPeriod}, nil
}

type HarvestUsers []*harvest.User

func (h *HarvestUsers) ById(id int) *harvest.User {
	for _, user := range *h {
		if user.Id() == id {
			return user
		}
	}
	return nil
}

func NewPlanUserFromForm(parser *FormParser, user *PlanUser) (*PlanUserDataEntry, []error) {
	prefix := fmt.Sprintf("%d-", user.User.Id())
	planUser := PlanUserDataEntry{
		User:                      user,
		BillingDegree:             parser.Float64(prefix + "billing-degree"),
		WorkingDegree:             parser.Float64(prefix + "working-degree"),
		VacationInterest:          parser.Float64(prefix + "vacation-interest"),
		RemainingVacationInterest: parser.Float64(prefix + "remaining-vacation-interest"),
		DaysOfIllness:             parser.Float64(prefix + "days-of-illness"),
	}
	errs := parser.GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &planUser, nil
}

func planYearsHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		var cachedPlanYears PlanYears
		planYears := cache.Get("PlanYears")
		if planYears == nil {
			cachedPlanYears = make(PlanYears)
		} else {
			cachedPlanYears = planYears.(PlanYears)
		}
		page := PageForSession(s)
		if r.Method == "GET" {
			var currentYear *FiscalYear
			var fiscalYears FiscalYears
			now := time.Now()
			for year, planYear := range cachedPlanYears {
				if year == now.Year() {
					currentYear = planYear.FiscalYear
				} else {
					fiscalYears = append(fiscalYears, planYear.FiscalYear)
				}
			}
			page.Set("CurrentYear", currentYear)
			page.Set("PlanYears", fiscalYears)
			renderTemplate(w, "plan-years", page)
			return
		}
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				s.AddError(err)
				http.Redirect(w, r, "/plan_years/new", http.StatusFound)
				return
			}
			formParser := NewFormParser(r.PostForm)
			fiscalYear := &FiscalYear{
				Year:                     formParser.Int("year"),
				BusinessDays:             formParser.Int("business-days"),
				CalendarWeeks:            formParser.Int("calendar-weeks"),
				BusinessDaysFirstQuarter: formParser.Int("business-days-first-quarter"),
			}
			planYear := &PlanYear{
				FiscalYear:                fiscalYear,
				AverageDaysOfIllness:      formParser.Float64("average-days-of-illness"),
				AverageDaysOfChildrenCare: formParser.Float64("average-days-of-children-care"),
				DefaultVacationInterest:   formParser.Float64("default-vacation-interest"),
			}
			errs := formParser.GetErrors()
			if len(errs) != 0 {
				for _, err := range errs {
					s.AddError(err)
				}
				http.Redirect(w, r, "/plan_years/new", http.StatusFound)
				return
			}
			job := func() error {
				return planApp.SavePlanYear(planYear)
			}
			workerQueue.AddJob(job)
			cachedPlanYears[planYear.FiscalYear.Year] = planYear
			cache.Store("PlanYears", cachedPlanYears)
			http.Redirect(w, r, "/plan_years", http.StatusFound)
			return
		}
	}
}

func planYearsNewHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		page := PageForSession(s)
		renderTemplate(w, "plan-years-new", page)
	}
}

func planYearsShowHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		yearString := path.Base(r.URL.Path)
		year, err := strconv.Atoi(yearString)
		if err != nil {
			s.AddError(err)
			http.Redirect(w, r, "/plan_years", http.StatusFound)
			return
		}
		var planYear *PlanYear
		var cachedPlanYears PlanYears
		planYears := cache.Get("PlanYears")
		if planYears != nil {
			cachedPlanYears = planYears.(PlanYears)
			planYear = cachedPlanYears.FindByYear(year)
			if planYear == nil {
				s.AddError(fmt.Errorf("Kein Planjahr für %d gefunden.", year))
				http.Redirect(w, r, "/plan_years", http.StatusFound)
				return
			}
		} else {
			job := func() error {
				planYear, err := planApp.LoadPlanYear(year)
				if err != nil {
					return err
				}
				cachedPlanYears = make(PlanYears)
				cachedPlanYears[planYear.FiscalYear.Year] = planYear
				return cache.Store("PlanYears", cachedPlanYears)
			}
			status := workerQueue.AddJob(job)
			if ok, err := status.Success(); !ok {
				if err != nil {
					s.AddError(fmt.Errorf("Kein Planjahr für %d gefunden.", year))
					http.Redirect(w, r, "/plan_years", http.StatusFound)
					return
				}
			}
		}
		page := PageForSession(s)
		page.Set("PlanYear", planYear)
		renderTemplate(w, "plan-years-show", page)
		return
	}
}

func planItemshandler() harvestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session, c *harvest.Harvest) {
		var cachedPlanItems PlanItems
		planItems := cache.Get("PlanItems")
		if planItems == nil {
			cachedPlanItems = make(PlanItems, 0)
		} else {
			cachedPlanItems = planItems.(PlanItems)
		}
		page := PageForSession(s)
		if r.Method == "GET" {
			timeframes := make([]map[string]interface{}, len(cachedPlanItems))
			for i, planItem := range cachedPlanItems {
				fp := planItem.FiscalPeriod
				timeframes[i] = map[string]interface{}{
					"Link":      template.URL(fp.ToQuery().Encode()),
					"StartDate": fp.StartDate,
					"EndDate":   fp.EndDate,
				}
			}
			page.Set("Timeframes", timeframes)
			renderTemplate(w, "plan-items", page)
		}
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				s.AddError(err)
				http.Redirect(w, r, "/plan_items/new", http.StatusFound)
				return
			}
			params := r.Form
			debug.Printf("Form params: %+#v\n", params)
			var users []*harvest.User
			err = c.Users.All(&users, nil)
			if err != nil {
				s.AddError(err)
				http.Redirect(w, r, "/plan_items/new", http.StatusFound)
				return
			}
			planItem, errs := NewPlanItemFromForm(params, users)
			if len(errs) != 0 {
				for _, err := range errs {
					s.AddError(err)
					http.Redirect(w, r, "/plan_items/new", http.StatusFound)
					return
				}
			}
			cachedPlanItems = append(cachedPlanItems, planItem)
			cache.Store("PlanItems", cachedPlanItems)
			http.Redirect(w, r, "/plan_items", http.StatusFound)
		}
	}
}

func planItemNewHandler() harvestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session, c *harvest.Harvest) {
		var users []*harvest.User
		err := c.Users.All(&users, nil)
		if err != nil {
			s.AddError(err)
			http.Redirect(w, r, "/plan_items", http.StatusFound)
			return
		}
		page := PageForSession(s)
		page.Set("Users", users)
		page.Set("Timeframes", fiscalYear.FiscalPeriods())
		debug.Printf("Users: %+#v\n", users)
		renderTemplate(w, "plan-items-new", page)
	}
}

func planItemsShowHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		var cachedPlanItems PlanItems
		planItems := cache.Get("PlanItems")
		if planItems == nil {
			s.AddError(fmt.Errorf("Keine Planungszeiträume gefunden"))
			http.Redirect(w, r, "/plan_items", http.StatusFound)
			return
		} else {
			cachedPlanItems = planItems.(PlanItems)
		}
		params := r.URL.Query()
		tf, err := harvest.TimeframeFromQuery(params)
		if err != nil {
			debug.Printf("Error fetching timeframe from params: sessionId=%s\tparams=%+#v\terror=%T:%v\n", s.id, params, err, err)
			s.AddDebugError(err)
			http.Redirect(w, r, "/plan_items", http.StatusFound)
			return
		}
		businessDays, err := strconv.Atoi(params.Get("business-days"))
		if err != nil {
			debug.Printf("Error fetching timeframe from params: sessionId=%s\tparams=%+#v\terror=%T:%v\n", s.id, params, err, err)
			s.AddDebugError(err)
			http.Redirect(w, r, "/plan_items", http.StatusFound)
			return
		}
		fiscalPeriod := FiscalPeriod{Timeframe: &tf, BusinessDays: businessDays}
		planItem := cachedPlanItems.FindByFiscalPeriod(&fiscalPeriod)
		if planItem == nil {
			debug.Printf("PlanItem not found in %+#v\n", cachedPlanItems)
			s.AddError(fmt.Errorf("Planungszeitraum nicht gefunden"))
			http.Redirect(w, r, "/plan_items", http.StatusFound)
			return
		}
		page := PageForSession(s)
		page.Set("PlanItem", planItem)
		renderTemplate(w, "plan-items-show", page)
	}
}
