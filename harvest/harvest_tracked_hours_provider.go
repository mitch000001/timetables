package harvest

import (
	"time"

	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

type HarvestProvider struct {
	taskConfig   TaskConfig
	userService  *harvest.UserService
	trackedHours timetables.TrackedHours
}

func (h *HarvestProvider) Fetch(year int) error {
	h.trackedHours = timetables.NewTrackedHours([]timetables.TrackingEntry{
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 15, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "1", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 20, time.Local), Type: timetables.Vacation},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 17, time.Local), Type: timetables.Billable},
		timetables.TrackingEntry{UserID: "2", Hours: timetables.NewFloat(8), TrackedAt: timetables.Date(2015, 1, 19, time.Local), Type: timetables.Sickness},
	})
	return nil
}

func (h *HarvestProvider) FetchUserEntries(userId, year int) error {
	var user harvest.User
	err := h.userService.Find(userId, &user, nil)
	if err != nil {
		return err
	}
	dayEntryService := h.userService.DayEntries(&user)
	harvestUserHours := NewHarvestUserTrackedHours(year, dayEntryService, h.taskConfig)
	trackedHours, err := harvestUserHours.TrackedHours()
	if err != nil {
		return err
	}
	h.trackedHours = trackedHours
	return nil
}

func (h HarvestProvider) TrackedHoursForYear(year int) timetables.TrackedHours {
	var trackedHours timetables.TrackedHours
	// TODO: Implement this properly
	trackedHours = h.trackedHours
	return trackedHours
}
