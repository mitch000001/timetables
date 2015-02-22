package main

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

func TestGetHoursForUserAndTimeframe(t *testing.T) {
	// empty cache
	cache = &InMemoryCache{}
	memCache := cache.(*InMemoryCache)
	entries := []*harvest.DayEntry{
		&harvest.DayEntry{ID: 12, Hours: 8.0},
	}
	client := &harvest.Harvest{
		Users: harvest.NewUserService(&crudProvider{entries: entries}, &userEndpoint{}),
	}
	user := harvest.User{ID: 2}
	timeframe := harvest.From(harvest.Date(2014, 01, 01, time.Local))

	hours, err := getHoursForUserAndTimeframe(&user, timeframe, false, client)

	if err != nil {
		t.Logf("Expected no error, got %T: %v\n", err, err)
		t.Fail()
	}
	if hours != 8.0 {
		t.Logf("Expected hours to equal 8.0, got %f\n", hours)
		t.Fail()
	}
	if len(memCache.store) != 1 {
		t.Logf("Expected cache to have one entry, got %d\n", len(memCache.store))
		t.Fail()
	}
	for k, v := range memCache.store {
		expectedKey := "user=2&timeframe={{2014-01-01 00:00:00 +0000 UTC} {2015-02-22 00:00:00 +0000 UTC}}&billable=false"
		if k != expectedKey {
			t.Logf("Expected key to equal '%s', got '%s'\n", expectedKey, k)
			t.Fail()
		}
		if !reflect.DeepEqual(v, entries) {
			t.Logf("Expected value to equal '%+#v', got '%+#v'\n", entries, v)
			t.Fail()
		}
	}
}

type userEndpoint struct {
	harvest.CrudTogglerEndpoint
}

func (u *userEndpoint) Path() string { return "" }

type dayentryEndpoint struct {
	harvest.CrudEndpoint
	entries []*harvest.DayEntry
}

func (d *dayentryEndpoint) All(data interface{}, params url.Values) error {
	*data.(*[]*harvest.DayEntry) = d.entries
	return nil
}

type crudProvider struct {
	entries []*harvest.DayEntry
	harvest.CrudEndpointProvider
}

func (c *crudProvider) CrudEndpoint(string) harvest.CrudEndpoint {
	return &dayentryEndpoint{entries: c.entries}
}
