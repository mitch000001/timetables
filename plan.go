package main

import (
	"html/template"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
)

type PlanItem struct {
	User                            *User
	FiscalPeriod                    *FiscalPeriod
	CumulatedBusinessDays           int
	BillingDegree                   float64
	WorkingDegree                   float64
	VacationInterest                float64
	RemainingVacationInterest       float64
	DaysOfIllness                   float64
	CumulatedEffectiveBillingDegree float64
}

func NewPlanItemFromForm(form url.Values) (*PlanItem, []error) {
	billingDegree := form.Get("billing-degree")
	workingDegree := form.Get("working-degree")
	vacationInterest := form.Get("vacation-interest")
	remainingVacationInterest := form.Get("remaining-vacation-interest")
	daysOfIllness := form.Get("days-of-illness")
	cumulatedEffectiveBillingDegree := form.Get("cumulated-effective-billing-degree")
	parser := NewFormParser()
	planItem := PlanItem{
		BillingDegree:                   parser.Float64(billingDegree),
		WorkingDegree:                   parser.Float64(workingDegree),
		VacationInterest:                parser.Float64(vacationInterest),
		RemainingVacationInterest:       parser.Float64(remainingVacationInterest),
		DaysOfIllness:                   parser.Float64(daysOfIllness),
		CumulatedEffectiveBillingDegree: parser.Float64(cumulatedEffectiveBillingDegree),
	}
	errs := parser.GetErrors()
	if len(errs) != 0 {
		return nil, errs
	}

	return &planItem, nil
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

func planItemHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *session) {
		var cachedPlanItems []*PlanItem
		planItems := cache.Get("PlanItems")
		if planItems == nil {
			cachedPlanItems = make([]*PlanItem, 0)
		} else {
			cachedPlanItems = planItems.([]*PlanItem)
		}
		page := PageForSession(s)
		if r.Method == "GET" {
			page.Set("PlanItems", cachedPlanItems)
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
			planItem, errs := NewPlanItemFromForm(params)
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

var planItemNewTemplate = template.Must(template.Must(layout.Clone()).Parse(`{{define "content"}}{{template "plan-item-new" .}}{{end}}`))

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
