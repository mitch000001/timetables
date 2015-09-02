package timetables

import "github.com/mitch000001/go-harvest/harvest"

	"github.com/mitch000001/go-harvest/harvest"
)

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
