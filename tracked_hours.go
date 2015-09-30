package timetables

import "github.com/mitch000001/timetables/date"

func NewTrackedHours(entries []TrackingEntry) TrackedHours {
	return TrackedHours{
		entries: entries,
	}
}

type TrackedHours struct {
	entries []TrackingEntry
}

func (t TrackedHours) BillableHours() *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Billable {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) BillableHoursForTimeframe(timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Billable && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) BillableHoursForUserAndTimeframe(userId string, timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Billable && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) VacationInterestHours() *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Vacation {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) VacationInterestHoursForTimeframe(timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Vacation && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) VacationInterestHoursForUserAndTimeframe(userId string, timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Vacation && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) SicknessInterestHours() *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Sickness {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) SicknessInterestHoursForTimeframe(timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Sickness && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) SicknessInterestHoursForUserAndTimeframe(userId string, timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == Sickness && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) ChildCareHours() *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == ChildCare {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) ChildCareHoursForTimeframe(timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == ChildCare && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) ChildCareHoursForUserAndTimeframe(userId string, timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == ChildCare && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) NonBillableRemainderHours() *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == NonBillable {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) NonBillableRemainderHoursForTimeframe(timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == NonBillable {
			if timeframe.IsInTimeframe(entry.TrackedAt) {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) NonBillableRemainderHoursForUserAndTimeframe(userId string, timeframe date.Timeframe) *Rat {
	hours := NewRat(0)
	for _, entry := range t.entries {
		if entry.Type == NonBillable {
			if timeframe.IsInTimeframe(entry.TrackedAt) {
				if entry.UserID == userId {
					hours = hours.Add(entry.Hours)
				}
			}
		}
	}
	return hours
}
