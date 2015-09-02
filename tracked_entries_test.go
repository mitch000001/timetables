package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

func TestTrackedEntriesBillableHours(t *testing.T) {
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
	}
	billingConfig := BillingConfig{}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
	}

	res := trackedEntries.BillableHours()

	expected := NewFloat(16)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesBillableHoursForTimeframe(t *testing.T) {
	timeframe := harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
	}
	billingConfig := BillingConfig{}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
	}
	var res *Float

	res = trackedEntries.BillableHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesVacationInterestHours(t *testing.T) {
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
	}
	billingConfig := BillingConfig{
		VacationTaskID: 12,
	}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
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
	timeframe := harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
	}
	billingConfig := BillingConfig{
		VacationTaskID: 12,
	}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
	}
	var res *Float

	res = trackedEntries.VacationInterestHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesSicknessInterestHours(t *testing.T) {
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	billingConfig := BillingConfig{
		VacationTaskID: 12,
		SicknessTaskID: 14,
	}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
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
	timeframe := harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	billingConfig := BillingConfig{
		VacationTaskID: 12,
		SicknessTaskID: 14,
	}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
	}
	var res *Float

	res = trackedEntries.SicknessInterestHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}

func TestTrackedEntriesNonBillableRemainderHours(t *testing.T) {
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 16, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 16, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	billingConfig := BillingConfig{
		VacationTaskID: 12,
		SicknessTaskID: 14,
	}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
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
	timeframe := harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)
	billableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	nonbillableEntries := []*harvest.DayEntry{
		&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 14, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 16, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{Hours: 8, TaskId: 16, UserId: 2, SpentAt: harvest.Date(2014, 1, 1, time.Local)},
	}
	billingConfig := BillingConfig{
		VacationTaskID: 12,
		SicknessTaskID: 14,
	}

	trackedEntries := TrackedEntries{
		billingConfig:      billingConfig,
		billableEntries:    billableEntries,
		nonbillableEntries: nonbillableEntries,
	}
	var res *Float

	res = trackedEntries.NonBillableRemainderHoursForTimeframe(timeframe)

	expected := NewFloat(8)

	if !reflect.DeepEqual(expected, res) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expected, res)
		t.Fail()
	}
}
