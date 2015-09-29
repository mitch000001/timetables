package presenter

import (
	"reflect"
	"testing"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/interaction"
)

func TestDeltaConverterConvert(t *testing.T) {
	tracked := interaction.Days{
		BillableDays:             timetables.NewRat(4),
		CumulatedBillableDays:    timetables.NewRat(4),
		NonbillableDays:          timetables.NewRat(1),
		CumulatedNonbillableDays: timetables.NewRat(1),
		VacationDays:             timetables.NewRat(5),
		CumulatedVacationDays:    timetables.NewRat(5),
		SicknessDays:             timetables.NewRat(1),
		CumulatedSicknessDays:    timetables.NewRat(1),
		ChildCareDays:            timetables.NewRat(3),
		CumulatedChildCareDays:   timetables.NewRat(3),
		OfficeDays:               timetables.NewRat(5),
		CumulatedOfficeDays:      timetables.NewRat(5),
		BillingDegree:            timetables.NewRat(0.8),
		CumulatedBillingDegree:   timetables.NewRat(0.8),
	}

	estimated := interaction.Days{
		BillableDays:             timetables.NewRat(10),
		CumulatedBillableDays:    timetables.NewRat(10),
		NonbillableDays:          timetables.NewRat(3),
		CumulatedNonbillableDays: timetables.NewRat(3),
		VacationDays:             timetables.NewRat(2),
		CumulatedVacationDays:    timetables.NewRat(2),
		SicknessDays:             timetables.NewRat(1),
		CumulatedSicknessDays:    timetables.NewRat(1),
		ChildCareDays:            timetables.NewRat(3),
		CumulatedChildCareDays:   timetables.NewRat(3),
		OfficeDays:               timetables.NewRat(13),
		CumulatedOfficeDays:      timetables.NewRat(13),
		BillingDegree:            timetables.NewRat(0.6),
		CumulatedBillingDegree:   timetables.NewRat(0.6),
	}

	deltaConverter := DeltaConverter{}

	delta := deltaConverter.Convert(tracked, estimated)

	expectedDelta := BillingDelta{
		BillableDaysDelta:             Delta{timetables.NewRat(4), timetables.NewRat(10)},
		CumulatedBillableDaysDelta:    Delta{timetables.NewRat(4), timetables.NewRat(10)},
		NonbillableDaysDelta:          Delta{timetables.NewRat(1), timetables.NewRat(3)},
		CumulatedNonbillableDaysDelta: Delta{timetables.NewRat(1), timetables.NewRat(3)},
		VacationDaysDelta:             Delta{timetables.NewRat(5), timetables.NewRat(2)},
		CumulatedVacationDaysDelta:    Delta{timetables.NewRat(5), timetables.NewRat(2)},
		SicknessDaysDelta:             Delta{timetables.NewRat(1), timetables.NewRat(1)},
		CumulatedSicknessDaysDelta:    Delta{timetables.NewRat(1), timetables.NewRat(1)},
		ChildCareDaysDelta:            Delta{timetables.NewRat(3), timetables.NewRat(3)},
		CumulatedChildCareDaysDelta:   Delta{timetables.NewRat(3), timetables.NewRat(3)},
		OfficeDaysDelta:               Delta{timetables.NewRat(5), timetables.NewRat(13)},
		CumulatedOfficeDaysDelta:      Delta{timetables.NewRat(5), timetables.NewRat(13)},
		BillingDegreeDelta:            Delta{timetables.NewRat(0.8), timetables.NewRat(0.6)},
		CumulatedBillingDegreeDelta:   Delta{timetables.NewRat(0.8), timetables.NewRat(0.6)},
	}

	if !reflect.DeepEqual(expectedDelta, delta) {
		t.Logf("Expected delta to equal\n%+v\n\tgot:\n%+v\n", expectedDelta, delta)
		t.Fail()
	}
}
