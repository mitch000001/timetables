package timetables

type PlanConfig struct {
	Year                  int
	BusinessDays          float64
	VacationDays          float64
	SicknessDays          float64
	ChildCareInterestDays float64
}

type UserConfig struct {
	userID                string
	hasChild              bool
	billingDegree         float64
	workingDegree         float64
	remainingVacationDays float64
}

func NewForecastBillingPeriodUserEntry(period Period, planConfig PlanConfig, userConfig UserConfig) ForecastBillingPeriodUserEntry {
	var forecastPeriod = ForecastBillingPeriodUserEntry{
		Period:                period,
		UserID:                userConfig.userID,
		RemainingVacationDays: NewRat(userConfig.remainingVacationDays),
	}
	shareOfYear := NewRat(period.BusinessDays).Div(NewRat(planConfig.BusinessDays))

	forecastPeriod.VacationDays = NewRat(userConfig.workingDegree).Mul(NewRat(planConfig.VacationDays)).Mul(shareOfYear)

	if userConfig.hasChild {
		forecastPeriod.ChildCareDays = NewRat(planConfig.ChildCareInterestDays).Mul(shareOfYear).Mul(NewRat(userConfig.workingDegree))
	} else {
		forecastPeriod.ChildCareDays = NewRat(0)
	}

	sicknessInterestShare := NewRat(planConfig.SicknessDays).Mul(shareOfYear)
	forecastPeriod.SicknessDays = sicknessInterestShare.Mul(NewRat(userConfig.workingDegree))

	nonOfficeDays := forecastPeriod.SicknessDays.Add(forecastPeriod.VacationDays).Add(forecastPeriod.RemainingVacationDays).Add(forecastPeriod.ChildCareDays)
	forecastPeriod.OfficeDays = NewRat(period.BusinessDays).Sub(nonOfficeDays)

	billingDegree := NewRat(userConfig.billingDegree)

	forecastPeriod.BillableDays = forecastPeriod.OfficeDays.Mul(billingDegree)
	forecastPeriod.NonbillableDays = forecastPeriod.OfficeDays.Mul(NewRat(1).Sub(billingDegree))

	forecastPeriod.EffectiveBillingDegree = forecastPeriod.BillableDays.Div(NewRat(period.BusinessDays))

	return forecastPeriod
}

type ForecastBillingPeriodUserEntry struct {
	ID                              string
	Period                          Period
	UserID                          string
	VacationDays                    *Rat
	CumulatedVacationDays           *Rat
	RemainingVacationDays           *Rat
	SicknessDays                    *Rat
	CumulatedSicknessDays           *Rat
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
