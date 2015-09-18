package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestTrackingEntriesSortedByDate(t *testing.T) {
	entries := TrackingEntries{
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 15, time.Local)},
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 25, time.Local)},
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 9, time.Local)},
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 12, time.Local)},
	}

	expectedEntries := TrackingEntries{
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 9, time.Local)},
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 12, time.Local)},
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 15, time.Local)},
		TrackingEntry{UserID: "1", TrackedAt: Date(2015, 1, 25, time.Local)},
	}

	sortedEntries := entries.SortByDate()

	if !reflect.DeepEqual(expectedEntries, sortedEntries) {
		t.Logf("Expected sorted entries to equal\n%+#v\n\tgot:\n%+#v\n", expectedEntries, sortedEntries)
		t.Fail()
	}
}
