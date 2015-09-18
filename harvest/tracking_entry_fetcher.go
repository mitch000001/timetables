package harvest

import (
	"github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables"
)

func NewTrackingEntryFetcher(dayEntryService *harvest.DayEntryService, config TaskConfig) TrackingEntryFetcher {
	return TrackingEntryFetcher{
		dayEntryService: dayEntryService,
		taskConfig:      config,
	}
}

type TrackingEntryFetcher struct {
	userService     *harvest.UserService
	dayEntryService *harvest.DayEntryService
	taskConfig      TaskConfig
}

func (t TrackingEntryFetcher) FetchForUser(userId, year int) ([]timetables.TrackingEntry, error) {
	var user harvest.User
	err := t.userService.Find(userId, &user, nil)
	if err != nil {
		return nil, err
	}
	dayEntryService := t.userService.DayEntries(&user)
	userEntryFetcher := UserEntryFetcher{dayEntryService: dayEntryService, year: year}
	entry, err := userEntryFetcher.FetchUserEntry()
	if err != nil {
		return nil, err
	}
	convertedEntries := DayEntryConverter{taskConfig: t.taskConfig}.ConvertUserEntry(entry)
	return convertedEntries, nil
}

func (t TrackingEntryFetcher) Fetch(year int) ([]timetables.TrackingEntry, error) {
	var users []*harvest.User
	err := t.userService.All(&users, nil)
	if err != nil {
		return nil, err
	}
	var entries []timetables.TrackingEntry
	for _, user := range users {
		fetchedEntries, err := t.FetchForUser(user.ID, year)
		if err != nil {
			return nil, err
		}
		entries = append(entries, fetchedEntries...)
	}
	return entries, nil
}
