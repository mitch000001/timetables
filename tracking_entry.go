package timetables

import "sort"

type TrackingEntry struct {
	UserID    string
	Hours     *Float
	TrackedAt ShortDate
	Type      TrackingEntryType
}

type TrackingEntries []TrackingEntry

func (t TrackingEntries) SortByDate() TrackingEntries {
	sortSet := sortedByDate(t)
	sort.Sort(sortSet)
	return TrackingEntries(sortSet)
}

type TrackingEntryType int

const (
	Billable TrackingEntryType = iota
	Vacation
	Sickness
	NonBillable
)

type sortedByDate TrackingEntries

func (s sortedByDate) Less(i, j int) bool { return s[i].TrackedAt.Before(s[j].TrackedAt.Time) }
func (s sortedByDate) Len() int           { return len(s) }
func (s sortedByDate) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
