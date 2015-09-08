package timetables

func CreateBillingPeriodUserEntry(period Period, user User) (BillingPeriodUserEntry, interface{}) {
	var billingPeriod = BillingPeriodUserEntry{
		UserID: user.ID,
		Period: period,
	}
	billingPeriod.VacationInterestHours = user.TrackedHours.VacationInterestHoursForUserAndTimeframe(user.ID, period.Timeframe)
	billingPeriod.SicknessInterestHours = user.TrackedHours.SicknessInterestHoursForUserAndTimeframe(user.ID, period.Timeframe)
	billingPeriod.BilledHours = user.TrackedHours.BillableHoursForUserAndTimeframe(user.ID, period.Timeframe)
	billingPeriod.OverheadAndSlacktimeHours = user.TrackedHours.NonBillableRemainderHoursForUserAndTimeframe(user.ID, period.Timeframe)
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
