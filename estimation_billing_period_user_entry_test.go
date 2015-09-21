package timetables

import (
	"fmt"
	"testing"
	"time"
)

func TestNewEstimationBillingPeriodUserEntry(t *testing.T) {
	tests := []struct {
		description     string
		period          Period
		planConfigInput PlanConfig
		userConfigInput UserConfig
		output          EstimationBillingPeriodUserEntry
	}{
		{
			description: "Has child, no remaining vacation, 250 business days, 10 in period",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			planConfigInput: PlanConfig{
				Year:                  2015,
				BusinessDays:          250,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			userConfigInput: UserConfig{
				userID:                        "1",
				hasChild:                      true,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0.0,
			},
			output: EstimationBillingPeriodUserEntry{
				Period:                        Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(10).Div(NewFloat(250))),
				RemainingVacationInterestDays: NewFloat(0),
				SicknessInterestDays:          NewFloat(5).Mul(NewFloat(10).Div(NewFloat(250))),
				ChildCareDays:                 NewFloat(5).Mul(NewFloat(10).Div(NewFloat(250))),
				BillableDays:                  NewFloat(10).Sub(NewFloat(35).Mul(NewFloat(10).Div(NewFloat(250)))).Mul(NewFloat(0.8)),
				NonbillableDays:               NewFloat(10).Sub(NewFloat(35).Mul(NewFloat(10).Div(NewFloat(250)))).Mul(NewFloat(0.2)),
				OfficeDays:                    NewFloat(10).Sub(NewFloat(35).Mul(NewFloat(10).Div(NewFloat(250)))),
				EffectiveBillingDegree:        NewFloat(10).Sub(NewFloat(35).Mul(NewFloat(10).Div(NewFloat(250)))).Mul(NewFloat(0.8)).Div(NewFloat(10)),
			},
		},
		{
			description: "Has child, no remaining vacation, 200 business days, 20 in period",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 26, time.Local), EndDate: Date(2015, 2, 22, time.Local)}, 20},
			planConfigInput: PlanConfig{
				Year:                  2015,
				BusinessDays:          200,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			userConfigInput: UserConfig{
				userID:                        "1",
				hasChild:                      true,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0.0,
			},
			output: EstimationBillingPeriodUserEntry{
				Period:                        Period{Timeframe{StartDate: Date(2015, 1, 26, time.Local), EndDate: Date(2015, 2, 22, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays: NewFloat(0),
				SicknessInterestDays:          NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				ChildCareDays:                 NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BillableDays:                  NewFloat(20).Sub(NewFloat(35).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.8)),
				NonbillableDays:               NewFloat(20).Sub(NewFloat(35).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.2)),
				OfficeDays:                    NewFloat(20).Sub(NewFloat(35).Mul(NewFloat(20).Div(NewFloat(200)))),
				EffectiveBillingDegree:        NewFloat(20).Sub(NewFloat(35).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			description: "Has no child, no remaining vacation, 200 business days, 20 in period",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 26, time.Local), EndDate: Date(2015, 2, 22, time.Local)}, 20},
			planConfigInput: PlanConfig{
				Year:                  2015,
				BusinessDays:          200,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			userConfigInput: UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			output: EstimationBillingPeriodUserEntry{
				Period:                        Period{Timeframe{StartDate: Date(2015, 1, 26, time.Local), EndDate: Date(2015, 2, 22, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays: NewFloat(0),
				SicknessInterestDays:          NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				ChildCareDays:                 NewFloat(0),
				BillableDays:                  NewFloat(20).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.8)),
				NonbillableDays:               NewFloat(20).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.2)),
				OfficeDays:                    NewFloat(20).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))),
				EffectiveBillingDegree:        NewFloat(20).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			description: "Has no child, 5 vacation days remaining, 200 business days, 20 in period",
			period:      Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
			planConfigInput: PlanConfig{
				Year:                  2015,
				BusinessDays:          200,
				VacationInterestDays:  25,
				SicknessInterestDays:  5,
				ChildCareInterestDays: 5,
			},
			userConfigInput: UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 5.0,
			},
			output: EstimationBillingPeriodUserEntry{
				Period:                        Period{Timeframe{StartDate: Date(2015, 26, 1, time.Local), EndDate: Date(2015, 22, 2, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays: NewFloat(5),
				SicknessInterestDays:          NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				ChildCareDays:                 NewFloat(0),
				BillableDays:                  NewFloat(20).Sub(NewFloat(5)).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.8)),
				NonbillableDays:               NewFloat(20).Sub(NewFloat(5)).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.2)),
				OfficeDays:                    NewFloat(20).Sub(NewFloat(5)).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))),
				EffectiveBillingDegree:        NewFloat(20).Sub(NewFloat(5)).Sub(NewFloat(30).Mul(NewFloat(20).Div(NewFloat(200)))).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
	}

	for _, test := range tests {
		period := NewEstimationBillingPeriodUserEntry(test.period, test.planConfigInput, test.userConfigInput)

		expectedPeriod := test.output

		// TODO: reflect.DeepEqual won't work with big.Float
		if fmt.Sprintf("%+v", expectedPeriod) != fmt.Sprintf("%+v", period) {
			t.Logf("Used configuration: %s\n", test.description)
			t.Logf("Expected estimation period to equal\n%+v\n\tgot:\n%+v\n", expectedPeriod, period)
			t.Fail()
		}
	}
}
