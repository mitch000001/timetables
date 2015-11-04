package timetables

type BillingPeriodUserEntryConverter struct {
}

func (b BillingPeriodUserEntryConverter) Convert(userEntry BillingPeriodUserEntry, workingDegree *Rat) BillingPeriodDayUserEntry {
	workingDegreeDayFactor := workingDegree.Mul(NewRat(8))
	return BillingPeriodDayUserEntry{
		UserID:                        userEntry.UserID,
		Period:                        userEntry.Period,
		VacationInterestDays:          userEntry.VacationInterestHours.Div(workingDegreeDayFactor),
		CumulatedVacationInterestDays: userEntry.CumulatedVacationInterestHours.Div(workingDegreeDayFactor),
		SicknessInterestDays:          userEntry.SicknessInterestHours.Div(workingDegreeDayFactor),
		CumulatedSicknessInterestDays: userEntry.CumulatedSicknessInterestHours.Div(workingDegreeDayFactor),
		ChildCareDays:                 userEntry.ChildCareHours.Div(workingDegreeDayFactor),
		CumulatedChildCareDays:        userEntry.CumulatedChildCareHours.Div(workingDegreeDayFactor),
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
	VacationInterestDays          *Rat
	CumulatedVacationInterestDays *Rat
	SicknessInterestDays          *Rat
	CumulatedSicknessInterestDays *Rat
	ChildCareDays                 *Rat
	CumulatedChildCareDays        *Rat
	BillableDays                  *Rat
	CumulatedBillableDays         *Rat
	OfficeDays                    *Rat
	CumulatedOfficeDays           *Rat
	NonbillableDays               *Rat
	CumulatedNonbillableDays      *Rat
	BillingDegree                 *Rat
	CumulatedBillingDegree        *Rat
}
