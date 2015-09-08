package harvest

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/go-harvest/harvest/mock"
	"github.com/mitch000001/timetables"
)

func TestHarvestProviderTrackedHours(t *testing.T) {
	provider := HarvestProvider{}

	trackedHours := provider.TrackedHours()

	expectedBillableHours := timetables.NewFloat(8)

	actualBillableHours := trackedHours.BillableHours()

	if expectedBillableHours.Cmp(actualBillableHours) != 0 {
		t.Logf("Expected billable Hours to equal\n%q\n\tgot\n%q\n", expectedBillableHours.String(), actualBillableHours.String())
		t.Fail()
	}
}

func TestHarvestUserTrackedHoursTrackedHours(t *testing.T) {
	dayEntryService := mock.DayEntryService{
		Entries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 3, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 8, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
			&harvest.DayEntry{ID: 4, UserId: 1, Hours: 8, TaskId: 9, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
			&harvest.DayEntry{ID: 5, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 5, time.Local)},
		},
		BillableTasks: []int{3, 5},
	}

	harvestFetcher := HarvestUserEntryFetcher{
		dayEntryService: mock.NewDayEntryService(dayEntryService),
	}
	harvestConverter := HarvestEntryConverter{
		VacationTaskID: 8,
		SicknessTaskID: 9,
	}

	harvestUserTrackedHours := HarvestUserTrackedHours{
		entryFetcher: harvestFetcher,
		converter:    harvestConverter,
	}

	var trackedHours timetables.TrackedHours
	var err error

	trackedHours, err = harvestUserTrackedHours.TrackedHours()

	if err != nil {
		t.Logf("Expected error to be nil, was %T:%v\n", err, err)
		t.Fail()
	}

	expectedHours := timetables.NewTrackedHours(
		[]timetables.TrackingEntry{
			timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 1, time.Local), Type: timetables.Billable},
			timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 2, time.Local), Type: timetables.Billable},
		},
		[]timetables.TrackingEntry{
			timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 3, time.Local), Type: timetables.Vacation},
			timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 4, time.Local), Type: timetables.Sickness},
			timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 5, time.Local), Type: timetables.NonBillable},
		},
	)

	if !reflect.DeepEqual(expectedHours, trackedHours) {
		t.Logf("Expected trackingEntries to equal\n%q\n\tgot\n%q\n", expectedHours, trackedHours)
		t.Fail()
	}
}
