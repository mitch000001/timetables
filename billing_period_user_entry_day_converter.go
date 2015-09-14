package timetables

type BillingPeriodUserEntryConverter struct {
}

func (b BillingPeriodUserEntryConverter) Convert(userEntry BillingPeriodUserEntry, workingDegree float64) BillingPeriodDayUserEntry {
	workingDegreeDayFactor := NewFloat(workingDegree).Mul(NewFloat(8))
	return BillingPeriodDayUserEntry{
		UserID:                        userEntry.UserID,
		Period:                        userEntry.Period,
		VacationInterestDays:          userEntry.VacationInterestHours.Div(workingDegreeDayFactor),
		CumulatedVacationInterestDays: userEntry.CumulatedVacationInterestHours.Div(workingDegreeDayFactor),
		SicknessInterestDays:          userEntry.SicknessInterestHours.Div(workingDegreeDayFactor),
		CumulatedSicknessInterestDays: userEntry.CumulatedSicknessInterestHours.Div(workingDegreeDayFactor),
		BillableDays:                  userEntry.BillableHours.Div(workingDegreeDayFactor),
		CumulatedBillableDays:         userEntry.CumulatedBillableHours.Div(workingDegreeDayFactor),
		OfficeDays:                    userEntry.OfficeHours.Div(workingDegreeDayFactor),
		CumulatedOfficeDays:           userEntry.CumulatedOfficeHours.Div(workingDegreeDayFactor),
		NonbillableDays:               userEntry.NonbillableHours.Div(workingDegreeDayFactor),
		CumulatedNonbillableDays:      userEntry.CumulatedNonbillableHours.Div(workingDegreeDayFactor),
		BillingDegree:                 userEntry.BillingDegree,
		CumulatedBillingDegree:        userEntry.CumulatedBillingDegree,
	}
}

type BillingPeriodDayUserEntry struct {
	UserID                        string
	Period                        Period
	VacationInterestDays          *Float
	CumulatedVacationInterestDays *Float
	SicknessInterestDays          *Float
	CumulatedSicknessInterestDays *Float
	BillableDays                  *Float
	CumulatedBillableDays         *Float
	OfficeDays                    *Float
	CumulatedOfficeDays           *Float
	NonbillableDays               *Float
	CumulatedNonbillableDays      *Float
	BillingDegree                 *Float
	CumulatedBillingDegree        *Float
}
