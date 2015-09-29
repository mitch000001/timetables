package presenter

import (
	"reflect"
	"testing"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/interaction"
)

func TestDeltaCalculatorCalculate(t *testing.T) {
	tracked := interaction.Days{
		BillableDays:    timetables.NewRat(4),
		NonbillableDays: timetables.NewRat(1),
		VacationDays:    timetables.NewRat(5),
		SicknessDays:    timetables.NewRat(1),
		ChildCareDays:   timetables.NewRat(3),
		OfficeDays:      timetables.NewRat(5),
		BillingDegree:   timetables.NewRat(0.8),
	}

	estimated := interaction.Days{
		BillableDays:    timetables.NewRat(10),
		NonbillableDays: timetables.NewRat(3),
		VacationDays:    timetables.NewRat(2),
		SicknessDays:    timetables.NewRat(1),
		ChildCareDays:   timetables.NewRat(3),
		OfficeDays:      timetables.NewRat(13),
		BillingDegree:   timetables.NewRat(0.6),
	}

	deltaConverter := DeltaConverter{}

	delta := deltaConverter.Convert(tracked, estimated)

	expectedDelta := BillingDelta{
		BillableDaysDelta:    Delta{timetables.NewRat(4), timetables.NewRat(10)},
		NonbillableDaysDelta: Delta{timetables.NewRat(1), timetables.NewRat(3)},
		VacationDaysDelta:    Delta{timetables.NewRat(5), timetables.NewRat(2)},
		SicknessDaysDelta:    Delta{timetables.NewRat(1), timetables.NewRat(1)},
		ChildCareDaysDelta:   Delta{timetables.NewRat(3), timetables.NewRat(3)},
		OfficeDaysDelta:      Delta{timetables.NewRat(5), timetables.NewRat(13)},
		BillingDegreeDelta:   Delta{timetables.NewRat(0.8), timetables.NewRat(0.6)},
	}

	if !reflect.DeepEqual(expectedDelta, delta) {
		t.Logf("Expected delta to equal\n%+v\n\tgot:\n%+v\n", expectedDelta, delta)
		t.Fail()
	}
}
