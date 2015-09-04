package timetables

type PeriodProvider interface {
	Period() Period
}

func NewPeriod(timeframe Timeframe, businessDays float64) Period {
	if businessDays <= 0 {
		businessDays = float64(timeframe.Days())
	}
	return Period{
		Timeframe:    timeframe,
		BusinessDays: businessDays,
	}
}

type Period struct {
	Timeframe    Timeframe
	BusinessDays float64
}
