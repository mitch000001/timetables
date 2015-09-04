package timetables

type PlanPeriod struct {
	BillingPeriod           BillingPeriodUserEntry
	EstimationBillingPeriod EstimationBillingPeriodUserEntry
}

type Plan struct {
	periods []PlanPeriod
}

func (p Plan) Periods() []PlanPeriod {
	return p.periods
}
