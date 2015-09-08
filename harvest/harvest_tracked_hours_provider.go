package harvest

import (
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

type HarvestProvider struct {
	taskConfig  TaskConfig
	userService *harvest.UserService
}

func (h HarvestProvider) TrackedHoursForYear(year int) timetables.TrackedHours {
	var trackedHours timetables.TrackedHours
	// TODO: Implement this properly
	trackedHours = timetables.NewTrackedHours([]timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 15, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 20, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 17, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 19, time.Local), Type: timetables.Sickness},
	})
	return trackedHours
}
