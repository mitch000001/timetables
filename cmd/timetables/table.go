package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

type Task struct {
	*harvest.Task
}

type table struct {
	Timeframe harvest.Timeframe
	Rows      []row
}

func newUserHours(user *harvest.User, timeframe harvest.Timeframe, billable bool) *userHours {
	return &userHours{
		user:      user,
		timeframe: timeframe,
		billable:  billable,
	}
}

type userHours struct {
	user      *harvest.User
	timeframe harvest.Timeframe
	billable  bool
	hours     float64
}

func (u *userHours) getHours() float64 {
	return u.hours
}

type row struct {
	User                   *harvest.User
	Hours                  float64
	Days                   float64
	CumulatedHours         float64
	CumulatedDays          float64
	BillableHours          float64
	BillableDays           float64
	BillingDegree          float64
	CumulatedBillableHours float64
	CumulatedBillableDays  float64
	CumulatedBillingDegree float64
}

func invalidateTableCacheForTimeframe(timeframe harvest.Timeframe, client *harvest.Harvest) {
	log.Printf("Invalidating table cache for timeframe %s\n", timeframe)
	start := time.Now()
	var t *table
	cacheValue := cache.Get(fmt.Sprintf("table:timeframe=%s", timeframe))
	if cacheValue == nil {
		t = &table{Timeframe: timeframe}
	} else {
		t = cacheValue.(*table)
	}
	err := populateTable(t, timeframe, client)
	if err != nil {
		log.Printf("%T: %v\n", err, err)
	}
	cache.Store(fmt.Sprintf("table:timeframe=%s", timeframe), t)
	log.Printf("Table cache invalidated, took %s", time.Since(start))
}

func populateTable(t *table, timeframe harvest.Timeframe, client *harvest.Harvest) error {
	var users []*harvest.User
	err := client.Users.All(&users, nil)
	if err != nil {
		return err
	}
	cumulationTimeframe := harvest.Timeframe{StartDate: harvest.Date(2015, 01, 01, time.Local), EndDate: timeframe.EndDate}
	var rows []row
	var multiErr multiError
	for _, user := range users {
		var hours float64
		var billableHours float64
		var cumulatedHours float64
		var cumulatedBillableHours float64
		hours, err = getHoursForUserAndTimeframe(newUserHours(user, timeframe, false), client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		billableHours, err = getHoursForUserAndTimeframe(newUserHours(user, timeframe, true), client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		// TODO: don't fetch all data since new years eve, use cached values
		cumulatedHours, err = getHoursForUserAndTimeframe(newUserHours(user, cumulationTimeframe, false), client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		// TODO: don't fetch all data since new years eve, use cached values
		cumulatedBillableHours, err = getHoursForUserAndTimeframe(newUserHours(user, cumulationTimeframe, true), client)
		if err != nil {
			multiErr.Add(err)
			continue
		}
		r := row{
			User:                   user,
			Hours:                  hours,
			Days:                   hours / 8,
			CumulatedHours:         cumulatedHours,
			CumulatedDays:          cumulatedHours / 8,
			BillableHours:          billableHours,
			BillableDays:           billableHours / 8,
			CumulatedBillableHours: cumulatedBillableHours,
			CumulatedBillableDays:  cumulatedBillableHours / 8,
			BillingDegree:          (billableHours / hours) * 100,
			CumulatedBillingDegree: (cumulatedBillableHours / cumulatedHours) * 100,
		}
		rows = append(rows, r)
	}
	if len(multiErr) != 0 {
		return multiErr
	}
	t.Rows = rows
	return nil
}

func getHoursForUserAndTimeframe(userHours *userHours, client *harvest.Harvest) (float64, error) {
	key := fmt.Sprintf("user=%d&timeframe=%s&billable=%t", userHours.user.Id(), userHours.timeframe, userHours.billable)
	dayEntries := cache.Get(key)
	var entries []*harvest.DayEntry
	var lastUpdate time.Time
	var cachedEntries []*harvest.DayEntry
	if dayEntries != nil {
		lastUpdate = time.Now()
		cachedEntries = dayEntries.([]*harvest.DayEntry)
	}
	params := harvest.Params{}
	params.ForTimeframe(userHours.timeframe).UpdatedSince(lastUpdate)
	if userHours.billable {
		params.OnlyBillable(userHours.billable)
	}
	err := client.Users.DayEntries(userHours.user).All(&entries, params.Values())
	if err != nil {
		return -1.0, err
	}
	newEntries := make(map[int]*harvest.DayEntry)
	for _, entry := range entries {
		newEntries[entry.ID] = entry
	}
	replacements := make(map[int]int)
	hours := 0.0
	for i, entry := range cachedEntries {
		if newEntry, ok := newEntries[entry.ID]; ok {
			replacements[entry.ID] = i
			hours += newEntry.Hours
		} else {
			hours += entry.Hours
		}
	}
	for id, i := range replacements {
		cachedEntries[i] = newEntries[id]
		delete(newEntries, id)
	}
	if len(newEntries) != 0 {
		for _, entry := range newEntries {
			cachedEntries = append(cachedEntries, entry)
			hours += entry.Hours
		}
	}
	cache.Store(key, cachedEntries)
	return hours, nil
}
