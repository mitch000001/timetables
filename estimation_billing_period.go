package timetables

import "github.com/mitch000001/go-harvest/harvest"

type PlanConfig struct {
	BusinessDays      float64
	VacationInterest  float64
	SicknessInterest  float64
	ChildCareInterest float64
}

type UserConfig struct {
	userID                    string
	hasChild                  bool
	billingDegree             float64
	workingDegree             float64
	remainingVacationInterest float64
}

func CreateEstimationBillingPeriod(period Period, planConfig PlanConfig, userConfig UserConfig) (EstimationBillingPeriod, interface{}) {
	var estimationPeriod = EstimationBillingPeriod{
		ID:                              "10",
		Timeframe:                       period.Timeframe,
		UserID:                          "1",
		BusinessDays:                    period.BusinessDays,
		CumulatedBusinessDays:           10.0,
		CumulatedVacationInterest:       1.71,
		RemainingVacationInterest:       userConfig.remainingVacationInterest,
		CumulatedSicknessInterest:       1.08,
		CumulatedBilledDays:             10.57,
		CumulatedEffectiveBillingDegree: 0.66,
	}
	shareOfYear := NewFloat(period.BusinessDays).Div(NewFloat(planConfig.BusinessDays))

	estimationPeriod.VacationInterest = NewFloat(userConfig.workingDegree).Mul(NewFloat(planConfig.VacationInterest)).Mul(shareOfYear)

	sicknessInterest := NewFloat(planConfig.SicknessInterest)
	if userConfig.hasChild {
		sicknessInterest = NewFloat(planConfig.SicknessInterest).Add(NewFloat(planConfig.ChildCareInterest))
	}
	sicknessInterestShare := sicknessInterest.Mul(shareOfYear)
	estimationPeriod.SicknessInterest = sicknessInterestShare.Mul(NewFloat(userConfig.workingDegree))

	unbilled := estimationPeriod.SicknessInterest.Add(estimationPeriod.VacationInterest).Add(NewFloat(estimationPeriod.RemainingVacationInterest))
	estimationPeriod.BilledDays = NewFloat(period.BusinessDays).Sub(unbilled).Mul(NewFloat(userConfig.billingDegree))

	estimationPeriod.EffectiveBillingDegree = estimationPeriod.BilledDays.Div(NewFloat(period.BusinessDays))

	return estimationPeriod, nil
}

type EstimationBillingPeriod struct {
	ID                              string
	Timeframe                       harvest.Timeframe
	UserID                          string
	BusinessDays                    float64
	CumulatedBusinessDays           float64
	VacationInterest                *Float
	CumulatedVacationInterest       float64
	RemainingVacationInterest       float64
	SicknessInterest                *Float
	CumulatedSicknessInterest       float64
	BilledDays                      *Float
	CumulatedBilledDays             float64
	EffectiveBillingDegree          *Float
	CumulatedEffectiveBillingDegree float64
}
