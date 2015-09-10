package timetables

func NewBillingPeriodUserEntry(period Period, userID string, entries TrackedHours) BillingPeriodUserEntry {
	cumulationTimeframe := Timeframe{
		StartDate: Date(period.Timeframe.StartDate.Year(), 1, 1, period.Timeframe.StartDate.Location()),
		EndDate:   period.Timeframe.EndDate,
	}
	var billingPeriod = BillingPeriodUserEntry{
		UserID: userID,
		Period: period,
	}
	billingPeriod.VacationInterestHours = entries.VacationInterestHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedVacationInterestHours = entries.VacationInterestHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.SicknessInterestHours = entries.SicknessInterestHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedSicknessInterestHours = entries.SicknessInterestHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.BillableHours = entries.BillableHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedBillableHours = entries.BillableHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.NonbillableHours = entries.NonBillableRemainderHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.CumulatedNonbillableHours = entries.NonBillableRemainderHoursForUserAndTimeframe(userID, cumulationTimeframe)
	billingPeriod.OfficeHours = billingPeriod.BillableHours.Add(billingPeriod.NonbillableHours)
	if billingPeriod.OfficeHours.Cmp(NewFloat(0)) != 0 {
		billingPeriod.BillingDegree = billingPeriod.BillableHours.Div(billingPeriod.OfficeHours)
	} else {
		billingPeriod.BillingDegree = NewFloat(0)
	}
	billingPeriod.CumulatedOfficeHours = billingPeriod.CumulatedBillableHours.Add(billingPeriod.CumulatedNonbillableHours)
	if billingPeriod.CumulatedOfficeHours.Cmp(NewFloat(0)) != 0 {
		billingPeriod.CumulatedBillingDegree = billingPeriod.CumulatedBillableHours.Div(billingPeriod.CumulatedOfficeHours)
	} else {
		billingPeriod.CumulatedBillingDegree = NewFloat(0)
	}
	return billingPeriod
}

type BillingPeriodUserEntry struct {
	ID                             string
	UserID                         string
	Period                         Period
	VacationInterestHours          *Float
	CumulatedVacationInterestHours *Float
	SicknessInterestHours          *Float
	CumulatedSicknessInterestHours *Float
	BillableHours                  *Float
	CumulatedBillableHours         *Float
	OfficeHours                    *Float
	CumulatedOfficeHours           *Float
	NonbillableHours               *Float
	CumulatedNonbillableHours      *Float
	BillingDegree                  *Float
	CumulatedBillingDegree         *Float
}
