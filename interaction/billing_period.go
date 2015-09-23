package interaction

import "github.com/mitch000001/timetables"

type BillingPeriod struct {
	StartDate timetables.ShortDate
	EndDate   timetables.ShortDate
}

type BillingPeriodEntry struct {
	User        User
	TrackedDays TrackedDays
}

type User struct {
	FirstName string
	LastName  string
}

type TrackedDays struct {
	BillableDays *timetables.Float
}

type Delta struct {
	Tracked   *timetables.Float
	Estimated *timetables.Float
	Delta     *timetables.Float
}

type TrackingEntryFetcher interface {
	Fetch(year int) ([]timetables.TrackingEntry, error)
}

type TrackingEntryUserFetcher interface {
	FetchForUser(userId string, year int) ([]timetables.TrackingEntry, error)
}
