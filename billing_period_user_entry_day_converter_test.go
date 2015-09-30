package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables/date"
)

func TestBillingPeriodUserEntryDayConverterConvert(t *testing.T) {
	userEntry := BillingPeriodUserEntry{
		ID:                             "17",
		UserID:                         "1",
		Period:                         Period{"1", date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 20},
		VacationInterestHours:          NewRat(8),
		CumulatedVacationInterestHours: NewRat(8),
		SicknessInterestHours:          NewRat(8),
		CumulatedSicknessInterestHours: NewRat(8),
		ChildCareHours:                 NewRat(8),
		CumulatedChildCareHours:        NewRat(8),
		BillableHours:                  NewRat(8),
		CumulatedBillableHours:         NewRat(8),
		OfficeHours:                    NewRat(8),
		CumulatedOfficeHours:           NewRat(8),
		NonbillableHours:               NewRat(8),
		CumulatedNonbillableHours:      NewRat(8),
		BillingDegree:                  NewRat(8),
		CumulatedBillingDegree:         NewRat(8),
	}

	converter := BillingPeriodUserEntryConverter{}

	dayEntries := converter.Convert(userEntry, 1)

	expectedEntries := BillingPeriodDayUserEntry{
		UserID:                        "1",
		Period:                        Period{"1", date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 20},
		VacationInterestDays:          NewRat(1),
		CumulatedVacationInterestDays: NewRat(1),
		SicknessInterestDays:          NewRat(1),
		CumulatedSicknessInterestDays: NewRat(1),
		ChildCareDays:                 NewRat(1),
		CumulatedChildCareDays:        NewRat(1),
		BillableDays:                  NewRat(1),
		CumulatedBillableDays:         NewRat(1),
		OfficeDays:                    NewRat(1),
		CumulatedOfficeDays:           NewRat(1),
		NonbillableDays:               NewRat(1),
		CumulatedNonbillableDays:      NewRat(1),
		BillingDegree:                 NewRat(8),
		CumulatedBillingDegree:        NewRat(8),
	}

	if !reflect.DeepEqual(expectedEntries, dayEntries) {
		t.Logf("Expected day entries to equal\n%+v\n\tgot:\n%+v\n", expectedEntries, dayEntries)
		t.Fail()
	}
}
