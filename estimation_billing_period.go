package timetables

import "github.com/mitch000001/go-harvest/harvest"

type PlanConfig struct {
	Year                      int
	BusinessDays              float64
	VacationInterest          float64
	SicknessInterest          float64
	ChildCareInterest         float64
	CumulatedBusinessDays     float64
	CumulatedVacationInterest float64
	CumulatedSicknessInterest float64
	CumulatedBilledDays       float64
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
		ID:                        "10",
		Timeframe:                 period.Timeframe,
		UserID:                    "1",
		BusinessDays:              NewFloat(period.BusinessDays),
		CumulatedBusinessDays:     NewFloat(period.BusinessDays).Add(NewFloat(planConfig.CumulatedBusinessDays)),
		RemainingVacationInterest: NewFloat(userConfig.remainingVacationInterest),
	}
	shareOfYear := NewFloat(period.BusinessDays).Div(NewFloat(planConfig.BusinessDays))

	estimationPeriod.VacationInterest = NewFloat(userConfig.workingDegree).Mul(NewFloat(planConfig.VacationInterest)).Mul(shareOfYear)

	estimationPeriod.CumulatedVacationInterest = estimationPeriod.VacationInterest.Add(NewFloat(planConfig.CumulatedVacationInterest))

	sicknessInterest := NewFloat(planConfig.SicknessInterest)
	if userConfig.hasChild {
		sicknessInterest = NewFloat(planConfig.SicknessInterest).Add(NewFloat(planConfig.ChildCareInterest))
	}
	sicknessInterestShare := sicknessInterest.Mul(shareOfYear)
	estimationPeriod.SicknessInterest = sicknessInterestShare.Mul(NewFloat(userConfig.workingDegree))

	estimationPeriod.CumulatedSicknessInterest = estimationPeriod.SicknessInterest.Add(NewFloat(planConfig.CumulatedSicknessInterest))

	unbilled := estimationPeriod.SicknessInterest.Add(estimationPeriod.VacationInterest).Add(estimationPeriod.RemainingVacationInterest)
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
	VacationInterest                *Float
	CumulatedVacationInterest       *Float
	RemainingVacationInterest       *Float
	SicknessInterest                *Float
	CumulatedSicknessInterest       *Float
	BilledDays                      *Float
	CumulatedBilledDays             *Float
	EffectiveBillingDegree          *Float
	CumulatedEffectiveBillingDegree *Float
}
