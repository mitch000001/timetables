package timetables

func NewEstimationBillingPeriod(period Period, planConfig PlanConfig) EstimationBillingPeriod {
	return EstimationBillingPeriod{
		Period:      period,
		planConfig:  planConfig,
		userEntries: make([]EstimationBillingPeriodUserEntry, 0),
	}
}

type EstimationBillingPeriod struct {
	ID          string
	Period      Period
	planConfig  PlanConfig
	userEntries []EstimationBillingPeriodUserEntry
}

func (e *EstimationBillingPeriod) AddUserEntry(userConfig UserConfig) {
	userEntry := NewEstimationBillingPeriodUserEntry(e.Period, e.planConfig, userConfig)
	e.userEntries = append(e.userEntries, userEntry)
}
