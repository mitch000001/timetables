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
		Period:      period,
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

func TestBilligPeriodUserEntries(t *testing.T) {
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

	entries := billingPeriod.UserEntries()

	expectedEntries := []BillingPeriodUserEntry{
		NewBillingPeriodUserEntry(period, userId, trackedHours),
	}

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Logf("Expected user entry to equal\n%q\n\tgot\n%q\n", expectedEntries, entries)
		t.Fail()
	}
}

func TestBillingPeriodMarshalText(t *testing.T) {
	period := BillingPeriod{
		ID:     "17",
		Period: Period{NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 25},
	}

	expected := "{17}:{{{2015-01-01}:{2015-01-25}}:{25}}:[]"

	var err error
	var marshaled []byte

	marshaled, err = period.MarshalText()

	if err != nil {
		t.Logf("Expected error to equal nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expected, string(marshaled)) {
		t.Logf("Expected marshaled value to equal\n%q\n\tgot:\n%q\n", expected, marshaled)
		t.Fail()
	}
}

func TestBillingPeriodUnmarshalText(t *testing.T) {
	expectedPeriod := BillingPeriod{
		ID:     "17",
		Period: Period{NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 25},
	}

	marshaled := "{17}:{{2015-01-01:2015-01-25}:{25}}:[]"

	period := BillingPeriod{}

	err := period.UnmarshalText([]byte(marshaled))

	if err != nil {
		t.Logf("Expected error to equal nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedPeriod, period) {
		t.Logf("Expected unmarshaled period to equal\n%+v\n\tgot\n%+v\n", expectedPeriod, period)
		t.Fail()
	}
}
