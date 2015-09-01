package timetables

import "github.com/mitch000001/go-harvest/harvest"

func NewPeriod(timeframe harvest.Timeframe, businessDays int) Period {
	return Period{
		Timeframe: timeframe,
	}
}

type Period struct {
	Timeframe harvest.Timeframe
}
