package interaction

import "github.com/mitch000001/timetables"

type BillingPeriod struct {
	StartDate timetables.ShortDate
	EndDate   timetables.ShortDate
	Entries   []BillingPeriodEntry
}

type BillingPeriodEntry struct {
	User          User
	TrackedDays   Days
	EstimatedDays Days
}

type User struct {
	FirstName string
	LastName  string
}

type Days struct {
	BillableDays    *timetables.Rat
	NonbillableDays *timetables.Rat
	VacationDays    *timetables.Rat
	SicknessDays    *timetables.Rat
	ChildCareDays   *timetables.Rat
	OfficeDays      *timetables.Rat
	BillingDegree   *timetables.Rat
}

type TrackingEntryFetcher interface {
	Fetch(year int) ([]timetables.TrackingEntry, error)
}

type TrackingEntryUserFetcher interface {
	FetchForUser(userId string, year int) ([]timetables.TrackingEntry, error)
}
