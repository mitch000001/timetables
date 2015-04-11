package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
)

type FiscalPeriod struct {
	ID int
	*harvest.Timeframe
	BusinessDays int
}

func (f *FiscalPeriod) MarshalDb(db *sql.DB) error {
	const insertSQL = `
		INSERT INTO fiscal_periods
			(starts_at, ends_at, business_days)
		VALUES
			($1, $2, $3)
		RETURNING id
	`
	return db.QueryRow(insertSQL, f.StartDate, f.EndDate, f.BusinessDays).Scan(&f.ID)
}

func InsertFiscalPeriod(db *sql.DB, f *FiscalPeriod, fy *FiscalYear) error {
	const insertSQL = `
		INSERT INTO fiscal_periods
			(starts_at, ends_at, business_days, fiscal_year_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`
	return db.QueryRow(insertSQL, f.StartDate, f.EndDate, f.BusinessDays, fy.ID).Scan(&f.ID)
}

func (f *FiscalPeriod) UnmarshalDb(db *sql.DB, id int) error {
	const findSQL = `
		SELECT
			id,
			starts_at,
			ends_at,
			business_days
		FROM fiscal_periods
		WHERE id = $1
	`
	return db.QueryRow(findSQL, id).Scan(
		&f.ID,
		&f.StartDate,
		&f.EndDate,
		&f.BusinessDays,
	)
}

const selectSQL = `
	SELECT
		id,
		starts_at,
		ends_at,
		business_days
	FROM fiscal_periods
	ORDER BY starts_at ASC
`

const findSQL = `
	SELECT
		id,
		starts_at,
		ends_at,
		business_days
	FROM fiscal_periods
	WHERE id IN ($1)
	ORDER BY starts_at ASC
`

func FindFiscalPeriods(db *sql.DB, ids []int) (FiscalPeriods, error) {
	var fiscalPeriods FiscalPeriods
	rows, err := db.Query(findSQL, ids)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		fp, err := scanFiscalPeriod(rows)
		if err != nil {
			return nil, err
		}
		fiscalPeriods = append(fiscalPeriods, fp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return fiscalPeriods, nil
}

const findForFiscalYearSQL = `
	SELECT
		id,
		starts_at,
		ends_at,
		business_days
	FROM fiscal_periods
	WHERE fiscal_year_id = $1
	ORDER BY starts_at ASC
`

func FindFiscalPeriodsForFiscalYear(db *sql.DB, fiscalYear *FiscalYear) (FiscalPeriods, error) {
	var fiscalPeriods FiscalPeriods
	rows, err := db.Query(findForFiscalYearSQL, fiscalYear.ID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		fp, err := scanFiscalPeriod(rows)
		if err != nil {
			return nil, err
		}
		fiscalPeriods = append(fiscalPeriods, fp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return fiscalPeriods, nil
}

func scanFiscalPeriod(rows *sql.Rows) (*FiscalPeriod, error) {
	var fiscalPeriod FiscalPeriod
	err := rows.Scan(
		&fiscalPeriod.ID,
		&fiscalPeriod.StartDate,
		&fiscalPeriod.EndDate,
		&fiscalPeriod.BusinessDays,
	)
	if err != nil {
		return nil, err
	}
	return &fiscalPeriod, nil
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
