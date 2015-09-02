package timetables

import (
	"fmt"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

func TestEstimationBillingPeriodCreate(t *testing.T) {
	tests := []struct {
		period          Period
		planConfigInput PlanConfig
		userConfigInput UserConfig
		output          EstimationBillingPeriod
	}{
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)}, 10},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  250,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         0,
				CumulatedVacationInterestDays: 0,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      true,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(10.0),
				CumulatedBusinessDays:           NewFloat(10.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(10).Div(NewFloat(250))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(10).Div(NewFloat(250))),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(10).Div(NewFloat(250))),
				CumulatedSicknessInterestDays:   NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(10).Div(NewFloat(250))),
				BilledDays:                      NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)).Div(NewFloat(10)),
				CumulatedEffectiveBillingDegree: NewFloat(10).Sub(NewFloat(1)).Sub(NewFloat(0.4)).Mul(NewFloat(0.8)).Div(NewFloat(10)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         0,
				CumulatedVacationInterestDays: 0,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      true,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(20.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Add(NewFloat(5)).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                      NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(1)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         0,
				CumulatedVacationInterestDays: 0,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(20.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                      NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(0.5)).Sub(NewFloat(2.5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         0,
				CumulatedVacationInterestDays: 0,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 5.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(20.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				RemainingVacationInterestDays:   NewFloat(5),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         0,
				CumulatedVacationInterestDays: 5,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 5.0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(20.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(5)),
				RemainingVacationInterestDays:   NewFloat(5),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(3.0)).Sub(NewFloat(5)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         2,
				CumulatedVacationInterestDays: 5,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(22.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(5)),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(22)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         2,
				CumulatedVacationInterestDays: 5,
				CumulatedSicknessInterestDays: 3,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(22.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(5)),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(3)),
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(22)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         2,
				CumulatedVacationInterestDays: 5,
				CumulatedSicknessInterestDays: 3,
				CumulatedBilledDays:           0,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(22.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(5)),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(3)),
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(22)),
			},
		},
		{
			Period{harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)}, 20},
			PlanConfig{
				Year:                          2015,
				BusinessDays:                  200,
				VacationInterestDays:          25,
				SicknessInterestDays:          5,
				ChildCareInterestDays:         5,
				CumulatedBusinessDays:         2,
				CumulatedVacationInterestDays: 5,
				CumulatedSicknessInterestDays: 3,
				CumulatedBilledDays:           10,
			},
			UserConfig{
				userID:                        "1",
				hasChild:                      false,
				billingDegree:                 0.8,
				workingDegree:                 1.0,
				remainingVacationInterestDays: 0,
			},
			EstimationBillingPeriod{
				ID:                              "10",
				Timeframe:                       harvest.Timeframe{StartDate: harvest.Date(2015, 26, 1, time.Local), EndDate: harvest.Date(2015, 22, 2, time.Local)},
				UserID:                          "1",
				BusinessDays:                    NewFloat(20.0),
				CumulatedBusinessDays:           NewFloat(22.0),
				VacationInterestDays:            NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedVacationInterestDays:   NewFloat(25).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(5)),
				RemainingVacationInterestDays:   NewFloat(0),
				SicknessInterestDays:            NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))),
				CumulatedSicknessInterestDays:   NewFloat(5).Mul(NewFloat(20).Div(NewFloat(200))).Add(NewFloat(3)),
				BilledDays:                      NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)),
				CumulatedBilledDays:             NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Add(NewFloat(10)),
				EffectiveBillingDegree:          NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Div(NewFloat(20)),
				CumulatedEffectiveBillingDegree: NewFloat(20).Sub(NewFloat(3.0)).Mul(NewFloat(0.8)).Add(NewFloat(10)).Div(NewFloat(22)),
			},
		},
	}

	for _, test := range tests {
		period, err := CreateEstimationBillingPeriod(test.period, test.planConfigInput, test.userConfigInput)

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