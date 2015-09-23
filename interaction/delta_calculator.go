package interaction

import "github.com/mitch000001/timetables"

type DeltaCalculator struct {
	tracked   timetables.BillingPeriodDayUserEntry
	estimated timetables.EstimationBillingPeriodUserEntry
}

func (d DeltaCalculator) Calculate() BillingDelta {
	return BillingDelta{}
}

type BillingDelta struct {
}
