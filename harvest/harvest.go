package harvest

import (
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

type HarvestUserEntryFetcher struct {
	dayEntryService *harvest.DayEntryService
}

func (h HarvestUserEntryFetcher) BillableEntries() ([]*harvest.DayEntry, error) {
	var entries []*harvest.DayEntry
	var params harvest.Params
	err := h.dayEntryService.All(&entries, params.ForTimeframe(harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)).Billable(true).Values())
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (h HarvestUserEntryFetcher) NonbillableEntries() ([]*harvest.DayEntry, error) {
	var entries []*harvest.DayEntry
	var params harvest.Params
	err := h.dayEntryService.All(&entries, params.ForTimeframe(harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)).Billable(false).Values())
	if err != nil {
		return nil, err
	}
	return entries, nil
}

type HarvestProvider struct {
}

func (h HarvestProvider) TrackedHours() timetables.TrackedHours {
	var trackedHours timetables.TrackedHours
	trackedHours = timetables.NewTrackedHours([]timetables.TrackingEntry{timetables.TrackingEntry{Hours: timetables.NewFloat(8)}}, nil)
	return trackedHours
}
