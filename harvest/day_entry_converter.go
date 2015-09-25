package harvest

import (
	"fmt"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

type TaskConfig struct {
	VacationID  int
	ChildCareID int
	SicknessID  int
}

type DayEntryConverter struct {
	taskConfig TaskConfig
}

func (h DayEntryConverter) ConvertNonbillable(entries []*harvest.DayEntry) []timetables.TrackingEntry {
	var trackingEntries []timetables.TrackingEntry
	for _, entry := range entries {
		trackingEntry := timetables.TrackingEntry{
			UserID:    fmt.Sprintf("%d", entry.UserId),
			Hours:     timetables.NewRat(entry.Hours),
			TrackedAt: timetables.Date(entry.SpentAt.Year(), entry.SpentAt.Month(), entry.SpentAt.Day(), entry.SpentAt.Location()),
		}
		switch entry.TaskId {
		case h.taskConfig.VacationID:
			trackingEntry.Type = timetables.Vacation
		case h.taskConfig.SicknessID:
			trackingEntry.Type = timetables.Sickness
		case h.taskConfig.ChildCareID:
			trackingEntry.Type = timetables.ChildCare
		default:
			trackingEntry.Type = timetables.NonBillable
		}
		trackingEntries = append(trackingEntries, trackingEntry)
	}
	return trackingEntries
}

func (h DayEntryConverter) ConvertBillable(entries []*harvest.DayEntry) []timetables.TrackingEntry {
	var trackingEntries []timetables.TrackingEntry
	for _, entry := range entries {
		trackingEntry := timetables.TrackingEntry{
			UserID:    fmt.Sprintf("%d", entry.UserId),
			Hours:     timetables.NewRat(entry.Hours),
			TrackedAt: timetables.Date(entry.SpentAt.Year(), entry.SpentAt.Month(), entry.SpentAt.Day(), entry.SpentAt.Location()),
			Type:      timetables.Billable,
		}
		trackingEntries = append(trackingEntries, trackingEntry)
	}
	return trackingEntries
}

func (h DayEntryConverter) ConvertUserEntry(userEntry UserEntry) []timetables.TrackingEntry {
	var trackingEntries []timetables.TrackingEntry
	trackingEntries = append(trackingEntries, h.ConvertBillable(userEntry.BillableEntries)...)
	trackingEntries = append(trackingEntries, h.ConvertNonbillable(userEntry.NonbillableEntries)...)
	return trackingEntries
}
