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
	timeframe                 harvest.Timeframe
	billingDegree             float64
	workingDegree             float64
	businessDays              float64
	cumulatedBusinessDays     float64
	remainingVacationInterest float64
}

func CreateEstimationBillingPeriod(planConfig PlanConfig, userConfig UserConfig) (EstimationBillingPeriod, interface{}) {
	var period = EstimationBillingPeriod{
		ID:                              "10",
		Timeframe:                       userConfig.timeframe,
		UserID:                          "1",
		BusinessDays:                    userConfig.businessDays,
		CumulatedBusinessDays:           10.0,
		CumulatedVacationInterest:       1.71,
		RemainingVacationInterest:       0,
		CumulatedSicknessInterest:       1.08,
		CumulatedBilledDays:             10.57,
		CumulatedEffectiveBillingDegree: 0.66,
	}
	shareOfYear := NewFloat(userConfig.businessDays).Div(NewFloat(planConfig.BusinessDays))

	period.VacationInterest = NewFloat(userConfig.workingDegree).Mul(NewFloat(planConfig.VacationInterest)).Mul(shareOfYear)

	sicknessInterestIncludingChildCare := NewFloat(planConfig.SicknessInterest).Add(NewFloat(planConfig.ChildCareInterest))
	sicknessInterestIncludingChildCareShare := sicknessInterestIncludingChildCare.Mul(shareOfYear)
	period.SicknessInterest = sicknessInterestIncludingChildCareShare.Mul(NewFloat(userConfig.workingDegree))

	unbilled := period.SicknessInterest.Add(period.VacationInterest).Add(NewFloat(period.RemainingVacationInterest))
	period.BilledDays = NewFloat(userConfig.businessDays).Sub(unbilled).Mul(NewFloat(userConfig.billingDegree))

	period.EffectiveBillingDegree = period.BilledDays.Div(NewFloat(userConfig.businessDays))

	return period, nil
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
