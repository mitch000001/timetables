package interaction

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/date"
)

func TestNewBillingPeriodEntry(t *testing.T) {
	user := timetables.User{
		FirstName: "Max",
		LastName:  "Muster",
	}

	expectedEntry := BillingPeriodEntry{
		User: User{
			FirstName: "Max",
			LastName:  "Muster",
		},
	}

	var actual BillingPeriodEntry

	actual = NewBillingPeriodEntry(user)

	if !reflect.DeepEqual(expectedEntry, actual) {
		t.Logf("Expected BillingPeriodEntry to equal\n%+v\n\tgot:\n%+v\n", expectedEntry, actual)
		t.Fail()
	}
}

func TestBillingPeriodEntryAddTrackingData(t *testing.T) {
	billingPeriodUserEntry := timetables.BillingPeriodUserEntry{
		UserID:                         "1",
		Period:                         timetables.Period{"1", date.Timeframe{StartDate: date.Date(2015, 2, 1, time.Local), EndDate: date.Date(2015, 2, 25, time.Local)}, 10},
		VacationInterestHours:          timetables.NewRat(8),
		CumulatedVacationInterestHours: timetables.NewRat(8),
		SicknessInterestHours:          timetables.NewRat(8),
		CumulatedSicknessInterestHours: timetables.NewRat(8),
		ChildCareHours:                 timetables.NewRat(8),
		CumulatedChildCareHours:        timetables.NewRat(8),
		BillableHours:                  timetables.NewRat(8),
		CumulatedBillableHours:         timetables.NewRat(8),
		OfficeHours:                    timetables.NewRat(8),
		CumulatedOfficeHours:           timetables.NewRat(8),
		NonbillableHours:               timetables.NewRat(8),
		CumulatedNonbillableHours:      timetables.NewRat(8),
		BillingDegree:                  timetables.NewRat(8),
		CumulatedBillingDegree:         timetables.NewRat(8),
	}

	expectedEntry := BillingPeriodEntry{
		User: User{WorkingDegree: timetables.NewRat(1)},
		TrackedDays: PeriodData{
			BillableDays:             timetables.NewRat(1),
			CumulatedBillableDays:    timetables.NewRat(1),
			NonbillableDays:          timetables.NewRat(1),
			CumulatedNonbillableDays: timetables.NewRat(1),
			VacationDays:             timetables.NewRat(1),
			CumulatedVacationDays:    timetables.NewRat(1),
			SicknessDays:             timetables.NewRat(1),
			CumulatedSicknessDays:    timetables.NewRat(1),
			ChildCareDays:            timetables.NewRat(1),
			CumulatedChildCareDays:   timetables.NewRat(1),
			OfficeDays:               timetables.NewRat(1),
			CumulatedOfficeDays:      timetables.NewRat(1),
			BillingDegree:            timetables.NewRat(8),
			CumulatedBillingDegree:   timetables.NewRat(8),
		},
	}

	entry := BillingPeriodEntry{
		User: User{WorkingDegree: timetables.NewRat(1)},
	}

	entry.AddTrackingData(billingPeriodUserEntry)

	if !reflect.DeepEqual(expectedEntry, entry) {
		t.Logf("Expected BillingPeriodEntry to equal\n%+v\n\tgot:\n%+v\n", expectedEntry, entry)
		t.Fail()
	}

	// Changing the users WorkingDegree
	entry.User.WorkingDegree = timetables.NewRat(0.5)

	expectedEntry = BillingPeriodEntry{
		User: User{WorkingDegree: timetables.NewRat(0.5)},
		TrackedDays: PeriodData{
			BillableDays:             timetables.NewRat(2),
			CumulatedBillableDays:    timetables.NewRat(2),
			NonbillableDays:          timetables.NewRat(2),
			CumulatedNonbillableDays: timetables.NewRat(2),
			VacationDays:             timetables.NewRat(2),
			CumulatedVacationDays:    timetables.NewRat(2),
			SicknessDays:             timetables.NewRat(2),
			CumulatedSicknessDays:    timetables.NewRat(2),
			ChildCareDays:            timetables.NewRat(2),
			CumulatedChildCareDays:   timetables.NewRat(2),
			OfficeDays:               timetables.NewRat(2),
			CumulatedOfficeDays:      timetables.NewRat(2),
			BillingDegree:            timetables.NewRat(8),
			CumulatedBillingDegree:   timetables.NewRat(8),
		},
	}

	entry.AddTrackingData(billingPeriodUserEntry)

	if !reflect.DeepEqual(expectedEntry, entry) {
		t.Logf("Expected BillingPeriodEntry to equal\n%+v\n\tgot:\n%+v\n", expectedEntry, entry)
		t.Fail()
	}
}

func TestBillingPeriodEntryAddEstimationData(t *testing.T) {
	estimationBillingPeriodUserEntry := timetables.EstimationBillingPeriodUserEntry{
		Period:                          timetables.Period{"1", date.Timeframe{StartDate: date.Date(2015, 26, 1, time.Local), EndDate: date.Date(2015, 22, 2, time.Local)}, 20},
		UserID:                          "1",
		VacationInterestDays:            timetables.NewRat(8),
		CumulatedVacationInterestDays:   timetables.NewRat(8),
		RemainingVacationInterestDays:   timetables.NewRat(8),
		SicknessInterestDays:            timetables.NewRat(8),
		CumulatedSicknessInterestDays:   timetables.NewRat(8),
		ChildCareDays:                   timetables.NewRat(8),
		CumulatedChildCareDays:          timetables.NewRat(8),
		BillableDays:                    timetables.NewRat(8),
		CumulatedBillableDays:           timetables.NewRat(8),
		NonbillableDays:                 timetables.NewRat(8),
		CumulatedNonbillableDays:        timetables.NewRat(8),
		OfficeDays:                      timetables.NewRat(8),
		CumulatedOfficeDays:             timetables.NewRat(8),
		EffectiveBillingDegree:          timetables.NewRat(8),
		CumulatedEffectiveBillingDegree: timetables.NewRat(8),
	}

	expectedEntry := BillingPeriodEntry{
		EstimatedDays: PeriodData{
			BillableDays:             timetables.NewRat(8),
			CumulatedBillableDays:    timetables.NewRat(8),
			NonbillableDays:          timetables.NewRat(8),
			CumulatedNonbillableDays: timetables.NewRat(8),
			VacationDays:             timetables.NewRat(8).Add(timetables.NewRat(8)),
			CumulatedVacationDays:    timetables.NewRat(8),
			SicknessDays:             timetables.NewRat(8),
			CumulatedSicknessDays:    timetables.NewRat(8),
			ChildCareDays:            timetables.NewRat(8),
			CumulatedChildCareDays:   timetables.NewRat(8),
			OfficeDays:               timetables.NewRat(8),
			CumulatedOfficeDays:      timetables.NewRat(8),
			BillingDegree:            timetables.NewRat(8),
			CumulatedBillingDegree:   timetables.NewRat(8),
		},
	}

	entry := BillingPeriodEntry{}

	entry.AddEstimationData(estimationBillingPeriodUserEntry)

	if !reflect.DeepEqual(expectedEntry, entry) {
		t.Logf("Expected BillingPeriodEntry to equal\n%+v\n\tgot:\n%+v\n", expectedEntry, entry)
		t.Fail()
	}
}
