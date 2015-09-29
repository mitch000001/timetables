package presenter

import "github.com/mitch000001/timetables/interaction"

const DefaultDayPrecision = 2

func NewBillingPeriodPresenter(billingPeriod interaction.BillingPeriod) BillingPeriodPresenter {
	return BillingPeriodPresenter{
		model: billingPeriod,
	}
}

type BillingPeriodPresenter struct {
	model interaction.BillingPeriod
}

func (b BillingPeriodPresenter) Present() BillingPeriod {
	return BillingPeriod{
		Entries: []BillingPeriodEntry{},
	}
}

type BillingPeriod struct {
	Entries []BillingPeriodEntry
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
