package main

import "github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"

type PlanItem struct {
	User                  *User
	Timeframe             harvest.Timeframe
	BusinessDays          int
	CumulatedBusinessDays int
	BillingDegree         float64
	WorkingDegree         float64
	VacationInterest      float64
}
