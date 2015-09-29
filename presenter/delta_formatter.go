package presenter

type BillingDeltaFormatter struct {
}

func (b BillingDeltaFormatter) Format(delta BillingDelta, precision int) FormattedBillingDelta {
	deltaFormatter := DeltaFormatter{}
	return FormattedBillingDelta{
		BillableDaysDelta:    deltaFormatter.Format(delta.BillableDaysDelta, precision),
		NonbillableDaysDelta: deltaFormatter.Format(delta.NonbillableDaysDelta, precision),
		VacationDaysDelta:    deltaFormatter.Format(delta.VacationDaysDelta, precision),
		SicknessDaysDelta:    deltaFormatter.Format(delta.SicknessDaysDelta, precision),
		ChildCareDaysDelta:   deltaFormatter.Format(delta.ChildCareDaysDelta, precision),
		OfficeDaysDelta:      deltaFormatter.Format(delta.OfficeDaysDelta, precision),
		BillingDegreeDelta:   deltaFormatter.Format(delta.BillingDegreeDelta, precision),
	}
}

type DeltaFormatter struct {
}

func (d DeltaFormatter) Format(delta Delta, precision int) FormattedDelta {
	return FormattedDelta{
		Tracked:   delta.Tracked.FloatString(precision),
		Estimated: delta.Estimated.FloatString(precision),
		Delta:     delta.Delta().FloatString(precision),
	}
}
