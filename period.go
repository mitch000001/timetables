package timetables

import "github.com/mitch000001/go-harvest/harvest"

func NewPeriod(timeframe harvest.Timeframe, businessDays float64) Period {
	return Period{
		Timeframe:    timeframe,
		BusinessDays: businessDays,
	}
}

type Period struct {
	Timeframe    harvest.Timeframe
	BusinessDays float64
}
