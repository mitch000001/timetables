package timetables

type BillingPeriodUserEntryConverter struct {
}

func (b BillingPeriodUserEntryConverter) Convert(userEntry BillingPeriodUserEntry, workingDegree *Rat) BillingPeriodDayUserEntry {
	workingDegreeDayFactor := workingDegree.Mul(NewRat(8))
	return BillingPeriodDayUserEntry{
		UserID:                   userEntry.UserID,
		Period:                   userEntry.Period,
		VacationDays:             userEntry.VacationHours.Div(workingDegreeDayFactor),
		CumulatedVacationDays:    userEntry.CumulatedVacationHours.Div(workingDegreeDayFactor),
		SicknessDays:             userEntry.SicknessHours.Div(workingDegreeDayFactor),
		CumulatedSicknessDays:    userEntry.CumulatedSicknessHours.Div(workingDegreeDayFactor),
		ChildCareDays:            userEntry.ChildCareHours.Div(workingDegreeDayFactor),
		CumulatedChildCareDays:   userEntry.CumulatedChildCareHours.Div(workingDegreeDayFactor),
		BillableDays:             userEntry.BillableHours.Div(workingDegreeDayFactor),
		CumulatedBillableDays:    userEntry.CumulatedBillableHours.Div(workingDegreeDayFactor),
		OfficeDays:               userEntry.OfficeHours.Div(workingDegreeDayFactor),
		CumulatedOfficeDays:      userEntry.CumulatedOfficeHours.Div(workingDegreeDayFactor),
		NonbillableDays:          userEntry.NonbillableHours.Div(workingDegreeDayFactor),
		CumulatedNonbillableDays: userEntry.CumulatedNonbillableHours.Div(workingDegreeDayFactor),
		BillingDegree:            userEntry.BillingDegree,
		CumulatedBillingDegree:   userEntry.CumulatedBillingDegree,
	}
}

type BillingPeriodDayUserEntry struct {
	UserID                   string
	Period                   Period
	VacationDays             *Rat
	CumulatedVacationDays    *Rat
	SicknessDays             *Rat
	CumulatedSicknessDays    *Rat
	ChildCareDays            *Rat
	CumulatedChildCareDays   *Rat
	BillableDays             *Rat
	CumulatedBillableDays    *Rat
	OfficeDays               *Rat
	CumulatedOfficeDays      *Rat
	NonbillableDays          *Rat
	CumulatedNonbillableDays *Rat
	BillingDegree            *Rat
	CumulatedBillingDegree   *Rat
}
