package interaction

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables"
)

func TestDeltaCalculatorCalculate(t *testing.T) {
	tracked := timetables.BillingPeriodDayUserEntry{
		Period:               timetables.Period{timetables.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 16},
		BillableDays:         timetables.NewFloat(4),
		NonbillableDays:      timetables.NewFloat(1),
		VacationInterestDays: timetables.NewFloat(5),
		SicknessInterestDays: timetables.NewFloat(1),
	}

	estimated := timetables.EstimationBillingPeriodUserEntry{
		Period:               timetables.Period{timetables.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 16},
		BillableDays:         timetables.NewFloat(10),
		NonbillableDays:      timetables.NewFloat(3),
		VacationInterestDays: timetables.NewFloat(2),
		SicknessInterestDays: timetables.NewFloat(1),
	}

	deltaCalculator := DeltaCalculator{
		tracked:   tracked,
		estimated: estimated,
	}

	delta := deltaCalculator.Calculate()

	expectedDelta := BillingDelta{}

	if !reflect.DeepEqual(expectedDelta, delta) {
		t.Logf("Expected delta to equal\n%+#v\n\tgot:\n%+#v\n", expectedDelta, delta)
		t.Fail()
	}
}
