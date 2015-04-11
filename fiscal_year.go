package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"sort"
	"time"
)

type FiscalYear struct {
	ID                       int
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

func (f *FiscalYear) MustAdd(fiscalPeriod *FiscalPeriod) {
	err := f.Add(fiscalPeriod)
	if err != nil {
		panic(err)
	}
}

func InsertFiscalYear(db *sql.DB, fiscalYear *FiscalYear) error {
	const insertSQL = `
		INSERT INTO fiscal_years
			(year, business_days, business_days_first_quarter, calendar_weeks)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`
	return db.QueryRow(insertSQL, fiscalYear.Year, fiscalYear.BusinessDays, fiscalYear.BusinessDaysFirstQuarter, fiscalYear.CalendarWeeks).Scan(&fiscalYear.ID)
}

func FindFiscalYearForYear(db *sql.DB, year int) (*FiscalYear, error) {
	const findSQL = `
		SELECT
			id,
			business_days,
			business_days_first_quarter,
			calendar_weeks,
			year
		FROM fiscal_years
		WHERE year = $1
	`
	row := db.QueryRow(findSQL, year)
	var fiscalYear FiscalYear
	err := row.Scan(
		&fiscalYear.ID,
		&fiscalYear.BusinessDays,
		&fiscalYear.BusinessDaysFirstQuarter,
		&fiscalYear.CalendarWeeks,
		&fiscalYear.Year,
	)
	if err != nil {
		return nil, err
	}
	return &fiscalYear, nil
}

func scanFiscalYear(rows *sql.Rows) (*FiscalYear, error) {
	var fiscalYear FiscalYear
	err := rows.Scan(
		&fiscalYear.ID,
		&fiscalYear.BusinessDays,
		&fiscalYear.BusinessDaysFirstQuarter,
		&fiscalYear.CalendarWeeks,
		&fiscalYear.Year,
	)
	if err != nil {
		return nil, err
	}
	return &fiscalYear, nil
}

type FiscalYears []*FiscalYear

func (f FiscalYears) Len() int           { return len(f) }
func (f FiscalYears) Less(i, j int) bool { return f[i].Year < f[j].Year }
func (f FiscalYears) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
