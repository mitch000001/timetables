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
		Period: period,
		UserID: "1",
		RemainingVacationInterestDays: NewFloat(userConfig.remainingVacationInterestDays),
	}
	shareOfYear := NewFloat(period.BusinessDays).Div(NewFloat(planConfig.BusinessDays))

	estimationPeriod.VacationInterestDays = NewFloat(userConfig.workingDegree).Mul(NewFloat(planConfig.VacationInterestDays)).Mul(shareOfYear)

	if userConfig.hasChild {
		estimationPeriod.ChildCareDays = NewFloat(planConfig.ChildCareInterestDays).Mul(shareOfYear).Mul(NewFloat(userConfig.workingDegree))
	} else {
		estimationPeriod.ChildCareDays = NewFloat(0)
	}

	sicknessInterestShare := NewFloat(planConfig.SicknessInterestDays).Mul(shareOfYear)
	estimationPeriod.SicknessInterestDays = sicknessInterestShare.Mul(NewFloat(userConfig.workingDegree))

	nonOfficeDays := estimationPeriod.SicknessInterestDays.Add(estimationPeriod.VacationInterestDays).Add(estimationPeriod.RemainingVacationInterestDays).Add(estimationPeriod.ChildCareDays)
	estimationPeriod.OfficeDays = NewFloat(period.BusinessDays).Sub(nonOfficeDays)

	estimationPeriod.BillableDays = estimationPeriod.OfficeDays.Mul(NewFloat(userConfig.billingDegree))
	estimationPeriod.NonbillableDays = estimationPeriod.OfficeDays.Sub(estimationPeriod.BillableDays)

	estimationPeriod.EffectiveBillingDegree = estimationPeriod.BillableDays.Div(NewFloat(period.BusinessDays))

	return estimationPeriod
}

type EstimationBillingPeriodUserEntry struct {
	ID                            string
	Period                        Period
	UserID                        string
	VacationInterestDays          *Float
	RemainingVacationInterestDays *Float
	SicknessInterestDays          *Float
	ChildCareDays                 *Float
	BillableDays                  *Float
	NonbillableDays               *Float
	OfficeDays                    *Float
	EffectiveBillingDegree        *Float
}
