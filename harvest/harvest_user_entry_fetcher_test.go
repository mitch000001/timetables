package harvest

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/go-harvest/harvest/mock"
)

func TestHarvestEntryFetcherBillableEntries(t *testing.T) {
	dayEntryService := mock.DayEntryService{
		Entries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
		},
		BillableTasks: []int{5},
	}

	harvestFetcher := HarvestUserEntryFetcher{
		dayEntryService: mock.NewDayEntryService(dayEntryService),
	}

	expectedEntries := []*harvest.DayEntry{
		&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
		&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
	}

	trackedEntries, err := harvestFetcher.BillableEntries()

	if err != nil {
		t.Logf("Expected error to be nil, was %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedEntries, trackedEntries) {
		t.Logf("Expected trackingEntries to equal\n%q\n\tgot\n%q\n", expectedEntries, trackedEntries)
		t.Fail()
	}
}

func TestHarvestEntryFetcherNonbillableEntries(t *testing.T) {
	dayEntryService := mock.DayEntryService{
		Entries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
		},
		BillableTasks: []int{5},
	}

	harvestFetcher := HarvestUserEntryFetcher{
		dayEntryService: mock.NewDayEntryService(dayEntryService),
	}

	expectedEntries := []*harvest.DayEntry{
		&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
	}

	trackedEntries, err := harvestFetcher.NonbillableEntries()

	if err != nil {
		t.Logf("Expected error to be nil, was %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedEntries, trackedEntries) {
		t.Logf("Expected trackingEntries to equal\n%q\n\tgot\n%q\n", expectedEntries, trackedEntries)
		t.Fail()
	}
}
