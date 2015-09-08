package harvest

import (
	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

type TaskConfig struct {
	VacationID int
	SicknessID int
}

func NewHarvestUserTrackedHours(year int, dayEntryService *harvest.DayEntryService, taskConfig TaskConfig) HarvestUserTrackedHours {
	entryFetcher := HarvestUserEntryFetcher{year, dayEntryService}
	converter := HarvestEntryConverter{taskConfig}
	return HarvestUserTrackedHours{
		entryFetcher: entryFetcher,
		converter:    converter,
	}
}

type HarvestUserTrackedHours struct {
	entryFetcher HarvestUserEntryFetcher
	converter    HarvestEntryConverter
}

func (h HarvestUserTrackedHours) TrackedHours() (timetables.TrackedHours, error) {
	var trackedHours timetables.TrackedHours
	billableEntries, err := h.entryFetcher.BillableEntries()
	if err != nil {
		return trackedHours, err
	}
	nonbillableEntries, err := h.entryFetcher.NonbillableEntries()
	if err != nil {
		return trackedHours, err
	}
	trackingEntries := make([]timetables.TrackingEntry, 0)
	trackingEntries = append(trackingEntries, h.converter.ConvertBillable(billableEntries)...)
	trackingEntries = append(trackingEntries, h.converter.ConvertNonbillable(nonbillableEntries)...)
	trackedHours = timetables.NewTrackedHours(trackingEntries)
	return trackedHours, nil
}
