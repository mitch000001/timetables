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
		NonbillableDaysDelta: Delta{
			tracked.NonbillableDays,
			estimated.NonbillableDays,
		},
		VacationDaysDelta: Delta{
			tracked.VacationDays,
			estimated.VacationDays,
		},
		SicknessDaysDelta: Delta{
			tracked.SicknessDays,
			estimated.SicknessDays,
		},
		ChildCareDaysDelta: Delta{
			tracked.ChildCareDays,
			estimated.ChildCareDays,
		},
		OfficeDaysDelta: Delta{
			tracked.OfficeDays,
			estimated.OfficeDays,
		},
		BillingDegreeDelta: Delta{
			tracked.BillingDegree,
			estimated.BillingDegree,
		},
	}
}
