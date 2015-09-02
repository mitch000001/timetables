package timetables

import (
	"math/big"

	"github.com/mitch000001/go-harvest/harvest"
)

func CreateBillingPeriod(period Period, entries TrackedEntries) (BillingPeriod, interface{}) {
	var billingPeriod = BillingPeriod{
		UserID:       entries.billingConfig.UserID,
		Timeframe:    period.Timeframe,
		BusinessDays: NewFloat(period.BusinessDays),
	}
	billingPeriod.VacationInterestHours = entries.VacationInterestHoursForTimeframe(period.Timeframe)
	billingPeriod.SicknessInterestHours = entries.SicknessInterestHoursForTimeframe(period.Timeframe)
	billingPeriod.BilledHours = entries.BillableHoursForTimeframe(period.Timeframe)
	billingPeriod.OverheadAndSlacktimeHours = entries.NonBillableRemainderHoursForTimeframe(period.Timeframe)
	billingPeriod.OfficeHours = billingPeriod.BilledHours.Add(billingPeriod.OverheadAndSlacktimeHours)
	if billingPeriod.OfficeHours.Cmp(big.NewFloat(0)) != 0 {
		billingPeriod.BillingDegree = billingPeriod.BilledHours.Div(billingPeriod.OfficeHours)
	} else {
		billingPeriod.BillingDegree = NewFloat(0)
	}
	return billingPeriod, nil
}

type BillingPeriod struct {
	UserID                    string
	Timeframe                 harvest.Timeframe
	BusinessDays              *Float
	VacationInterestHours     *Float
	SicknessInterestHours     *Float
	BilledHours               *Float
	OfficeHours               *Float
	OverheadAndSlacktimeHours *Float
	BillingDegree             *Float
}
