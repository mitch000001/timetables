package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables/date"
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
			period:      Period{"1", date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
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
				Period:                        Period{"1", date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				UserID:                        "1",
				VacationInterestDays:          NewRat(25).Mul(NewRat(10).Div(NewRat(250))),
				RemainingVacationInterestDays: NewRat(0),
				SicknessInterestDays:          NewRat(5).Mul(NewRat(10).Div(NewRat(250))),
				ChildCareDays:                 NewRat(5).Mul(NewRat(10).Div(NewRat(250))),
				BillableDays:                  NewRat(10).Sub(NewRat(35).Mul(NewRat(10).Div(NewRat(250)))).Mul(NewRat(0.8)),
				NonbillableDays:               NewRat(10).Sub(NewRat(35).Mul(NewRat(10).Div(NewRat(250)))).Mul(NewRat(1).Sub(NewRat(0.8))),
				OfficeDays:                    NewRat(10).Sub(NewRat(35).Mul(NewRat(10).Div(NewRat(250)))),
				EffectiveBillingDegree:        NewRat(10).Sub(NewRat(35).Mul(NewRat(10).Div(NewRat(250)))).Mul(NewRat(0.8)).Div(NewRat(10)),
			},
		},
		{
			description: "Has child, no remaining vacation, 200 business days, 20 in period",
			period:      Period{"1", date.Timeframe{StartDate: date.Date(2015, 1, 26, time.Local), EndDate: date.Date(2015, 2, 22, time.Local)}, 20},
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
				Period:                        Period{"1", date.Timeframe{StartDate: date.Date(2015, 1, 26, time.Local), EndDate: date.Date(2015, 2, 22, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewRat(25).Mul(NewRat(20).Div(NewRat(200))),
				RemainingVacationInterestDays: NewRat(0),
				SicknessInterestDays:          NewRat(5).Mul(NewRat(20).Div(NewRat(200))),
				ChildCareDays:                 NewRat(5).Mul(NewRat(20).Div(NewRat(200))),
				BillableDays:                  NewRat(20).Sub(NewRat(35).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(0.8)),
				NonbillableDays:               NewRat(20).Sub(NewRat(35).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(1).Sub(NewRat(0.8))),
				OfficeDays:                    NewRat(20).Sub(NewRat(35).Mul(NewRat(20).Div(NewRat(200)))),
				EffectiveBillingDegree:        NewRat(20).Sub(NewRat(35).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(0.8)).Div(NewRat(20)),
			},
		},
		{
			description: "Has no child, no remaining vacation, 200 business days, 20 in period",
			period:      Period{"1", date.Timeframe{StartDate: date.Date(2015, 1, 26, time.Local), EndDate: date.Date(2015, 2, 22, time.Local)}, 20},
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
				Period:                        Period{"1", date.Timeframe{StartDate: date.Date(2015, 1, 26, time.Local), EndDate: date.Date(2015, 2, 22, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewRat(25).Mul(NewRat(20).Div(NewRat(200))),
				RemainingVacationInterestDays: NewRat(0),
				SicknessInterestDays:          NewRat(5).Mul(NewRat(20).Div(NewRat(200))),
				ChildCareDays:                 NewRat(0),
				BillableDays:                  NewRat(20).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(0.8)),
				NonbillableDays:               NewRat(20).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(1).Sub(NewRat(0.8))),
				OfficeDays:                    NewRat(20).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))),
				EffectiveBillingDegree:        NewRat(20).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(0.8)).Div(NewRat(20)),
			},
		},
		{
			description: "Has no child, 5 vacation days remaining, 200 business days, 20 in period",
			period:      Period{"1", date.Timeframe{StartDate: date.Date(2015, 26, 1, time.Local), EndDate: date.Date(2015, 22, 2, time.Local)}, 20},
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
				Period:                        Period{"1", date.Timeframe{StartDate: date.Date(2015, 26, 1, time.Local), EndDate: date.Date(2015, 22, 2, time.Local)}, 20},
				UserID:                        "1",
				VacationInterestDays:          NewRat(25).Mul(NewRat(20).Div(NewRat(200))),
				RemainingVacationInterestDays: NewRat(5),
				SicknessInterestDays:          NewRat(5).Mul(NewRat(20).Div(NewRat(200))),
				ChildCareDays:                 NewRat(0),
				BillableDays:                  NewRat(20).Sub(NewRat(5)).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(0.8)),
				NonbillableDays:               NewRat(20).Sub(NewRat(5)).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(1).Sub(NewRat(0.8))),
				OfficeDays:                    NewRat(20).Sub(NewRat(5)).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))),
				EffectiveBillingDegree:        NewRat(20).Sub(NewRat(5)).Sub(NewRat(30).Mul(NewRat(20).Div(NewRat(200)))).Mul(NewRat(0.8)).Div(NewRat(20)),
			},
		},
	}

	for _, test := range tests {
		period := NewEstimationBillingPeriodUserEntry(test.period, test.planConfigInput, test.userConfigInput)

		expectedPeriod := test.output

		compareEstimationBillingPeriodUserEntry(t, expectedPeriod, period)
		if t.Failed() {
			t.Logf("Used configuration: %s\n", test.description)
		}
	}
}

