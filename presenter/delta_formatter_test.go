package presenter

import (
	"reflect"
	"testing"

	"github.com/mitch000001/timetables"
)

func TestBillingDeltaFormatterFormat(t *testing.T) {
	billingDelta := BillingDelta{
		BillableDaysDelta:             Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedBillableDaysDelta:    Delta{timetables.NewRat(8), timetables.NewRat(7)},
		NonbillableDaysDelta:          Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedNonbillableDaysDelta: Delta{timetables.NewRat(8), timetables.NewRat(7)},
		VacationDaysDelta:             Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedVacationDaysDelta:    Delta{timetables.NewRat(8), timetables.NewRat(7)},
		SicknessDaysDelta:             Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedSicknessDaysDelta:    Delta{timetables.NewRat(8), timetables.NewRat(7)},
		ChildCareDaysDelta:            Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedChildCareDaysDelta:   Delta{timetables.NewRat(8), timetables.NewRat(7)},
		OfficeDaysDelta:               Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedOfficeDaysDelta:      Delta{timetables.NewRat(8), timetables.NewRat(7)},
		BillingDegreeDelta:            Delta{timetables.NewRat(8), timetables.NewRat(7)},
		CumulatedBillingDegreeDelta:   Delta{timetables.NewRat(8), timetables.NewRat(7)},
	}

	formatter := BillingDeltaFormatter{}

	formatted := formatter.Format(billingDelta, 2)

	expectedDelta := FormattedBillingDelta{
		BillableDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedBillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
		NonbillableDaysDelta:          FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedNonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
		VacationDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedVacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
		SicknessDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedSicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
		ChildCareDaysDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
		OfficeDaysDelta:               FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedOfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
		BillingDegreeDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
		CumulatedBillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
	}

	if !reflect.DeepEqual(expectedDelta, formatted) {
		t.Logf("Expected formatted delta to equal\n%+v\n\tgot:\n%+v\n", expectedDelta, formatted)
		t.Fail()
	}
}

func TestDeltaFormatterFormat(t *testing.T) {
	tests := []struct {
		precision      int
		delta          Delta
		formattedDelta FormattedDelta
	}{
		{
			precision:      2,
			delta:          Delta{timetables.NewRat(6), timetables.NewRat(8)},
			formattedDelta: FormattedDelta{"6.00", "8.00", "-2.00"},
		},
		{
			precision:      5,
			delta:          Delta{timetables.NewRat(6.123456), timetables.NewRat(8.123454)},
			formattedDelta: FormattedDelta{"6.12346", "8.12345", "-2.00000"},
		},
		{
			precision:      2,
			delta:          Delta{},
			formattedDelta: FormattedDelta{},
		},
		{
			precision:      2,
			delta:          Delta{Tracked: timetables.NewRat(4)},
			formattedDelta: FormattedDelta{},
		},
		{
			precision:      2,
			delta:          Delta{Estimated: timetables.NewRat(4)},
			formattedDelta: FormattedDelta{},
		},
	}

	for _, test := range tests {
		formatter := DeltaFormatter{}

		formatted := formatter.Format(test.delta, test.precision)

		if !reflect.DeepEqual(test.formattedDelta, formatted) {
			t.Logf("Expected formatted Delta to equal\n%+v\n\tgot\n%+v\n", test.formattedDelta, formatted)
			t.Fail()
		}
	}
}
