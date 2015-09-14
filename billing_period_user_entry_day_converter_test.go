package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestBillingPeriodUserEntryDayConverterConvert(t *testing.T) {
	userEntry := BillingPeriodUserEntry{
		ID:                             "17",
		UserID:                         "1",
		Period:                         Period{NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 20},
		VacationInterestHours:          NewFloat(8),
		CumulatedVacationInterestHours: NewFloat(8),
		SicknessInterestHours:          NewFloat(8),
		CumulatedSicknessInterestHours: NewFloat(8),
		BillableHours:                  NewFloat(8),
		CumulatedBillableHours:         NewFloat(8),
		OfficeHours:                    NewFloat(8),
		CumulatedOfficeHours:           NewFloat(8),
		NonbillableHours:               NewFloat(8),
		CumulatedNonbillableHours:      NewFloat(8),
		BillingDegree:                  NewFloat(8),
		CumulatedBillingDegree:         NewFloat(8),
	}

	converter := BillingPeriodUserEntryConverter{}

	dayEntries := converter.Convert(userEntry, 1)

	expectedEntries := BillingPeriodDayUserEntry{
		UserID:                        "1",
		Period:                        Period{NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 20},
		VacationInterestDays:          NewFloat(1),
		CumulatedVacationInterestDays: NewFloat(1),
		SicknessInterestDays:          NewFloat(1),
		CumulatedSicknessInterestDays: NewFloat(1),
		BillableDays:                  NewFloat(1),
		CumulatedBillableDays:         NewFloat(1),
		OfficeDays:                    NewFloat(1),
		CumulatedOfficeDays:           NewFloat(1),
		NonbillableDays:               NewFloat(1),
		CumulatedNonbillableDays:      NewFloat(1),
		BillingDegree:                 NewFloat(8),
		CumulatedBillingDegree:        NewFloat(8),
	}

	if !reflect.DeepEqual(expectedEntries, dayEntries) {
		t.Logf("Expected day entries to equal\n%q\n\tgot:\n%q\n", expectedEntries, dayEntries)
		t.Fail()
	}
}
