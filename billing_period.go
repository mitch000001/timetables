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

func (t TrackedEntries) BillableHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.billableEntries {
		hours = hours.Add(NewFloat(entry.Hours))
	}
	return hours
}

func (t TrackedEntries) BillableHoursForTimeframe(timeframe harvest.Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.billableEntries {
		if timeframe.IsInTimeframe(entry.SpentAt) {
			hours = hours.Add(NewFloat(entry.Hours))
		}
	}
	return hours
}

func (t TrackedEntries) VacationInterestHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.VacationTaskID {
			hours = hours.Add(NewFloat(entry.Hours))
		}
	}
	return hours
}

func (t TrackedEntries) VacationInterestHoursForTimeframe(timeframe harvest.Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.VacationTaskID && timeframe.IsInTimeframe(entry.SpentAt) {
			hours = hours.Add(NewFloat(entry.Hours))
		}
	}
	return hours
}

func (t TrackedEntries) SicknessInterestHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.SicknessTaskID {
			hours = hours.Add(NewFloat(entry.Hours))
		}
	}
	return hours
}

func (t TrackedEntries) SicknessInterestHoursForTimeframe(timeframe harvest.Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId == t.billingConfig.SicknessTaskID && timeframe.IsInTimeframe(entry.SpentAt) {
			hours = hours.Add(NewFloat(entry.Hours))
		}
	}
	return hours
}

func (t TrackedEntries) NonBillableRemainderHours() *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId != t.billingConfig.SicknessTaskID && entry.TaskId != t.billingConfig.VacationTaskID {
			hours = hours.Add(NewFloat(entry.Hours))
		}
	}
	return hours
}

func (t TrackedEntries) NonBillableRemainderHoursForTimeframe(timeframe harvest.Timeframe) *Float {
	hours := NewFloat(0)
	for _, entry := range t.nonbillableEntries {
		if entry.TaskId != t.billingConfig.SicknessTaskID && entry.TaskId != t.billingConfig.VacationTaskID {
			if timeframe.IsInTimeframe(entry.SpentAt) {
				hours = hours.Add(NewFloat(entry.Hours))
			}
		}
	}
	return hours
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
