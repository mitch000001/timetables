package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBillingPeriod(t *testing.T) {
	period := Period{
		Timeframe:    NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}

	billingPeriod := NewBillingPeriod(period)

	expectedBillingPeriod := BillingPeriod{
		ID:          "",
		userEntries: make([]BillingPeriodUserEntry, 0),
		period:      period,
	}

	if !reflect.DeepEqual(expectedBillingPeriod, billingPeriod) {
		t.Logf("Expected billing period to equal\n%q\n\tgot\n%q\n", expectedBillingPeriod, billingPeriod)
		t.Fail()
	}
}

func TestBillingPeriodAddUserEntry(t *testing.T) {
	period := Period{
		Timeframe:    NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}
	userId := "1"
	trackedHours := NewTrackedHours([]TrackingEntry{
		TrackingEntry{UserID: "1", Hours: NewFloat(8), Type: Billable, TrackedAt: Date(2015, 1, 5, time.Local)},
	})

	billingPeriod := NewBillingPeriod(period)

	billingPeriod.AddUserEntry(userId, trackedHours)

	expectedUserEntry := NewBillingPeriodUserEntry(period, userId, trackedHours)

	if len(billingPeriod.userEntries) != 1 {
		t.Logf("Expected user entries to have 1 item, got %d\n", len(billingPeriod.userEntries))
		t.FailNow()
	}

	actualEntry := billingPeriod.userEntries[0]

	if !reflect.DeepEqual(expectedUserEntry, actualEntry) {
		t.Logf("Expected user entry to equal\n%q\n\tgot\n%q\n", expectedUserEntry, actualEntry)
		t.Fail()
	}
}
