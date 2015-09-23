package presenter

import "github.com/mitch000001/timetables/interaction"

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
}

func NewBillingPeriodEntryPresenter(entry interaction.BillingPeriodEntry) BillingPeriodEntryPresenter {
	return BillingPeriodEntryPresenter{
		model: entry,
	}
}

type BillingPeriodEntryPresenter struct {
	model interaction.BillingPeriodEntry
}

func (b BillingPeriodEntryPresenter) Present() BillingPeriodEntry {
	var entry BillingPeriodEntry
	entry.Name = b.model.User.FirstName
	return entry
}