func compareEstimationBillingPeriodUserEntry(t *testing.T, a, b EstimationBillingPeriodUserEntry) {
	if a.ID != b.ID {
		t.Logf("Expected ID to equal %q, got %q\n", a.ID, b.ID)
		t.Fail()
	}
	if !reflect.DeepEqual(a.Period, b.Period) {
		t.Logf("Expected Period to equal\n%+v\n\tgot\n%+v\n", a.Period, b.Period)
		t.Fail()
	}
	if a.UserID != b.UserID {
		t.Logf("Expected UserID to equal %q, got %q\n", a.UserID, b.UserID)
		t.Fail()
	}
	if a.VacationInterestDays.Cmp(b.VacationInterestDays) != 0 {
		t.Logf("Expected VacationInterestDays to equal\n%+v\n\tgot\n%+v\n", a.VacationInterestDays, b.VacationInterestDays)
		t.Fail()
	}
	if a.RemainingVacationInterestDays.Cmp(b.RemainingVacationInterestDays) != 0 {
		t.Logf("Expected RemainingVacationInterestDays to equal\n%+v\n\tgot\n%+v\n", a.RemainingVacationInterestDays, b.RemainingVacationInterestDays)
		t.Fail()
	}
	if a.SicknessInterestDays.Cmp(b.SicknessInterestDays) != 0 {
		t.Logf("Expected SicknessInterestDays to equal\n%+v\n\tgot\n%+v\n", a.SicknessInterestDays, b.SicknessInterestDays)
		t.Fail()
	}
	if a.ChildCareDays.Cmp(b.ChildCareDays) != 0 {
		t.Logf("Expected ChildCareDays to equal\n%+v\n\tgot\n%+v\n", a.ChildCareDays, b.ChildCareDays)
		t.Fail()
	}
	if a.BillableDays.Cmp(b.BillableDays) != 0 {
		t.Logf("Expected BillableDays to equal\n%+v\n\tgot\n%+v\n", a.BillableDays, b.BillableDays)
		t.Fail()
	}
	if a.NonbillableDays.Cmp(b.NonbillableDays) != 0 {
		t.Logf("Expected NonbillableDays to equal\n%+v\n\tgot\n%+v\n", a.NonbillableDays, b.NonbillableDays)
		t.Logf("Expected NonbillableDays to equal\n%+v\n\tgot\n%+v\n", a.NonbillableDays.FloatString(53), b.NonbillableDays.FloatString(53))
		t.Fail()
	}
	if a.OfficeDays.Cmp(b.OfficeDays) != 0 {
		t.Logf("Expected OfficeDays to equal\n%+v\n\tgot\n%+v\n", a.OfficeDays, b.OfficeDays)
		t.Fail()
	}
	if a.EffectiveBillingDegree.Cmp(b.EffectiveBillingDegree) != 0 {
		t.Logf("Expected EffectiveBillingDegree to equal\n%+v\n\tgot\n%+v\n", a.EffectiveBillingDegree, b.EffectiveBillingDegree)
		t.Fail()
	}
}
