package timetables

func CreateBillingPeriodUserEntry(period Period, userID string, entries TrackedHours) (BillingPeriodUserEntry, interface{}) {
	var billingPeriod = BillingPeriodUserEntry{
		UserID: userID,
		Period: period,
	}
	billingPeriod.VacationInterestHours = entries.VacationInterestHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.SicknessInterestHours = entries.SicknessInterestHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.BilledHours = entries.BillableHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.OverheadAndSlacktimeHours = entries.NonBillableRemainderHoursForUserAndTimeframe(userID, period.Timeframe)
	billingPeriod.OfficeHours = billingPeriod.BilledHours.Add(billingPeriod.OverheadAndSlacktimeHours)
	if billingPeriod.OfficeHours.Cmp(NewFloat(0)) != 0 {
		billingPeriod.BillingDegree = billingPeriod.BilledHours.Div(billingPeriod.OfficeHours)
	} else {
		billingPeriod.BillingDegree = NewFloat(0)
	}
	return billingPeriod, nil
}

type BillingPeriodUserEntry struct {
	ID                        string
	UserID                    string
	Period                    Period
	VacationInterestHours     *Float
	SicknessInterestHours     *Float
	BilledHours               *Float
	OfficeHours               *Float
	OverheadAndSlacktimeHours *Float
	BillingDegree             *Float
}
