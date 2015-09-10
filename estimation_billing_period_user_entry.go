package timetables

type PlanConfig struct {
	Year                  int
	BusinessDays          float64
	VacationInterestDays  float64
	SicknessInterestDays  float64
	ChildCareInterestDays float64
}

type UserConfig struct {
	userID                        string
	hasChild                      bool
	billingDegree                 float64
	workingDegree                 float64
	remainingVacationInterestDays float64
}

func NewEstimationBillingPeriodUserEntry(period Period, planConfig PlanConfig, userConfig UserConfig) EstimationBillingPeriodUserEntry {
	var estimationPeriod = EstimationBillingPeriodUserEntry{
		ID:     "10",
		Period: period,
		UserID: "1",
		RemainingVacationInterestDays: NewFloat(userConfig.remainingVacationInterestDays),
	}
	shareOfYear := NewFloat(period.BusinessDays).Div(NewFloat(planConfig.BusinessDays))

	estimationPeriod.VacationInterestDays = NewFloat(userConfig.workingDegree).Mul(NewFloat(planConfig.VacationInterestDays)).Mul(shareOfYear)

	sicknessInterest := NewFloat(planConfig.SicknessInterestDays)
	if userConfig.hasChild {
		sicknessInterest = NewFloat(planConfig.SicknessInterestDays).Add(NewFloat(planConfig.ChildCareInterestDays))
	}
	sicknessInterestShare := sicknessInterest.Mul(shareOfYear)
	estimationPeriod.SicknessInterestDays = sicknessInterestShare.Mul(NewFloat(userConfig.workingDegree))

	unbilled := estimationPeriod.SicknessInterestDays.Add(estimationPeriod.VacationInterestDays).Add(estimationPeriod.RemainingVacationInterestDays)
	estimationPeriod.BilledDays = NewFloat(period.BusinessDays).Sub(unbilled).Mul(NewFloat(userConfig.billingDegree))

	estimationPeriod.EffectiveBillingDegree = estimationPeriod.BilledDays.Div(NewFloat(period.BusinessDays))

	return estimationPeriod
}

type EstimationBillingPeriodUserEntry struct {
	ID                            string
	Period                        Period
	UserID                        string
	VacationInterestDays          *Float
	RemainingVacationInterestDays *Float
	SicknessInterestDays          *Float
	BilledDays                    *Float
	EffectiveBillingDegree        *Float
}
