package main

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

type FiscalPeriod struct {
	*harvest.Timeframe
	BusinessDays int
}

func (f *FiscalPeriod) ToQuery() url.Values {
	params := f.Timeframe.ToQuery()
	params.Set("business-days", fmt.Sprintf("%d", f.BusinessDays))
	return params
}

func FiscalPeriodFromQuery(params url.Values) (*FiscalPeriod, error) {
	tf, err := harvest.TimeframeFromQuery(params)
	if err != nil {
		return nil, err
	}
	businessDays, err := strconv.Atoi(params.Get("business-days"))
	if err != nil {
		return nil, err
	}
	return &FiscalPeriod{Timeframe: &tf, BusinessDays: businessDays}, nil
}

func NewFiscalPeriod(start time.Time, end time.Time, businessDays int) *FiscalPeriod {
	timeframe := harvest.Timeframe{StartDate: harvest.NewShortDate(start), EndDate: harvest.NewShortDate(end)}
	return &FiscalPeriod{
		Timeframe:    &timeframe,
		BusinessDays: businessDays,
	}
}

func (f *FiscalPeriod) InBetween(date time.Time) bool {
	return f.StartDate.Before(date) && f.EndDate.After(date)
}

func (f *FiscalPeriod) Overlapping(other *FiscalPeriod) bool {
	return f.InBetween(other.StartDate.Time) || f.InBetween(other.EndDate.Time)
}

type FiscalPeriods []*FiscalPeriod

func (f FiscalPeriods) Len() int           { return len(f) }
func (f FiscalPeriods) Less(i, j int) bool { return f[i].StartDate.Before(f[j].StartDate.Time) }
func (f FiscalPeriods) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

type ReverseSortedFiscalPeriods []*FiscalPeriod

func (f ReverseSortedFiscalPeriods) Len() int { return len(f) }
func (f ReverseSortedFiscalPeriods) Less(i, j int) bool {
	return f[j].StartDate.Before(f[i].StartDate.Time)
}
func (f ReverseSortedFiscalPeriods) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

type FiscalYear struct {
	fiscalPeriods            FiscalPeriods
	Year                     int
	BusinessDays             int
	CalendarWeeks            int
	BusinessDaysFirstQuarter int
}

func (f *FiscalYear) init() {
	if f.fiscalPeriods == nil {
		f.fiscalPeriods = make(FiscalPeriods, 0)
	}
}

func (f *FiscalYear) BusinessDaysInFirstQuarter() int {
	return f.BusinessDays / 4
}

// CurrentFiscalPeriod returns the FiscalPeriod for the present day
// it returns nil if there is none.
//
// The method is optimized for the common case, i.e. the current period is the last element
// in the slice of periods
func (f *FiscalYear) CurrentFiscalPeriod() *FiscalPeriod {
	f.init()
	now := time.Now()
	reverseSorted := ReverseSortedFiscalPeriods(f.fiscalPeriods)
	defer sort.Sort(f.fiscalPeriods)
	idx := sort.Search(len(f.fiscalPeriods), func(i int) bool {
		return reverseSorted[i].InBetween(now)
	})
	if idx != len(f.fiscalPeriods) {
		return f.fiscalPeriods[idx]
	}
	return nil
}

func (f *FiscalYear) PastFiscalPeriods() FiscalPeriods {
	f.init()
	now := time.Now()
	var pastFiscalPeriods FiscalPeriods
	for _, fp := range f.fiscalPeriods {
		if !fp.InBetween(now) {
			pastFiscalPeriods = append(pastFiscalPeriods, fp)
		}
	}
	sort.Sort(pastFiscalPeriods)
	return pastFiscalPeriods
}

func (f *FiscalYear) FiscalPeriods() FiscalPeriods {
	return f.fiscalPeriods
}

func (f *FiscalYear) Add(fiscalPeriod *FiscalPeriod) error {
	f.init()
	if fiscalPeriod.StartDate.Year() != f.Year || fiscalPeriod.EndDate.Year() != f.Year {
		return fmt.Errorf("Der Abrechnungszeitraum wurde für das falsche Jahr angelegt.")
	}
	idx := sort.Search(len(f.fiscalPeriods), func(i int) bool {
		return f.fiscalPeriods[i].Overlapping(fiscalPeriod)
	})
	if idx == len(f.fiscalPeriods) && idx != 0 {
		return fmt.Errorf("Die Abrechnungszeiträume dürfen sich nicht überlappen.")
	}
	f.fiscalPeriods = append(f.fiscalPeriods, fiscalPeriod)
	sort.Sort(f.fiscalPeriods)
	return nil
}

func (f *FiscalYear) Delete(fiscalPeriod *FiscalPeriod) {
	f.init()
	var newFiscalPeriods FiscalPeriods
	for _, fp := range f.fiscalPeriods {
		if !reflect.DeepEqual(fp, fiscalPeriod) && !reflect.DeepEqual(fp.Timeframe, fiscalPeriod.Timeframe) {
			newFiscalPeriods = append(newFiscalPeriods, fp)
		}
	}
	sort.Sort(newFiscalPeriods)
	f.fiscalPeriods = newFiscalPeriods
}

type FiscalYears []*FiscalYear

func (f FiscalYears) Len() int           { return len(f) }
func (f FiscalYears) Less(i, j int) bool { return f[i].Year < f[j].Year }
func (f FiscalYears) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
