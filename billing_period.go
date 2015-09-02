package timetables

import "github.com/mitch000001/go-harvest/harvest"

type BillingConfig struct {
	UserID         string
	VacationTaskID int
	SicknessTaskID int
}

type TrackedEntries struct {
	billingConfig      BillingConfig
	billableEntries    []*harvest.DayEntry
	nonbillableEntries []*harvest.DayEntry
}

func (t TrackedEntries) Billable() []*harvest.DayEntry {
	return t.billableEntries
}

func (t TrackedEntries) BillableForTimeframe(timeframe harvest.Timeframe) []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.billableEntries {
		if timeframe.IsInTimeframe(entry.SpentAt) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (t TrackedEntries) VacationInterest() []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.VacationTaskID {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (t TrackedEntries) VacationInterestForTimeframe(timeframe harvest.Timeframe) []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.VacationTaskID && timeframe.IsInTimeframe(entry.SpentAt) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (t TrackedEntries) SicknessInterest() []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.SicknessTaskID {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (t TrackedEntries) SicknessInterestForTimeframe(timeframe harvest.Timeframe) []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.SicknessTaskID && timeframe.IsInTimeframe(entry.SpentAt) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (t TrackedEntries) NonBillableRemainder() []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId != t.billingConfig.SicknessTaskID && entry.TaskId != t.billingConfig.VacationTaskID {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (t TrackedEntries) NonBillableRemainderForTimeframe(timeframe harvest.Timeframe) []*harvest.DayEntry {
	var entries []*harvest.DayEntry
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId != t.billingConfig.SicknessTaskID && entry.TaskId != t.billingConfig.VacationTaskID {
			if timeframe.IsInTimeframe(entry.SpentAt) {
				entries = append(entries, entry)
			}
		}
	}
	return entries
}

func CreateBillingPeriod(period Period, entries TrackedEntries) (BillingPeriod, interface{}) {
	var billingPeriod = BillingPeriod{
		UserID:                entries.billingConfig.UserID,
		Timeframe:             period.Timeframe,
		BusinessDays:          NewFloat(period.BusinessDays),
		CumulatedBusinessDays: NewFloat(period.BusinessDays),
		VacationInterest:      NewFloat(0),
	}
	billingPeriod.CumulatedVacationInterest = billingPeriod.VacationInterest
	billingPeriod.SicknessInterest = NewFloat(0)
	billingPeriod.CumulatedSicknessInterest = billingPeriod.SicknessInterest
	billingPeriod.BilledDays = NewFloat(2)
	billingPeriod.CumulatedBilledDays = billingPeriod.BilledDays
	billingPeriod.OfficeDays = NewFloat(3)
	billingPeriod.CumulatedOfficeDays = billingPeriod.OfficeDays
	billingPeriod.OverheadAndSlacktime = NewFloat(1)
	billingPeriod.CumulatedOverheadAndSlacktime = billingPeriod.OverheadAndSlacktime
	billingPeriod.BillingDegree = billingPeriod.BilledDays.Div(billingPeriod.OfficeDays)
	billingPeriod.CumulatedBillingDegree = billingPeriod.CumulatedBilledDays.Div(billingPeriod.CumulatedOfficeDays)
	return billingPeriod, nil
}

type BillingPeriod struct {
	UserID                        string
	Timeframe                     harvest.Timeframe
	BusinessDays                  *Float
	CumulatedBusinessDays         *Float
	VacationInterest              *Float
	CumulatedVacationInterest     *Float
	SicknessInterest              *Float
	CumulatedSicknessInterest     *Float
	BilledDays                    *Float
	CumulatedBilledDays           *Float
	OfficeDays                    *Float
	CumulatedOfficeDays           *Float
	OverheadAndSlacktime          *Float
	CumulatedOverheadAndSlacktime *Float
	BillingDegree                 *Float
	CumulatedBillingDegree        *Float
}
