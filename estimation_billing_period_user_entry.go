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
		UserID: userConfig.userID,
		RemainingVacationInterestDays: NewRat(userConfig.remainingVacationInterestDays),
	}
	shareOfYear := NewRat(period.BusinessDays).Div(NewRat(planConfig.BusinessDays))

	estimationPeriod.VacationInterestDays = NewRat(userConfig.workingDegree).Mul(NewRat(planConfig.VacationInterestDays)).Mul(shareOfYear)

	if userConfig.hasChild {
		estimationPeriod.ChildCareDays = NewRat(planConfig.ChildCareInterestDays).Mul(shareOfYear).Mul(NewRat(userConfig.workingDegree))
	} else {
		estimationPeriod.ChildCareDays = NewRat(0)
	}

	sicknessInterestShare := NewRat(planConfig.SicknessInterestDays).Mul(shareOfYear)
	estimationPeriod.SicknessInterestDays = sicknessInterestShare.Mul(NewRat(userConfig.workingDegree))

	nonOfficeDays := estimationPeriod.SicknessInterestDays.Add(estimationPeriod.VacationInterestDays).Add(estimationPeriod.RemainingVacationInterestDays).Add(estimationPeriod.ChildCareDays)
	estimationPeriod.OfficeDays = NewRat(period.BusinessDays).Sub(nonOfficeDays)

	billingDegree := NewRat(userConfig.billingDegree)

	estimationPeriod.BillableDays = estimationPeriod.OfficeDays.Mul(billingDegree)
	estimationPeriod.NonbillableDays = estimationPeriod.OfficeDays.Mul(NewRat(1).Sub(billingDegree))

	estimationPeriod.EffectiveBillingDegree = estimationPeriod.BillableDays.Div(NewRat(period.BusinessDays))

	return estimationPeriod
}

type EstimationBillingPeriodUserEntry struct {
	ID                              string
	Period                          Period
	UserID                          string
	VacationInterestDays            *Rat
	CumulatedVacationInterestDays   *Rat
	RemainingVacationInterestDays   *Rat
	SicknessInterestDays            *Rat
	CumulatedSicknessInterestDays   *Rat
	ChildCareDays                   *Rat
	CumulatedChildCareDays          *Rat
	BillableDays                    *Rat
	CumulatedBillableDays           *Rat
	NonbillableDays                 *Rat
	CumulatedNonbillableDays        *Rat
	OfficeDays                      *Rat
	CumulatedOfficeDays             *Rat
	EffectiveBillingDegree          *Rat
	CumulatedEffectiveBillingDegree *Rat
}
