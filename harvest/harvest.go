package harvest

import (
	"fmt"
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

type HarvestEntryConverter struct {
	VacationTaskID int
	SicknessTaskID int
}

func (h HarvestEntryConverter) ConvertNonbillable(entries []*harvest.DayEntry) []timetables.TrackingEntry {
	var trackingEntries []timetables.TrackingEntry
	for _, entry := range entries {
		trackingEntry := timetables.TrackingEntry{
			UserID:    fmt.Sprintf("%d", entry.UserId),
			Hours:     timetables.NewFloat(entry.Hours),
			TrackedAt: timetables.Date(entry.SpentAt.Year(), entry.SpentAt.Month(), entry.SpentAt.Day(), entry.SpentAt.Location()),
		}
		if entry.TaskId == h.VacationTaskID {
			trackingEntry.Type = timetables.Vacation
		} else if entry.TaskId == h.SicknessTaskID {
			trackingEntry.Type = timetables.Sickness
		} else {
			trackingEntry.Type = timetables.NonBillable
		}
		trackingEntries = append(trackingEntries, trackingEntry)
	}
	return trackingEntries
}

func (h HarvestEntryConverter) ConvertBillable(entries []*harvest.DayEntry) []timetables.TrackingEntry {
	var trackingEntries []timetables.TrackingEntry
	for _, entry := range entries {
		trackingEntry := timetables.TrackingEntry{
			UserID:    fmt.Sprintf("%d", entry.UserId),
			Hours:     timetables.NewFloat(entry.Hours),
			TrackedAt: timetables.Date(entry.SpentAt.Year(), entry.SpentAt.Month(), entry.SpentAt.Day(), entry.SpentAt.Location()),
			Type:      timetables.Billable,
		}
		trackingEntries = append(trackingEntries, trackingEntry)
	}
	return trackingEntries
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
