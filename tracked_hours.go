package timetables

func NewTrackedHours(entries []TrackingEntry) TrackedHours {
	return TrackedHours{
		entries: entries,
	}
}

type TrackedHours struct {
	entries []TrackingEntry
}

func (t TrackedHours) BillableHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Billable {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) BillableHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Billable && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) BillableHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Billable && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) VacationInterestHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Vacation {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) VacationInterestHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Vacation && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) VacationInterestHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Vacation && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) SicknessInterestHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Sickness {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) SicknessInterestHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Sickness && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) SicknessInterestHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == Sickness && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) ChildCareHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == ChildCare {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) ChildCareHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == ChildCare && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) ChildCareHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == ChildCare && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) NonBillableRemainderHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == NonBillable {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedHours) NonBillableRemainderHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.entries {
		if entry.Type == NonBillable {
			if timeframe.IsInTimeframe(entry.TrackedAt) {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedHours) NonBillableRemainderHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
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
