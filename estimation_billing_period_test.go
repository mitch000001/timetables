package timetables

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateEstimationBillingPeriodUserEntry(t *testing.T) {
	tests := []struct {
		period          Period
		planConfigInput PlanConfig
		userConfigInput UserConfig
		output          EstimationBillingPeriodUserEntry
	}{
		{
			Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 25, 1, time.Local)}, 10},
			PlanConfig{
				Year:                  2015,
				BusinessDays:          250,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      true,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0.0,
			},
			EstimationBillingPeriodUserEntry{
				ID:                            "10",
				Period:                        Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 25, 1, time.Local)}, 10},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(10).Div(NewFloat(250))),
				RemainingVacationInterestDays: NewFloat(0),
				SicknessInterestDays:          NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(10).Div(NewFloat(250))),
				BilledDays:                    NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:        NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)).Div(NewFloat(10)),
			},
		},
		{
			Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                  2015,
				BusinessDays:          200,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      true,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0.0,
			},
			EstimationBillingPeriodUserEntry{
				ID:                            "10",
				Period:                        Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays: NewFloat(0),
				SicknessInterestDays:          NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                    NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:        NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                  2015,
				BusinessDays:          200,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			EstimationBillingPeriodUserEntry{
				ID:                            "10",
				Period:                        Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays: NewFloat(0),
				SicknessInterestDays:          NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                    NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:        NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                  2015,
				BusinessDays:          200,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 5.0,
			},
			EstimationBillingPeriodUserEntry{
				ID:                            "10",
				Period:                        Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays: NewFloat(5),
				SicknessInterestDays:          NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                    NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:        NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
	}

	for _, test := range tests {
		period, err := CreateEstimationBillingPeriodUserEntry(test.period, test.planConfigInput, test.userConfigInput)

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
