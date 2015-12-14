package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables/date"
)

func TestNewForecastBillingPeriod(t *testing.T) {
	period := Period{
		Timeframe:    date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}
	planConfig := PlanConfig{
		Year: 2015,
	}

	billingPeriod := NewForecastBillingPeriod(period, planConfig)

	expectedBillingPeriod := ForecastBillingPeriod{
		ID:          "",
		userEntries: make([]ForecastBillingPeriodUserEntry, 0),
		Period:      period,
		planConfig:  planConfig,
	}

	if !reflect.DeepEqual(expectedBillingPeriod, billingPeriod) {
		t.Logf("Expected billing period to equal\n%q\n\tgot\n%q\n", expectedBillingPeriod, billingPeriod)
		t.Fail()
	}
}

func TestForecastBillingPeriodAddUserEntry(t *testing.T) {
	period := Period{
		Timeframe:    date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}
	planConfig := PlanConfig{
		Year:                  2015,
		BusinessDays:          200,
		VacationInterestDays:  25,
		SicknessInterestDays:  5,
		ChildCareInterestDays: 5,
	}

	billingPeriod := ForecastBillingPeriod{
		ID:          "",
		userEntries: make([]ForecastBillingPeriodUserEntry, 0),
		Period:      period,
		planConfig:  planConfig,
	}

	userConfig := UserConfig{
		userID:                        "1",
		hasChild:                      false,
		billingDegree:                 0.8,
		workingDegree:                 1.0,
		remainingVacationInterestDays: 5.0,
	}

	expectedUserEntry := NewForecastBillingPeriodUserEntry(period, planConfig, userConfig)

	billingPeriod.AddUserEntry(userConfig)

	if len(billingPeriod.userEntries) != 1 {
		t.Logf("Expected one user entry, got %d\n", len(billingPeriod.userEntries))
		t.FailNow()
	}

	if !reflect.DeepEqual(expectedUserEntry, billingPeriod.userEntries[0]) {
		t.Logf("Expected user entry to equal\n%q\n\tgot:\n%q\n", expectedUserEntry, billingPeriod.userEntries[0])
		t.Fail()
	}
}

func TestForecastBillingPeriodUserEntries(t *testing.T) {
	period := Period{
		Timeframe:    date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}
	planConfig := PlanConfig{
		Year:                  2015,
		BusinessDays:          200,
		VacationInterestDays:  25,
		SicknessInterestDays:  5,
		ChildCareInterestDays: 5,
	}

	billingPeriod := ForecastBillingPeriod{
		ID:          "",
		userEntries: make([]ForecastBillingPeriodUserEntry, 0),
		Period:      period,
		planConfig:  planConfig,
	}

	userConfig := UserConfig{
		userID:                        "1",
		hasChild:                      false,
		billingDegree:                 0.8,
		workingDegree:                 1.0,
		remainingVacationInterestDays: 5.0,
	}

	expectedEntries := []ForecastBillingPeriodUserEntry{
		NewForecastBillingPeriodUserEntry(period, planConfig, userConfig),
	}

	billingPeriod.AddUserEntry(userConfig)

	entries := billingPeriod.UserEntries()

	if len(entries) != 1 {
		t.Logf("Expected one user entry, got %d\n", len(entries))
		t.FailNow()
	}

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Logf("Expected user entry to equal\n%q\n\tgot:\n%q\n", expectedEntries, entries)
		t.Fail()
	}
}
