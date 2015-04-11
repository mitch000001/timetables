package main

import "database/sql"

type FiscalApp struct {
	db *sql.DB
}

func NewFiscalApp(db *sql.DB) *FiscalApp {
	return &FiscalApp{db}
}

func (f *FiscalApp) saveFiscalYear(fiscalYear *FiscalYear) error {
	err := InsertFiscalYear(f.db, fiscalYear)
	if err != nil {
		return err
	}
	for _, fp := range fiscalYear.fiscalPeriods {
		err = InsertFiscalPeriod(f.db, fp, fiscalYear)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FiscalApp) loadFiscalYear(year int) (*FiscalYear, error) {
	fiscalYear, err := FindFiscalYearForYear(f.db, year)
	if err != nil {
		return nil, err
	}
	fiscalPeriods, err := FindFiscalPeriodsForFiscalYear(f.db, fiscalYear)
	if err != nil {
		return nil, err
	}
	for _, fp := range fiscalPeriods {
		fiscalYear.MustAdd(fp)
	}
	return fiscalYear, nil
}
