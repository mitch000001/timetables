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
	trackedHours = timetables.NewTrackedHours(
		h.converter.ConvertBillable(billableEntries),
		h.converter.ConvertNonbillable(nonbillableEntries),
	)
	return trackedHours, nil
}

type HarvestProvider struct {
}

func (h HarvestProvider) TrackedHours() timetables.TrackedHours {
	var trackedHours timetables.TrackedHours
	trackedHours = timetables.NewTrackedHours([]timetables.TrackingEntry{timetables.TrackingEntry{Hours: timetables.NewFloat(8)}}, nil)
	return trackedHours
}
