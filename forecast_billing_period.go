package timetables

func NewForecastBillingPeriod(period Period, planConfig PlanConfig) ForecastBillingPeriod {
	return ForecastBillingPeriod{
		Period:      period,
		planConfig:  planConfig,
		userEntries: make([]ForecastBillingPeriodUserEntry, 0),
	}
}

type ForecastBillingPeriod struct {
	ID          string
	Period      Period
	planConfig  PlanConfig
	userEntries []ForecastBillingPeriodUserEntry
}

func (e *ForecastBillingPeriod) AddUserEntry(userConfig UserConfig) {
	userEntry := NewForecastBillingPeriodUserEntry(e.Period, e.planConfig, userConfig)
	e.userEntries = append(e.userEntries, userEntry)
}

func (e *ForecastBillingPeriod) UserEntries() []ForecastBillingPeriodUserEntry {
	return e.userEntries
}
