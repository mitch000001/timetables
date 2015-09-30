package timetables

import "github.com/mitch000001/timetables/date"

func NewBillingPeriodUserEntry(period Period, userID string, entries TrackedHours) BillingPeriodUserEntry {
	cumulationTimeframe := date.Timeframe{
		StartDate: date.Date(period.Timeframe.StartDate.Year(), 1, 1, period.Timeframe.StartDate.Location()),
		EndDate:   period.Timeframe.EndDate,
	}
	var billingPeriod = BillingPeriodUserEntry{
		UserID: userID,
		Period: period,
	}
	billingPeriod.ChildCareHours = entries.ChildCareHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedChildCareHours = entries.ChildCareHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.VacationInterestHours = entries.VacationInterestHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedVacationInterestHours = entries.VacationInterestHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.SicknessInterestHours = entries.SicknessInterestHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedSicknessInterestHours = entries.SicknessInterestHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.BillableHours = entries.BillableHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedBillableHours = entries.BillableHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.NonbillableHours = entries.NonBillableRemainderHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedNonbillableHours = entries.NonBillableRemainderHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.OfficeHours = billingPeriod.BillableHours.Add(billingPeriod.NonbillableHours)
	if billingPeriod.OfficeHours.Cmp(NewRat(0)) != 0 {
		billingPeriod.BillingDegree = billingPeriod.BillableHours.Div(billingPeriod.OfficeHours)
	} else {
		billingPeriod.BillingDegree = NewRat(0)
	}
	billingPeriod.CumulatedOfficeHours = billingPeriod.CumulatedBillableHours.Add(billingPeriod.CumulatedNonbillableHours)
	if billingPeriod.CumulatedOfficeHours.Cmp(NewRat(0)) != 0 {
		billingPeriod.CumulatedBillingDegree = billingPeriod.CumulatedBillableHours.Div(billingPeriod.CumulatedOfficeHours)
	} else {
		billingPeriod.CumulatedBillingDegree = NewRat(0)
	}
	return billingPeriod
}

type BillingPeriodUserEntry struct {
	ID                             string
	UserID                         string
	Period                         Period
	VacationInterestHours          *Rat
	CumulatedVacationInterestHours *Rat
	SicknessInterestHours          *Rat
	CumulatedSicknessInterestHours *Rat
	ChildCareHours                 *Rat
	CumulatedChildCareHours        *Rat
	BillableHours                  *Rat
	CumulatedBillableHours         *Rat
	OfficeHours                    *Rat
	CumulatedOfficeHours           *Rat
	NonbillableHours               *Rat
	CumulatedNonbillableHours      *Rat
	BillingDegree                  *Rat
	CumulatedBillingDegree         *Rat
}
