package harvest

import (
	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

type HarvestProvider struct {
	taskConfig  TaskConfig
	userService *harvest.UserService
	userEntries map[int]HarvestUserEntry
}

func (h *HarvestProvider) Fetch(year int) error {
	var users []*harvest.User
	err := h.userService.All(&users, nil)
	if err != nil {
		return err
	}
	for _, user := range users {
		err = h.FetchUserEntries(user.ID, year)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *HarvestProvider) FetchUserEntries(userId, year int) error {
	var user harvest.User
	err := h.userService.Find(userId, &user, nil)
	if err != nil {
		return err
	}
	dayEntryService := h.userService.DayEntries(&user)
	harvestUserHours := HarvestUserEntryFetcher{year, dayEntryService}
	trackedHours, err := harvestUserHours.FetchUserEntry()
	if err != nil {
		return err
	}
	if h.userEntries == nil {
		h.userEntries = make(map[int]HarvestUserEntry)
	}
	h.userEntries[userId] = trackedHours
	return nil
}

func (h HarvestProvider) TrackedHoursForYear(year int) timetables.TrackedHours {
	converter := HarvestEntryConverter{h.taskConfig}
	var trackedEntries []timetables.TrackingEntry
	for _, userEntry := range h.userEntries {
		trackedEntries = append(trackedEntries, converter.ConvertUserEntry(userEntry)...)
	}
	return timetables.NewTrackedHours(trackedEntries)
}
