package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestTrackedEntriesBillableHours(t *testing.T) {
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}

	res := trackedEntries.BillableHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesBillableHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.BillableHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesBillableHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 2, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.BillableHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesVacationInterestHours(t *testing.T) {
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.VacationInterestHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesVacationInterestHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.VacationInterestHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesVacationInterestHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 4, time.Local), Type: Vacation},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.VacationInterestHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesSicknessInterestHours(t *testing.T) {
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.SicknessInterestHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesSicknessInterestHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.SicknessInterestHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesSicknessInterestHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 4, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.SicknessInterestHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesNonBillableRemainderHours(t *testing.T) {
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.NonBillableRemainderHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesNonBillableRemainderHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.NonBillableRemainderHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesNonBillableRemainderHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackedHours{
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackedHours{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local), Type: NonBillable},
		TrackedHours{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 5, time.Local), Type: NonBillable},
	}

	trackedEntries := TrackedEntries{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedEntries.NonBillableRemainderHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}
