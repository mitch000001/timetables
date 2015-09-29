package presenter

import "github.com/mitch000001/timetables/interaction"

type DeltaConverter struct {
}

func (d DeltaConverter) Convert(tracked, estimated interaction.Days) BillingDelta {
	return BillingDelta{
		BillableDaysDelta: Delta{
			tracked.BillableDays,
			estimated.BillableDays,
		},
		CumulatedBillableDaysDelta: Delta{
			tracked.CumulatedBillableDays,
			estimated.CumulatedBillableDays,
		},
		NonbillableDaysDelta: Delta{
			tracked.NonbillableDays,
			estimated.NonbillableDays,
		},
		CumulatedNonbillableDaysDelta: Delta{
			tracked.CumulatedNonbillableDays,
			estimated.CumulatedNonbillableDays,
		},
		VacationDaysDelta: Delta{
			tracked.VacationDays,
			estimated.VacationDays,
		},
		CumulatedVacationDaysDelta: Delta{
			tracked.CumulatedVacationDays,
			estimated.CumulatedVacationDays,
		},
		SicknessDaysDelta: Delta{
			tracked.SicknessDays,
			estimated.SicknessDays,
		},
		CumulatedSicknessDaysDelta: Delta{
			tracked.CumulatedSicknessDays,
			estimated.CumulatedSicknessDays,
		},
		ChildCareDaysDelta: Delta{
			tracked.ChildCareDays,
			estimated.ChildCareDays,
		},
		CumulatedChildCareDaysDelta: Delta{
			tracked.CumulatedChildCareDays,
			estimated.CumulatedChildCareDays,
		},
		OfficeDaysDelta: Delta{
			tracked.OfficeDays,
			estimated.OfficeDays,
		},
		CumulatedOfficeDaysDelta: Delta{
			tracked.CumulatedOfficeDays,
			estimated.CumulatedOfficeDays,
		},
		BillingDegreeDelta: Delta{
			tracked.BillingDegree,
			estimated.BillingDegree,
		},
		CumulatedBillingDegreeDelta: Delta{
			tracked.CumulatedBillingDegree,
			estimated.CumulatedBillingDegree,
		},
	}
}
