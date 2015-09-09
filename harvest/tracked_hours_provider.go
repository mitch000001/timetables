package harvest

import (
	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

func NewTrackedHoursProvider(taskConfig TaskConfig, userService *harvest.UserService) TrackedHoursProvider {
	provider := TrackedHoursProvider{
		taskConfig:  taskConfig,
		userService: userService,
		userEntries: make(map[int]UserEntry),
	}
	return provider
}

type TrackedHoursProvider struct {
	taskConfig  TaskConfig
	userService *harvest.UserService
	userEntries map[int]UserEntry
}

func (h *TrackedHoursProvider) Fetch(year int) error {
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

func (h *TrackedHoursProvider) FetchUserEntries(userId, year int) error {
	var user harvest.User
	err := h.userService.Find(userId, &user, nil)
	if err != nil {
		return err
	}
	dayEntryService := h.userService.DayEntries(&user)
	harvestUserHours := UserEntryFetcher{year, dayEntryService}
	trackedHours, err := harvestUserHours.FetchUserEntry()
	if err != nil {
		return err
	}
	if h.userEntries == nil {
		h.userEntries = make(map[int]UserEntry)
	}
	h.userEntries[userId] = trackedHours
	return nil
}

func (h TrackedHoursProvider) TrackedHoursForYear(year int) timetables.TrackedHours {
	converter := DayEntryConverter{h.taskConfig}
	var trackedEntries []timetables.TrackingEntry
	for _, userEntry := range h.userEntries {
		trackedEntries = append(trackedEntries, converter.ConvertUserEntry(userEntry)...)
	}
	return timetables.NewTrackedHours(trackedEntries)
}
