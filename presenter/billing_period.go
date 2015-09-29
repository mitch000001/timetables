package presenter

import "github.com/mitch000001/timetables/interaction"

const DefaultDayPrecision = 2
const DefaultDateFormat = "02.01.2006"

func NewBillingPeriodPresenter(billingPeriod interaction.BillingPeriod) BillingPeriodPresenter {
	return BillingPeriodPresenter{
		model:      billingPeriod,
		DateFormat: DefaultDateFormat,
	}
}

type BillingPeriodPresenter struct {
	model      interaction.BillingPeriod
	DateFormat string
}

func (b BillingPeriodPresenter) Present() BillingPeriod {
	var entries []BillingPeriodEntry
	for _, entry := range b.model.Entries {
		entryPresenter := NewBillingPeriodEntryPresenter(entry)
		entries = append(entries, entryPresenter.Present())
	}
	return BillingPeriod{
		StartDate: b.model.StartDate.Format(b.DateFormat),
		EndDate:   b.model.EndDate.Format(b.DateFormat),
		Entries:   entries,
	}
}

type BillingPeriod struct {
	StartDate string
	EndDate   string
	Entries   []BillingPeriodEntry
}

type BillingPeriodEntry struct {
	Name string
	FormattedBillingDelta
}

type BillingDelta struct {
	BillableDaysDelta    Delta
	NonbillableDaysDelta Delta
	VacationDaysDelta    Delta
	SicknessDaysDelta    Delta
	ChildCareDaysDelta   Delta
	OfficeDaysDelta      Delta
	BillingDegreeDelta   Delta
}

type FormattedBillingDelta struct {
	BillableDaysDelta    FormattedDelta
	NonbillableDaysDelta FormattedDelta
	VacationDaysDelta    FormattedDelta
	SicknessDaysDelta    FormattedDelta
	ChildCareDaysDelta   FormattedDelta
	OfficeDaysDelta      FormattedDelta
	BillingDegreeDelta   FormattedDelta
}

func NewBillingPeriodEntryPresenter(entry interaction.BillingPeriodEntry) BillingPeriodEntryPresenter {
	return BillingPeriodEntryPresenter{
		model:        entry,
		DayPrecision: DefaultDayPrecision,
	}
}

type BillingPeriodEntryPresenter struct {
	model        interaction.BillingPeriodEntry
	DayPrecision int
}

func (b BillingPeriodEntryPresenter) Present() BillingPeriodEntry {
	var entry BillingPeriodEntry
	entry.Name = b.model.User.FirstName
	deltaConverter := DeltaConverter{}
	billingDelta := deltaConverter.Convert(b.model.TrackedDays, b.model.EstimatedDays)
	entry.FormattedBillingDelta = BillingDeltaFormatter{}.Format(billingDelta, b.DayPrecision)
	return entry
}
