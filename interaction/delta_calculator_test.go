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
		BillableDays:         timetables.NewRat(4),
		NonbillableDays:      timetables.NewRat(1),
		VacationInterestDays: timetables.NewRat(5),
		SicknessInterestDays: timetables.NewRat(1),
	}

	estimated := timetables.EstimationBillingPeriodUserEntry{
		Period:               timetables.Period{timetables.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 16},
		BillableDays:         timetables.NewRat(10),
		NonbillableDays:      timetables.NewRat(3),
		VacationInterestDays: timetables.NewRat(2),
		SicknessInterestDays: timetables.NewRat(1),
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
