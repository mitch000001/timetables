package harvest

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/go-harvest/harvest/mock"
	"github.com/mitch000001/timetables"
)

func TestTrackedHoursProviderTrackedHours(t *testing.T) {
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
	provider := TrackedHoursProvider{
		taskConfig:  taskConfig,
		userService: mock.NewUserService(&userEndpoint),
	}

	year := 2015

	trackedHours := provider.TrackedHoursForYear(year)

	expectedTrackedHours := timetables.TrackedHours{}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}

	err := provider.Fetch(year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	trackedHours = provider.TrackedHoursForYear(year)

	expectedTrackedHours = timetables.NewTrackedHours([]timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 15, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 20, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 17, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 19, time.Local), Type: timetables.Sickness},
	})

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}
}

func TestTrackedHoursProviderFetch(t *testing.T) {
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
	provider := TrackedHoursProvider{
		taskConfig:  taskConfig,
		userService: mock.NewUserService(&userEndpoint),
	}
	year := 2015

	trackedHours := provider.userEntries[1]

	expectedTrackedHours := UserEntry{}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}

	err := provider.Fetch(year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	trackedHours = provider.userEntries[1]

	expectedTrackedHours = UserEntry{
		BillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 15, time.Local)},
		},
		NonbillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 3, SpentAt: harvest.Date(2015, 1, 20, time.Local)},
		},
	}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}

	trackedHours = provider.userEntries[2]

	expectedTrackedHours = UserEntry{
		BillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 3, UserId: 2, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 17, time.Local)},
		},
		NonbillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 4, UserId: 2, Hours: 8, TaskId: 8, SpentAt: harvest.Date(2015, 1, 19, time.Local)},
		},
	}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}

	// Added user after fetch
	userEndpoint.Users = append(userEndpoint.Users, &harvest.User{ID: 3})
	userEndpoint.DayEntryEndpoint.Entries = append(userEndpoint.DayEntryEndpoint.Entries, &harvest.DayEntry{ID: 20, UserId: 3, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 23, time.Local)})

	err = provider.Fetch(year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	trackedHours = provider.userEntries[3]

	expectedTrackedHours = UserEntry{
		BillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 20, UserId: 3, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 23, time.Local)},
		},
		NonbillableEntries: []*harvest.DayEntry{},
	}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}
}

func TestTrackedHoursProviderFetchUserEntries(t *testing.T) {
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
	provider := TrackedHoursProvider{
		taskConfig:  taskConfig,
		userService: mock.NewUserService(&userEndpoint),
	}
	year := 2015
	userId := 1

	err := provider.FetchUserEntries(userId, year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	trackedHours := provider.userEntries[userId]

	expectedTrackedHours := UserEntry{
		BillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 1, UserId: 1, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 15, time.Local)},
		},
		NonbillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 2, UserId: 1, Hours: 8, TaskId: 3, SpentAt: harvest.Date(2015, 1, 20, time.Local)},
		},
	}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}

	// for another userId
	userId = 2
	userEndpoint.DayEntryEndpoint.UserId = userId

	err = provider.FetchUserEntries(userId, year)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	trackedHours = provider.userEntries[userId]

	expectedTrackedHours = UserEntry{
		BillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 3, UserId: 2, Hours: 8, TaskId: 5, SpentAt: harvest.Date(2015, 1, 17, time.Local)},
		},
		NonbillableEntries: []*harvest.DayEntry{
			&harvest.DayEntry{ID: 4, UserId: 2, Hours: 8, TaskId: 8, SpentAt: harvest.Date(2015, 1, 19, time.Local)},
		},
	}

	if !reflect.DeepEqual(expectedTrackedHours, trackedHours) {
		t.Logf("Expected tracked hours to equal\n%q\n\tgot\n%q\n", expectedTrackedHours, trackedHours)
		t.Fail()
	}
}
