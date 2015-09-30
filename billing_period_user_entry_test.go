package timetables

import (
	"fmt"
	"testing"
	"time"

	"github.com/mitch000001/timetables/date"
)

func TestNewBillingPeriodUserEntry(t *testing.T) {
	tests := []struct {
		description    string
		period         Period
		userId         string
		trackedEntries TrackedHours
		output         BillingPeriodUserEntry
	}{
		{
			description: "Two billable and one nonbillabe Day, no sickness, no vacation, no childcare",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(0),
				CumulatedVacationInterestHours: NewRat(0),
				SicknessInterestHours:          NewRat(0),
				CumulatedSicknessInterestHours: NewRat(0),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(16),
				CumulatedBillableHours:         NewRat(16),
				OfficeHours:                    NewRat(24),
				CumulatedOfficeHours:           NewRat(24),
				NonbillableHours:               NewRat(8),
				CumulatedNonbillableHours:      NewRat(8),
				BillingDegree:                  NewRat(16).Div(NewRat(24)),
				CumulatedBillingDegree:         NewRat(16).Div(NewRat(24)),
			},
		},
		{
			description: "One billable and one nonbillabe Day, no sickness, no vacation, no childcare, TrackedHours contain data of other user",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "2", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(0),
				CumulatedVacationInterestHours: NewRat(0),
				SicknessInterestHours:          NewRat(0),
				CumulatedSicknessInterestHours: NewRat(0),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(8),
				CumulatedBillableHours:         NewRat(8),
				OfficeHours:                    NewRat(16),
				CumulatedOfficeHours:           NewRat(16),
				NonbillableHours:               NewRat(8),
				CumulatedNonbillableHours:      NewRat(8),
				BillingDegree:                  NewRat(8).Div(NewRat(16)),
				CumulatedBillingDegree:         NewRat(8).Div(NewRat(16)),
			},
		},
		{
			description: "Two billable and one nonbillabe Day, no sickness, no vacation, no childcare, only six hours a day",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "2",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(6), Type: Billable, UserID: "2", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(6), Type: Billable, UserID: "2", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(6), Type: NonBillable, UserID: "2", TrackedAt: date.Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "2",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(0),
				CumulatedVacationInterestHours: NewRat(0),
				SicknessInterestHours:          NewRat(0),
				CumulatedSicknessInterestHours: NewRat(0),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(12),
				CumulatedBillableHours:         NewRat(12),
				OfficeHours:                    NewRat(18),
				CumulatedOfficeHours:           NewRat(18),
				NonbillableHours:               NewRat(6),
				CumulatedNonbillableHours:      NewRat(6),
				BillingDegree:                  NewRat(12).Div(NewRat(18)),
				CumulatedBillingDegree:         NewRat(12).Div(NewRat(18)),
			},
		},
		{
			description: "Two billable, one nonbillabe and one vacation Day",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Vacation, UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(8),
				CumulatedVacationInterestHours: NewRat(8),
				SicknessInterestHours:          NewRat(0),
				CumulatedSicknessInterestHours: NewRat(0),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(16),
				CumulatedBillableHours:         NewRat(16),
				OfficeHours:                    NewRat(24),
				CumulatedOfficeHours:           NewRat(24),
				NonbillableHours:               NewRat(8),
				CumulatedNonbillableHours:      NewRat(8),
				BillingDegree:                  NewRat(16).Div(NewRat(24)),
				CumulatedBillingDegree:         NewRat(16).Div(NewRat(24)),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness and one vacation Day",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Vacation, UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Sickness, UserID: "1", TrackedAt: date.Date(2015, 1, 5, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(8),
				CumulatedVacationInterestHours: NewRat(8),
				SicknessInterestHours:          NewRat(8),
				CumulatedSicknessInterestHours: NewRat(8),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(16),
				CumulatedBillableHours:         NewRat(16),
				OfficeHours:                    NewRat(24),
				CumulatedOfficeHours:           NewRat(24),
				NonbillableHours:               NewRat(8),
				CumulatedNonbillableHours:      NewRat(8),
				BillingDegree:                  NewRat(16).Div(NewRat(24)),
				CumulatedBillingDegree:         NewRat(16).Div(NewRat(24)),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness, one childCare and one vacation Day",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Vacation, UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Sickness, UserID: "1", TrackedAt: date.Date(2015, 1, 5, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: ChildCare, UserID: "1", TrackedAt: date.Date(2015, 1, 6, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(8),
				CumulatedVacationInterestHours: NewRat(8),
				SicknessInterestHours:          NewRat(8),
				CumulatedSicknessInterestHours: NewRat(8),
				ChildCareHours:                 NewRat(8),
				CumulatedChildCareHours:        NewRat(8),
				BillableHours:                  NewRat(16),
				CumulatedBillableHours:         NewRat(16),
				OfficeHours:                    NewRat(24),
				CumulatedOfficeHours:           NewRat(24),
				NonbillableHours:               NewRat(8),
				CumulatedNonbillableHours:      NewRat(8),
				BillingDegree:                  NewRat(16).Div(NewRat(24)),
				CumulatedBillingDegree:         NewRat(16).Div(NewRat(24)),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness and one vacation Day, but not in timeframe",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 2, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 2, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 2, 3, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Vacation, UserID: "1", TrackedAt: date.Date(2015, 2, 4, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Sickness, UserID: "1", TrackedAt: date.Date(2015, 2, 5, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 1, 1, time.Local), EndDate: date.Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(0),
				CumulatedVacationInterestHours: NewRat(0),
				SicknessInterestHours:          NewRat(0),
				CumulatedSicknessInterestHours: NewRat(0),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(0),
				CumulatedBillableHours:         NewRat(0),
				OfficeHours:                    NewRat(0),
				CumulatedOfficeHours:           NewRat(0),
				NonbillableHours:               NewRat(0),
				CumulatedNonbillableHours:      NewRat(0),
				BillingDegree:                  NewRat(0),
				CumulatedBillingDegree:         NewRat(0),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness and one vacation Day, and the same one month ago",
			period:      Period{date.Timeframe{StartDate: date.Date(2015, 2, 1, time.Local), EndDate: date.Date(2015, 2, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Vacation, UserID: "1", TrackedAt: date.Date(2015, 1, 4, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Sickness, UserID: "1", TrackedAt: date.Date(2015, 1, 5, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 2, 1, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Billable, UserID: "1", TrackedAt: date.Date(2015, 2, 2, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: NonBillable, UserID: "1", TrackedAt: date.Date(2015, 2, 3, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Vacation, UserID: "1", TrackedAt: date.Date(2015, 2, 4, time.Local)},
					TrackingEntry{Hours: NewRat(8), Type: Sickness, UserID: "1", TrackedAt: date.Date(2015, 2, 5, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{date.Timeframe{StartDate: date.Date(2015, 2, 1, time.Local), EndDate: date.Date(2015, 2, 25, time.Local)}, 10},
				VacationInterestHours:          NewRat(8),
				CumulatedVacationInterestHours: NewRat(16),
				SicknessInterestHours:          NewRat(8),
				CumulatedSicknessInterestHours: NewRat(16),
				ChildCareHours:                 NewRat(0),
				CumulatedChildCareHours:        NewRat(0),
				BillableHours:                  NewRat(16),
				CumulatedBillableHours:         NewRat(32),
				OfficeHours:                    NewRat(24),
				CumulatedOfficeHours:           NewRat(48),
				NonbillableHours:               NewRat(8),
				CumulatedNonbillableHours:      NewRat(16),
				BillingDegree:                  NewRat(16).Div(NewRat(24)),
				CumulatedBillingDegree:         NewRat(32).Div(NewRat(48)),
			},
		},
	}
	for _, test := range tests {
		res := NewBillingPeriodUserEntry(test.period, test.userId, test.trackedEntries)

		// TODO: reflect.DeepEqual won't work with big.Rat
		// if !reflect.DeepEqual(test.output, res) {
		if fmt.Sprintf("%+v", test.output) != fmt.Sprintf("%+v", res) {
			t.Logf("TrackedHours: %s\n", test.description)
			t.Logf("Expected billing period to equal\n%+v\n\tgot:\n%+v\n", test.output, res)
			t.Fail()
		}
	}
}
