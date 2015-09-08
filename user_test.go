package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	harvestID := 2
	firstName := "Max"
	lastName := "Forn"
	email := "max.forn@example.com"

	user := CreateUser(harvestID, firstName, lastName, email)

	if user.ID == "" {
		t.Logf("Expected ID to be set properly\n")
		t.Fail()
	}

	if user.HarvestID != harvestID {
		t.Logf("Expected HarvestID to equal %d, got %d\n", harvestID, user.HarvestID)
		t.Fail()
	}

	if user.FirstName != firstName {
		t.Logf("Expected FirstName to equal %q, got %q\n", firstName, user.FirstName)
		t.Fail()
	}

	if user.LastName != lastName {
		t.Logf("Expected LastName to equal %q, got %q\n", lastName, user.LastName)
		t.Fail()
	}

	if user.Email != email {
		t.Logf("Expected Email to equal %q, got %q\n", email, user.Email)
		t.Fail()
	}
}

func TestUserBillableHoursForTimeframe(t *testing.T) {
	user := User{
		ID: "1",
		TrackedHours: TrackedHours{
			billableHours: []TrackingEntry{
				TrackingEntry{Hours: NewFloat(8), UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local), Type: Billable},
			},
		},
	}
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local)

	billableHours := user.BillableHoursForTimeframe(timeframe)

	expectedHours := NewFloat(8)

	if !reflect.DeepEqual(expectedHours, billableHours) {
		t.Logf("Expected result to equal\n%#v\n\tgot:\n%#v\n", expectedHours, billableHours)
		t.Fail()
	}
}
