package harvest

import (
	"fmt"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

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
