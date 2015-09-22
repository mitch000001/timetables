package harvest

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

func TestDayEntryConverterConvertNonbillable(t *testing.T) {
	converter := DayEntryConverter{
		taskConfig: TaskConfig{
			VacationID:  5,
			ChildCareID: 11,
			SicknessID:  8,
		},
	}

	harvestEntries := []*harvest.DayEntry{
		&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
		&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
		&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 11, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
	}

	entries := converter.ConvertNonbillable(harvestEntries)

	expectedEntries := []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 1, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 2, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 3, time.Local), Type: timetables.NonBillable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 4, time.Local), Type: timetables.ChildCare},
	}

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Logf("Expected convertedEntries to equal\n%q\n\tgot\n%q\n", expectedEntries, entries)
		t.Fail()
	}
}

func TestDayEntryConverterConvertBillable(t *testing.T) {
	converter := DayEntryConverter{
		taskConfig: TaskConfig{
			VacationID:  10,
			ChildCareID: 20,
			SicknessID:  14,
		},
	}

	harvestEntries := []*harvest.DayEntry{
		&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 3, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 4, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
		&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 9, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
	}

	entries := converter.ConvertBillable(harvestEntries)

	expectedEntries := []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 1, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 2, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 3, time.Local), Type: timetables.Billable},
	}

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Logf("Expected convertedEntries to equal\n%q\n\tgot\n%q\n", expectedEntries, entries)
		t.Fail()
	}
}

func TestDayEntryConverterConvertUserEntry(t *testing.T) {
	converter := DayEntryConverter{
		taskConfig: TaskConfig{
			VacationID:  5,
			ChildCareID: 11,
			SicknessID:  8,
		},
	}

	harvestUserEntry := UserEntry{
		BillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 3, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 4, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 9, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
		},
		NonbillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 4, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
			&harvest.DayEntry{ID: 5, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 5, time.Local)},
			&harvest.DayEntry{ID: 6, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 6, time.Local)},
			&harvest.DayEntry{ID: 7, UserId: 1, Hours: 8, TaskId: 11, SpentAt: harvest.Date(2015, 1, 7, time.Local)},
		},
	}

	entries := converter.ConvertUserEntry(harvestUserEntry)

	expectedEntries := []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 1, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 2, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 3, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 4, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 5, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 6, time.Local), Type: timetables.NonBillable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 7, time.Local), Type: timetables.ChildCare},
	}

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Logf("Expected convertedEntries to equal\n%q\n\tgot\n%q\n", expectedEntries, entries)
		t.Fail()
	}
}
