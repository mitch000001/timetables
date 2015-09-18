package harvest

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/go-harvest/harvest/mock"
	"github.com/mitch000001/timetables"
)

func TestNewTrackingEntryFetcher(t *testing.T) {
	dayEntryService := mock.DayEntryEndpoint{
		Entries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 8, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
			&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 11, SpentAt: harvest.Date(2015, 1, 5, time.Local)},
		},
		BillableTasks: []int{5},
		UserId:        1,
	}
	config := TaskConfig{
		VacationID: 11,
		SicknessID: 8,
	}

	fetcher := NewTrackingEntryFetcher(mock.NewDayEntryService(&dayEntryService), config)

	expectedFetcher := TrackingEntryFetcher{
		dayEntryService: mock.NewDayEntryService(&dayEntryService),
		config:          config,
	}

	if !reflect.DeepEqual(expectedFetcher, fetcher) {
		t.Logf("Expected new fetcher to equal\n%+#v\n\tgot\n%+#v\n", expectedFetcher, fetcher)
		t.Fail()
	}
}

func TestTrackingEntryFetcherFetchForUser(t *testing.T) {
	userEndpoint := mock.UserEndpoint{
		Users: []*harvest.User{
			&harvest.User{ID: 1},
			&harvest.User{ID: 2},
		},
		DayEntryEndpoint: mock.DayEntryEndpoint{
			Entries: []*harvest.DayEntry{
				&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
				&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
				&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 8, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
				&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 13, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
				&harvest.DayEntry{ID: 3, UserId: 1, Hours: 8, TaskId: 11, SpentAt: harvest.Date(2015, 1, 5, time.Local)},
			},
			BillableTasks: []int{5},
			UserId:        1,
		},
	}
	config := TaskConfig{
		VacationID: 11,
		SicknessID: 8,
	}

	fetcher := TrackingEntryFetcher{
		userService: mock.NewUserService(&userEndpoint),
		config:      config,
	}

	expectedResult := []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 1, time.Local)},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 2, time.Local)},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Sickness, TrackedAt: timetables.Date(2015, 1, 3, time.Local)},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.NonBillable, TrackedAt: timetables.Date(2015, 1, 4, time.Local)},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Vacation, TrackedAt: timetables.Date(2015, 1, 5, time.Local)},
	}

	var result []timetables.TrackingEntry
	var err error

	result, err = fetcher.FetchForUser(1, 2015)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedResult, result) {
		t.Logf("Expected result to equal\n%+#v\n\tgot:\n%+#v\n", expectedResult, result)
		t.Fail()
	}
}

func TestTrackingEntryFetcherFetch(t *testing.T) {
	userEndpoint := mock.UserEndpoint{
		Users: []*harvest.User{
			&harvest.User{ID: 1},
			&harvest.User{ID: 2},
		},
		DayEntryEndpoint: mock.DayEntryEndpoint{
			Entries: []*harvest.DayEntry{
				&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 15, time.Local)},
				&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 3, SpentAt: harvest.Date(2015, 1, 20, time.Local)},
				&harvest.DayEntry{ID: 3, UserId: 2, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 17, time.Local)},
				&harvest.DayEntry{ID: 4, UserId: 2, Hours: 8, TaskId: 8, SpentAt: harvest.Date(2015, 1, 19, time.Local)},
			},
			BillableTasks: []int{5},
			UserId:        1,
		},
	}
	taskConfig := TaskConfig{
		VacationID: 3,
		SicknessID: 8,
	}
	fetcher := TrackingEntryFetcher{
		config:      taskConfig,
		userService: mock.NewUserService(&userEndpoint),
	}
	year := 2015

	entries, err := fetcher.Fetch(year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	expectedTrackingEntries := []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 15, time.Local)},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Vacation, TrackedAt: timetables.Date(2015, 1, 20, time.Local)},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 17, time.Local)},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), Type: timetables.Sickness, TrackedAt: timetables.Date(2015, 1, 19, time.Local)},
	}

	if !reflect.DeepEqual(expectedTrackingEntries, entries) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackingEntries, entries)
		t.Fail()
	}

	// Added user after fetch
	userEndpoint.Users = append(userEndpoint.Users, &harvest.User{ID: 3})
	userEndpoint.DayEntryEndpoint.Entries = append(userEndpoint.DayEntryEndpoint.Entries, &harvest.DayEntry{ID: 20, UserId: 3, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 23, time.Local)})

	entries, err = fetcher.Fetch(year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	expectedTrackingEntries = []timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 15, time.Local)},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), Type: timetables.Vacation, TrackedAt: timetables.Date(2015, 1, 20, time.Local)},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 17, time.Local)},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), Type: timetables.Sickness, TrackedAt: timetables.Date(2015, 1, 19, time.Local)},
		timetables.TrackingEntry{UserID: "3", Hours: timetables.NewFloat(8), Type: timetables.Billable, TrackedAt: timetables.Date(2015, 1, 23, time.Local)},
	}

	if !reflect.DeepEqual(expectedTrackingEntries, entries) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackingEntries, entries)
		t.Fail()
	}
}
