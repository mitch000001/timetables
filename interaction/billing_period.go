package interaction

import (
	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/date"
)

type BillingPeriod struct {
	StartDate date.ShortDate
	EndDate   date.ShortDate
	Entries   []BillingPeriodEntry
}

type BillingPeriodEntry struct {
	User          User
	TrackedDays   PeriodData
	EstimatedDays PeriodData
}

type User struct {
	FirstName string
	LastName  string
}

type PeriodData struct {
	BillableDays             *timetables.Rat
	CumulatedBillableDays    *timetables.Rat
	NonbillableDays          *timetables.Rat
	CumulatedNonbillableDays *timetables.Rat
	VacationDays             *timetables.Rat
	CumulatedVacationDays    *timetables.Rat
	SicknessDays             *timetables.Rat
	CumulatedSicknessDays    *timetables.Rat
	ChildCareDays            *timetables.Rat
	CumulatedChildCareDays   *timetables.Rat
	OfficeDays               *timetables.Rat
	CumulatedOfficeDays      *timetables.Rat
	BillingDegree            *timetables.Rat
	CumulatedBillingDegree   *timetables.Rat
}

type TrackingEntryFetcher interface {
	Fetch(year int) ([]timetables.TrackingEntry, error)
}

type TrackingEntryUserFetcher interface {
	FetchForUser(userId string, year int) ([]timetables.TrackingEntry, error)
}
