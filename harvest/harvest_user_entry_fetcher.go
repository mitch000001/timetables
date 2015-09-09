package harvest

import (
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

type HarvestUserEntryFetcher struct {
	year            int
	dayEntryService *harvest.DayEntryService
}

func (h HarvestUserEntryFetcher) BillableEntries() ([]*harvest.DayEntry, error) {
	var entries []*harvest.DayEntry
	var params harvest.Params
	err := h.dayEntryService.All(&entries, params.ForTimeframe(harvest.NewTimeframe(h.year, 1, 1, h.year, 12, 31, time.Local)).Billable(true).Values())
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (h HarvestUserEntryFetcher) NonbillableEntries() ([]*harvest.DayEntry, error) {
	var entries []*harvest.DayEntry
	var params harvest.Params
	err := h.dayEntryService.All(&entries, params.ForTimeframe(harvest.NewTimeframe(h.year, 1, 1, h.year, 12, 31, time.Local)).Billable(false).Values())
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (h HarvestUserEntryFetcher) FetchUserEntry() (HarvestUserEntry, error) {
	var entry HarvestUserEntry
	billable, err := h.BillableEntries()
	if err != nil {
		return entry, err
	}
	nonbillable, err := h.NonbillableEntries()
	if err != nil {
		return entry, err
	}
	entry = HarvestUserEntry{
		BillableEntries:    billable,
		NonbillableEntries: nonbillable,
	}
	return entry, nil
}

type HarvestUserEntry struct {
	BillableEntries    []*harvest.DayEntry
	NonbillableEntries []*harvest.DayEntry
}
