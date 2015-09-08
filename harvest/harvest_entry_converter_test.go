package harvest

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

func TestHarvestEntryConverterConvertNonbillable(t *testing.T) {
	converter := HarvestEntryConverter{
		VacationTaskID: 5,
		SicknessTaskID: 8,
	}

	harvestEntries := []*harvest.DayEntry{
		&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
		&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
	}

	entries := converter.ConvertNonbillable(harvestEntries)

	expectedEntries := []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 1, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 2, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 3, time.Local), Type: timetables.NonBillable},
	}

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Logf("Expected convertedEntries to equal\n%q\n\tgot\n%q\n", expectedEntries, entries)
		t.Fail()
	}
}

func TestHarvestEntryConverterConvertBillable(t *testing.T) {
	converter := HarvestEntryConverter{
		VacationTaskID: 10,
		SicknessTaskID: 14,
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
