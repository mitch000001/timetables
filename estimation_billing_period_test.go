package timetables

import (
	"fmt"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

func TestEstimationBillingPeriodCreate(t *testing.T) {
	tests := []struct {
		timeframe       harvest.Timeframe
		planConfigInput PlanConfig
		userConfigInput UserConfig
		output          EstimationBillingPeriod
	}{
		{
			harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)},
			PlanConfig{
				BusinessDays:      250,
				VacationInterest:  25,
				SicknessInterest:  5,
				ChildCareInterest: 5,
			},
			UserConfig{
				userID:                    "1",
				hasChild:                  true,
				billingDegree:             0.8,
				workingDegree:             1.0,
				businessDays:              10,
				remainingVacationInterest: 0.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)},
				UserID:                          "1",
				BusinessDays:                    10.0,
				CumulatedBusinessDays:           10.0,
				VacationInterest:                NewFloat(25).Mul(NewFloat(10).Div(NewFloat(250))),
				CumulatedVacationInterest:       1.71,
				RemainingVacationInterest:       0,
				SicknessInterest:                NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(10).Div(NewFloat(250))),
				CumulatedSicknessInterest:       1.08,
				BilledDays:                      NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             10.57,
				EffectiveBillingDegree:          NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)).Div(NewFloat(10)),
				CumulatedEffectiveBillingDegree: 0.66,
			},
		},
		{
			harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
			PlanConfig{
				BusinessDays:      200,
				VacationInterest:  25,
				SicknessInterest:  5,
				ChildCareInterest: 5,
			},
			UserConfig{
				userID:                    "1",
				hasChild:                  true,
				billingDegree:             0.8,
				workingDegree:             1.0,
				businessDays:              20,
				remainingVacationInterest: 0.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    20.0,
				CumulatedBusinessDays:           10.0,
				VacationInterest:                NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterest:       1.71,
				RemainingVacationInterest:       0,
				SicknessInterest:                NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterest:       1.08,
				BilledDays:                      NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             10.57,
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: 0.66,
			},
		},
		{
			harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
			PlanConfig{
				BusinessDays:      200,
				VacationInterest:  25,
				SicknessInterest:  5,
				ChildCareInterest: 5,
			},
			UserConfig{
				userID:                    "1",
				hasChild:                  false,
				billingDegree:             0.8,
				workingDegree:             1.0,
				businessDays:              20,
				remainingVacationInterest: 0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    20.0,
				CumulatedBusinessDays:           10.0,
				VacationInterest:                NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterest:       1.71,
				RemainingVacationInterest:       0,
				SicknessInterest:                NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterest:       1.08,
				BilledDays:                      NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             10.57,
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: 0.66,
			},
		},
		{
			harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
			PlanConfig{
				BusinessDays:      200,
				VacationInterest:  25,
				SicknessInterest:  5,
				ChildCareInterest: 5,
			},
			UserConfig{
				userID:                    "1",
				hasChild:                  false,
				billingDegree:             0.8,
				workingDegree:             1.0,
				businessDays:              20,
				remainingVacationInterest: 5.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    20.0,
				CumulatedBusinessDays:           10.0,
				VacationInterest:                NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterest:       1.71,
				RemainingVacationInterest:       5,
				SicknessInterest:                NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterest:       1.08,
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             10.57,
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: 0.66,
			},
		},
	}

	for _, test := range tests {
		period, err := CreateEstimationBillingPeriod(test.timeframe, test.planConfigInput, test.userConfigInput)

		if err != nil {
			t.Logf("Expected no error, got %T:%v", err, err)
			t.Fail()
		}

		expectedPeriod := test.output

		// TODO: reflect.DeepEqual won't work with big.Float
		if fmt.Sprintf("%#v", expectedPeriod) != fmt.Sprintf("%#v", period) {
			t.Logf("Expected estimation period to equal\n%#v\n\tgot:\n%#v\n", expectedPeriod, period)
			t.Fail()
		}
	}
}
