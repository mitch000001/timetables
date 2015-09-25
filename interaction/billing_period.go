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
	BillableDays *timetables.Rat
}

type Delta struct {
	Tracked   *timetables.Rat
	Estimated *timetables.Rat
	Delta     *timetables.Rat
}

type TrackingEntryFetcher interface {
	Fetch(year int) ([]timetables.TrackingEntry, error)
}

type TrackingEntryUserFetcher interface {
	FetchForUser(userId string, year int) ([]timetables.TrackingEntry, error)
}
