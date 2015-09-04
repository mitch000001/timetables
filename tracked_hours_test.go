package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTrackedHours(t *testing.T) {
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := NewTrackedHours(billableHours, nonbillableHours)

	if !reflect.DeepEqual(billableHours, trackedHours.billableHours) {
		t.Logf("Expected billableHours to equal\n%q\n\tgot\n%q\n", billableHours, trackedHours.billableHours)
		t.Fail()
	}

	if !reflect.DeepEqual(nonbillableHours, trackedHours.nonbillableHours) {
		t.Logf("Expected nonbillableHours to equal\n%q\n\tgot\n%q\n", nonbillableHours, trackedHours.nonbillableHours)
		t.Fail()
	}
}

func TestTrackedHoursBillableHours(t *testing.T) {
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}

	res := trackedHours.BillableHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursBillableHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.BillableHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursBillableHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 2, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.BillableHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursVacationInterestHours(t *testing.T) {
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.VacationInterestHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursVacationInterestHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.VacationInterestHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursVacationInterestHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 4, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.VacationInterestHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursSicknessInterestHours(t *testing.T) {
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.SicknessInterestHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursSicknessInterestHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.SicknessInterestHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursSicknessInterestHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 4, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.SicknessInterestHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursNonBillableRemainderHours(t *testing.T) {
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.NonBillableRemainderHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursNonBillableRemainderHoursForTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.NonBillableRemainderHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursNonBillableRemainderHoursForUserAndTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
	}
	nonbillableHours := []TrackingEntry{
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewFloat(8), UserID: "2", TrackedAt: Date(2015, 1, 5, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		billableHours:    billableHours,
		nonbillableHours: nonbillableHours,
	}
	var res *Float

	res = trackedHours.NonBillableRemainderHoursForUserAndTimeframe("1", timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}