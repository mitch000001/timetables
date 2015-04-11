package main

import (
	"database/sql"
	"fmt"
)

type PlanApp struct {
	db    *sql.DB
	cache Cache
}

func NewPlanApp(db *sql.DB, cache Cache) *PlanApp {
	return &PlanApp{db: db, cache: cache}
}

func (p *PlanApp) SavePlanYear(planYear *PlanYear) error {
	fiscalApp := NewFiscalApp(p.db)
	err := fiscalApp.SaveFiscalYear(planYear.FiscalYear)
	if err != nil {
		return err
	}
	err = p.insertPlanYear(planYear)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlanApp) LoadAllPlanYears() (PlanYears, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

func (p *PlanApp) LoadPlanYear(year int) (*PlanYear, error) {
	fiscalApp := NewFiscalApp(p.db)
	fiscalYear, err := fiscalApp.LoadFiscalYear(year)
	if err != nil {
		return nil, err
	}
	planYear, err := p.loadPlanYearForFiscalYear(fiscalYear)
	if err != nil {
		return nil, err
	}
	planYear.FiscalYear = fiscalYear
	return planYear, nil
}

func (p *PlanApp) insertPlanYear(planYear *PlanYear) error {
	const insertSQL = `
		INSERT INTO plan_years
			(average_days_of_illness, average_days_of_children_care, default_vacation_interest, fiscal_year_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`
	return p.db.QueryRow(insertSQL,
		planYear.AverageDaysOfIllness,
		planYear.AverageDaysOfChildrenCare,
		planYear.DefaultVacationInterest,
		planYear.FiscalYear.ID,
	).Scan(&planYear.ID)
}

func (p *PlanApp) loadPlanYearForFiscalYear(fiscalYear *FiscalYear) (*PlanYear, error) {
	const findSQL = `
		SELECT
			id,
			average_days_of_illness,
			average_days_of_children_care,
			default_vacation_interest
		FROM plan_years
		WHERE fiscal_year_id = $1
	`
	var planYear PlanYear
	err := p.db.QueryRow(findSQL, fiscalYear.ID).Scan(
		&planYear.ID,
		&planYear.AverageDaysOfIllness,
		&planYear.AverageDaysOfChildrenCare,
		&planYear.DefaultVacationInterest,
	)
	if err != nil {
		return nil, err
	}
	return &planYear, nil
}

func (p *PlanApp) loadAllPlanYears() {}
