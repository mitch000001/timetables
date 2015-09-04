package timetables

import (
	"reflect"
	"testing"
)

func TestPlanPeriods(t *testing.T) {
	billingPeriod := BillingPeriodUserEntry{
		ID:     "1",
		UserID: "1",
	}
	estimationBillingPeriod := EstimationBillingPeriodUserEntry{
		ID:     "32",
		UserID: "1",
	}

	plan := Plan{
		periods: []PlanPeriod{
			PlanPeriod{
				BillingPeriod:           billingPeriod,
				EstimationBillingPeriod: estimationBillingPeriod,
			},
		},
	}
	var periods []PlanPeriod

	periods = plan.Periods()

	if len(periods) != 1 {
		t.Logf("Expected periods to have one item, got %d\n", len(periods))
		t.FailNow()
	}

	if !reflect.DeepEqual(billingPeriod, periods[0].BillingPeriod) {
		t.Logf("Expected billingPeriod to equal\n%q\n\tgot\n%q\n", billingPeriod, periods[0].BillingPeriod)
		t.Fail()
	}

	if !reflect.DeepEqual(estimationBillingPeriod, periods[0].EstimationBillingPeriod) {
		t.Logf("Expected estimationBillingPeriod to equal\n%q\n\tgot\n%q\n", estimationBillingPeriod, periods[0].EstimationBillingPeriod)
		t.Fail()
	}
}
