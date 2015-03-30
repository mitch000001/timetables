package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
)

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
	FiscalPeriod *FiscalPeriod
	PlanUsers    []*PlanUser
}

type PlanUser struct {
	FiscalPeriod              *FiscalPeriod
	User                      *harvest.User
	BillingDegree             float64
	WorkingDegree             float64
	VacationInterest          float64
	RemainingVacationInterest float64
	DaysOfIllness             float64
}

func (p *PlanUser) BusinessDays() float64 {
	return float64(p.FiscalPeriod.BusinessDays) * p.WorkingDegree
}

func (p *PlanUser) CumulatedBusinessDays() float64 {
	return float64(p.FiscalPeriod.BusinessDays) * p.WorkingDegree
}

func (p *PlanUser) BillableDays() float64 {
	return (p.BusinessDays() - p.VacationInterest - p.RemainingVacationInterest - p.DaysOfIllness) * p.WorkingDegree
}

func (p *PlanUser) CumulatedBillableDays() float64 {
	return p.BillableDays()
}

func (p *PlanUser) EffectiveBillingDegree() float64 {
	return p.BillableDays() / p.BusinessDays()
}

func NewPlanItemFromForm(form url.Values, users []*harvest.User) (*PlanItem, []error) {
	var errors []error
	var planUsers []*PlanUser
	timeframe := form.Get("timeframe")
	timeframeQuery, err := url.ParseQuery(timeframe)
	if err != nil {
		errors = append(errors, err)
	}
	fiscalPeriod, err := FiscalPeriodFromQuery(timeframeQuery)
	if err != nil {
		errors = append(errors, err)
	}
	for _, user := range users {
		planUser, errs := NewPlanUserFromForm(form, user)
		if len(errs) > 0 {
			errors = append(errors, errs...)
		}
		if planUser != nil {
			planUser.FiscalPeriod = fiscalPeriod
			planUsers = append(planUsers, planUser)
		}
	}
	if len(errors) > 0 {
		return nil, errors
	}
	return &PlanItem{PlanUsers: planUsers, FiscalPeriod: fiscalPeriod}, nil
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

func NewPlanUserFromForm(form url.Values, user *harvest.User) (*PlanUser, []error) {
	prefix := fmt.Sprintf("%d-", user.Id())
	billingDegree := form.Get(prefix + "billing-degree")
	workingDegree := form.Get(prefix + "working-degree")
	vacationInterest := form.Get(prefix + "vacation-interest")
	remainingVacationInterest := form.Get(prefix + "remaining-vacation-interest")
	daysOfIllness := form.Get(prefix + "days-of-illness")
	parser := NewFormParser()
	planUser := PlanUser{
		User:                      user,
		BillingDegree:             parser.Float64(billingDegree),
		WorkingDegree:             parser.Float64(workingDegree),
		VacationInterest:          parser.Float64(vacationInterest),
		RemainingVacationInterest: parser.Float64(remainingVacationInterest),
		DaysOfIllness:             parser.Float64(daysOfIllness),
	}
	errs := parser.GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &planUser, nil
}

type FormParser struct {
	errors []error
}

func NewFormParser() *FormParser {
	return &FormParser{make([]error, 0)}
}

func (f *FormParser) GetErrors() []error {
	return f.errors
}

func (f *FormParser) Float64(input string) float64 {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		f.errors = append(f.errors, err)
	}
	return num
}

var planItemsTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "plan-items" .}}{{end}}`))

func planItemshandler() harvestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session, c *harvest.Harvest) {
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
	return func(w http.ResponseWriter, r *http.Request, s *session, c *harvest.Harvest) {
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
	return func(w http.ResponseWriter, r *http.Request, s *session) {
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
