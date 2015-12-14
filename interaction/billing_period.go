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

func NewBillingPeriodEntry(user timetables.User) BillingPeriodEntry {
	var entry BillingPeriodEntry
	entry.User = User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return entry
}

type BillingPeriodEntry struct {
	User           User
	TrackedDays    PeriodData
	ForecastedDays PeriodData
}

func (b *BillingPeriodEntry) AddTrackingData(billingPeriodEntry timetables.BillingPeriodUserEntry) {
	billingPeriodDayEntry := timetables.BillingPeriodUserEntryConverter{}.Convert(billingPeriodEntry, b.User.WorkingDegree)
	b.TrackedDays = PeriodData{
		BillableDays:             billingPeriodDayEntry.BillableDays,
		CumulatedBillableDays:    billingPeriodDayEntry.CumulatedBillableDays,
		NonbillableDays:          billingPeriodDayEntry.NonbillableDays,
		CumulatedNonbillableDays: billingPeriodDayEntry.CumulatedNonbillableDays,
		VacationDays:             billingPeriodDayEntry.VacationDays,
		CumulatedVacationDays:    billingPeriodDayEntry.CumulatedVacationDays,
		SicknessDays:             billingPeriodDayEntry.SicknessDays,
		CumulatedSicknessDays:    billingPeriodDayEntry.CumulatedSicknessDays,
		ChildCareDays:            billingPeriodDayEntry.ChildCareDays,
		CumulatedChildCareDays:   billingPeriodDayEntry.CumulatedChildCareDays,
		OfficeDays:               billingPeriodDayEntry.OfficeDays,
		CumulatedOfficeDays:      billingPeriodDayEntry.CumulatedOfficeDays,
		BillingDegree:            billingPeriodDayEntry.BillingDegree,
		CumulatedBillingDegree:   billingPeriodDayEntry.CumulatedBillingDegree,
	}
}

func (b *BillingPeriodEntry) AddForecastData(forecastPeriodEntry timetables.ForecastBillingPeriodUserEntry) {
	b.ForecastedDays = PeriodData{
		BillableDays:             forecastPeriodEntry.BillableDays,
		CumulatedBillableDays:    forecastPeriodEntry.CumulatedBillableDays,
		NonbillableDays:          forecastPeriodEntry.NonbillableDays,
		CumulatedNonbillableDays: forecastPeriodEntry.CumulatedNonbillableDays,
		VacationDays:             forecastPeriodEntry.VacationDays.Add(forecastPeriodEntry.RemainingVacationDays),
		CumulatedVacationDays:    forecastPeriodEntry.CumulatedVacationDays,
		SicknessDays:             forecastPeriodEntry.SicknessDays,
		CumulatedSicknessDays:    forecastPeriodEntry.CumulatedSicknessDays,
		ChildCareDays:            forecastPeriodEntry.ChildCareDays,
		CumulatedChildCareDays:   forecastPeriodEntry.CumulatedChildCareDays,
		OfficeDays:               forecastPeriodEntry.OfficeDays,
		CumulatedOfficeDays:      forecastPeriodEntry.CumulatedOfficeDays,
		BillingDegree:            forecastPeriodEntry.EffectiveBillingDegree,
		CumulatedBillingDegree:   forecastPeriodEntry.CumulatedEffectiveBillingDegree,
	}
}

type User struct {
	FirstName     string
	LastName      string
	WorkingDegree *timetables.Rat
	BillingDegree *timetables.Rat
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
