package timetables

import "github.com/mitch000001/go-harvest/harvest"

type PlanConfig struct {
	Year                          int
	BusinessDays                  float64
	VacationInterestDays          float64
	SicknessInterestDays          float64
	ChildCareInterestDays         float64
	CumulatedBusinessDays         float64
	CumulatedVacationInterestDays float64
	CumulatedSicknessInterestDays float64
	CumulatedBilledDays           float64
}

type UserConfig struct {
	userID                        string
	hasChild                      bool
	billingDegree                 float64
	workingDegree                 float64
	remainingVacationInterestDays float64
}

func CreateEstimationBillingPeriod(period Period, planConfig PlanConfig, userConfig UserConfig) (EstimationBillingPeriod, interface{}) {
	var estimationPeriod = EstimationBillingPeriod{
		ID:                            "10",
		Timeframe:                     period.Timeframe,
		UserID:                        "1",
		BusinessDays:                  NewFloat(period.BusinessDays),
		CumulatedBusinessDays:         NewFloat(period.BusinessDays).Add(NewFloat(planConfig.CumulatedBusinessDays)),
		RemainingVacationInterestDays: NewFloat(userConfig.remainingVacationInterestDays),
	}
	shareOfYear := NewFloat(period.BusinessDays).Div(NewFloat(planConfig.BusinessDays))

	estimationPeriod.VacationInterestDays = NewFloat(userConfig.workingDegree).Mul(NewFloat(planConfig.VacationInterestDays)).Mul(shareOfYear)

	estimationPeriod.CumulatedVacationInterestDays = estimationPeriod.VacationInterestDays.Add(NewFloat(planConfig.CumulatedVacationInterestDays))

	sicknessInterest := NewFloat(planConfig.SicknessInterestDays)
	if userConfig.hasChild {
		sicknessInterest = NewFloat(planConfig.SicknessInterestDays).Add(NewFloat(planConfig.ChildCareInterestDays))
	}
	sicknessInterestShare := sicknessInterest.Mul(shareOfYear)
	estimationPeriod.SicknessInterestDays = sicknessInterestShare.Mul(NewFloat(userConfig.workingDegree))

	estimationPeriod.CumulatedSicknessInterestDays = estimationPeriod.SicknessInterestDays.Add(NewFloat(planConfig.CumulatedSicknessInterestDays))

	unbilled := estimationPeriod.SicknessInterestDays.Add(estimationPeriod.VacationInterestDays).Add(estimationPeriod.RemainingVacationInterestDays)
	estimationPeriod.BilledDays = NewFloat(period.BusinessDays).Sub(unbilled).Mul(NewFloat(userConfig.billingDegree))

	estimationPeriod.CumulatedBilledDays = estimationPeriod.BilledDays.Add(NewFloat(planConfig.CumulatedBilledDays))

	estimationPeriod.EffectiveBillingDegree = estimationPeriod.BilledDays.Div(NewFloat(period.BusinessDays))

	estimationPeriod.CumulatedEffectiveBillingDegree = estimationPeriod.CumulatedBilledDays.Div(estimationPeriod.CumulatedBusinessDays)

	return estimationPeriod, nil
}

type EstimationBillingPeriod struct {
	ID                              string
	Timeframe                       harvest.Timeframe
	UserID                          string
	BusinessDays                    *Float
	CumulatedBusinessDays           *Float
	VacationInterestDays            *Float
	CumulatedVacationInterestDays   *Float
	RemainingVacationInterestDays   *Float
	SicknessInterestDays            *Float
	CumulatedSicknessInterestDays   *Float
	BilledDays                      *Float
	CumulatedBilledDays             *Float
	EffectiveBillingDegree          *Float
	CumulatedEffectiveBillingDegree *Float
}
