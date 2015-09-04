package timetables

type BillingConfig struct {
	UserID         string
	VacationTaskID int
	SicknessTaskID int
}

type TrackingEntry struct {
	UserID    string
	Hours     *Float
	TrackedAt ShortDate
	Type      TrackingEntryType
}

type TrackingEntryType int

const (
	Billable TrackingEntryType = iota
	Vacation
	Sickness
	NonBillable
)

type TrackedHoursProvider interface {
	BillableHours() []TrackingEntry
	NonbillableHours() []TrackingEntry
}

type TrackedEntries struct {
	billableHours    []TrackingEntry
	nonbillableHours []TrackingEntry
}

func (t TrackedEntries) BillableHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.billableHours {
		hours = hours.Add(entry.Hours)
	}
	return hours
}

func (t TrackedEntries) BillableHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.billableHours {
		if timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedEntries) BillableHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.billableHours {
		if timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedEntries) VacationInterestHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type == Vacation {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedEntries) VacationInterestHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type == Vacation && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedEntries) VacationInterestHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type == Vacation && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedEntries) SicknessInterestHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type == Sickness {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedEntries) SicknessInterestHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type == Sickness && timeframe.IsInTimeframe(entry.TrackedAt) {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedEntries) SicknessInterestHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type == Sickness && timeframe.IsInTimeframe(entry.TrackedAt) {
			if entry.UserID == userId {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedEntries) NonBillableRemainderHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type != Sickness && entry.Type != Vacation {
			hours = hours.Add(entry.Hours)
		}
	}
	return hours
}

func (t TrackedEntries) NonBillableRemainderHoursForTimeframe(timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type != Sickness && entry.Type != Vacation {
			if timeframe.IsInTimeframe(entry.TrackedAt) {
				hours = hours.Add(entry.Hours)
			}
		}
	}
	return hours
}

func (t TrackedEntries) NonBillableRemainderHoursForUserAndTimeframe(userId string, timeframe Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableHours {
		if entry.Type != Sickness && entry.Type != Vacation {
			if timeframe.IsInTimeframe(entry.TrackedAt) {
				if entry.UserID == userId {
					hours = hours.Add(entry.Hours)
				}
			}
		}
	}
	return hours
}
