package presenter

type BillingDeltaFormatter struct {
}

func (b BillingDeltaFormatter) Format(delta BillingDelta, precision int) FormattedBillingDelta {
	deltaFormatter := DeltaFormatter{}
	return FormattedBillingDelta{
		BillableDaysDelta:             deltaFormatter.Format(delta.BillableDaysDelta, precision),
		CumulatedBillableDaysDelta:    deltaFormatter.Format(delta.CumulatedBillableDaysDelta, precision),
		NonbillableDaysDelta:          deltaFormatter.Format(delta.NonbillableDaysDelta, precision),
		CumulatedNonbillableDaysDelta: deltaFormatter.Format(delta.CumulatedNonbillableDaysDelta, precision),
		VacationDaysDelta:             deltaFormatter.Format(delta.VacationDaysDelta, precision),
		CumulatedVacationDaysDelta:    deltaFormatter.Format(delta.CumulatedVacationDaysDelta, precision),
		SicknessDaysDelta:             deltaFormatter.Format(delta.SicknessDaysDelta, precision),
		CumulatedSicknessDaysDelta:    deltaFormatter.Format(delta.CumulatedSicknessDaysDelta, precision),
		ChildCareDaysDelta:            deltaFormatter.Format(delta.ChildCareDaysDelta, precision),
		CumulatedChildCareDaysDelta:   deltaFormatter.Format(delta.CumulatedChildCareDaysDelta, precision),
		OfficeDaysDelta:               deltaFormatter.Format(delta.OfficeDaysDelta, precision),
		CumulatedOfficeDaysDelta:      deltaFormatter.Format(delta.CumulatedOfficeDaysDelta, precision),
		BillingDegreeDelta:            deltaFormatter.Format(delta.BillingDegreeDelta, precision),
		CumulatedBillingDegreeDelta:   deltaFormatter.Format(delta.CumulatedBillingDegreeDelta, precision),
	}
}

type DeltaFormatter struct {
}

func (d DeltaFormatter) Format(delta Delta, precision int) FormattedDelta {
	if delta.Tracked == nil || delta.Estimated == nil {
		return FormattedDelta{}
	}
	return FormattedDelta{
		Tracked:   delta.Tracked.FloatString(precision),
		Estimated: delta.Estimated.FloatString(precision),
		Delta:     delta.Delta().FloatString(precision),
	}
}
