package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables/date"
)

func TestNewTrackedHours(t *testing.T) {
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
	}

	trackedHours := NewTrackedHours(entries)

	if !reflect.DeepEqual(entries, trackedHours.entries) {
		t.Logf("Expected entires to equal\n%q\n\tgot\n%q\n", entries, trackedHours.entries)
		t.Fail()
	}
}

func TestTrackedHoursBillableHours(t *testing.T) {
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}

	res := trackedHours.BillableHours()

	expected := NewRat(16)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursBillableHoursForTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.BillableHoursForTimeframe(timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursBillableHoursForUserAndTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "2", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.BillableHoursForUserAndTimeframe("1", timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursVacationHours(t *testing.T) {
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.VacationHours()

	expected := NewRat(16)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursVacationHoursForTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.VacationHoursForTimeframe(timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursVacationHoursForUserAndTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "2", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.VacationHoursForUserAndTimeframe("1", timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursSicknessHours(t *testing.T) {
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.SicknessHours()

	expected := NewRat(16)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursSicknessHoursForTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.SicknessHoursForTimeframe(timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursSicknessHoursForUserAndTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "2", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.SicknessHoursForUserAndTimeframe("1", timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursNonBillableRemainderHours(t *testing.T) {
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.NonBillableRemainderHours()

	expected := NewRat(16)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursNonBillableRemainderHoursForTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.NonBillableRemainderHoursForTimeframe(timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursNonBillableRemainderHoursForUserAndTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: NonBillable},
		TrackingEntry{Hours: NewRat(8), UserID: "2", TrackedAt: date.Date(2015, 1, 5, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.NonBillableRemainderHoursForUserAndTimeframe("1", timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursChildCareHours(t *testing.T) {
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: Vacation},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: Sickness},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: ChildCare},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: ChildCare},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.ChildCareHours()

	expected := NewRat(16)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursChildCareHoursForTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: ChildCare},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: ChildCare},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.ChildCareHoursForTimeframe(timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedHoursChildCareHoursForUserAndTimeframe(t *testing.T) {
	timeframe := date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	entries := []TrackingEntry{
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local), Type: Billable},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2014, 2, 1, time.Local), Type: ChildCare},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local), Type: ChildCare},
		TrackingEntry{Hours: NewRat(8), UserID: "2", TrackedAt: date.Date(2015, 1, 4, time.Local), Type: ChildCare},
		TrackingEntry{Hours: NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local), Type: NonBillable},
	}

	trackedHours := TrackedHours{
		entries: entries,
	}
	var res *Rat

	res = trackedHours.ChildCareHoursForUserAndTimeframe("1", timeframe)

	expected := NewRat(8)

	if expected.Cmp(res) != 0 {
		t.Logf("Expected result to equal\n%+v\n\tgot:\n%+v\n", expected, res)
		t.Fail()
	}
}
