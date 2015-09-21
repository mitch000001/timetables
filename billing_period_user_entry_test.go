package timetables

import (
	"fmt"
	"testing"
	"time"
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
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(0),
				CumulatedVacationInterestHours: NewFloat(0),
				SicknessInterestHours:          NewFloat(0),
				CumulatedSicknessInterestHours: NewFloat(0),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(16),
				CumulatedBillableHours:         NewFloat(16),
				OfficeHours:                    NewFloat(24),
				CumulatedOfficeHours:           NewFloat(24),
				NonbillableHours:               NewFloat(8),
				CumulatedNonbillableHours:      NewFloat(8),
				BillingDegree:                  NewFloat(16).Div(NewFloat(24)),
				CumulatedBillingDegree:         NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			description: "One billable and one nonbillabe Day, no sickness, no vacation, no childcare, TrackedHours contain data of other user",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "2", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(0),
				CumulatedVacationInterestHours: NewFloat(0),
				SicknessInterestHours:          NewFloat(0),
				CumulatedSicknessInterestHours: NewFloat(0),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(8),
				CumulatedBillableHours:         NewFloat(8),
				OfficeHours:                    NewFloat(16),
				CumulatedOfficeHours:           NewFloat(16),
				NonbillableHours:               NewFloat(8),
				CumulatedNonbillableHours:      NewFloat(8),
				BillingDegree:                  NewFloat(8).Div(NewFloat(16)),
				CumulatedBillingDegree:         NewFloat(8).Div(NewFloat(16)),
			},
		},
		{
			description: "Two billable and one nonbillabe Day, no sickness, no vacation, no childcare, only six hours a day",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "2",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(6), Type: Billable, UserID: "2", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(6), Type: Billable, UserID: "2", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(6), Type: NonBillable, UserID: "2", TrackedAt: Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "2",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(0),
				CumulatedVacationInterestHours: NewFloat(0),
				SicknessInterestHours:          NewFloat(0),
				CumulatedSicknessInterestHours: NewFloat(0),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(12),
				CumulatedBillableHours:         NewFloat(12),
				OfficeHours:                    NewFloat(18),
				CumulatedOfficeHours:           NewFloat(18),
				NonbillableHours:               NewFloat(6),
				CumulatedNonbillableHours:      NewFloat(6),
				BillingDegree:                  NewFloat(12).Div(NewFloat(18)),
				CumulatedBillingDegree:         NewFloat(12).Div(NewFloat(18)),
			},
		},
		{
			description: "Two billable, one nonbillabe and one vacation Day",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(8),
				CumulatedVacationInterestHours: NewFloat(8),
				SicknessInterestHours:          NewFloat(0),
				CumulatedSicknessInterestHours: NewFloat(0),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(16),
				CumulatedBillableHours:         NewFloat(16),
				OfficeHours:                    NewFloat(24),
				CumulatedOfficeHours:           NewFloat(24),
				NonbillableHours:               NewFloat(8),
				CumulatedNonbillableHours:      NewFloat(8),
				BillingDegree:                  NewFloat(16).Div(NewFloat(24)),
				CumulatedBillingDegree:         NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness and one vacation Day",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Sickness, UserID: "1", TrackedAt: Date(2015, 1, 5, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(8),
				CumulatedVacationInterestHours: NewFloat(8),
				SicknessInterestHours:          NewFloat(8),
				CumulatedSicknessInterestHours: NewFloat(8),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(16),
				CumulatedBillableHours:         NewFloat(16),
				OfficeHours:                    NewFloat(24),
				CumulatedOfficeHours:           NewFloat(24),
				NonbillableHours:               NewFloat(8),
				CumulatedNonbillableHours:      NewFloat(8),
				BillingDegree:                  NewFloat(16).Div(NewFloat(24)),
				CumulatedBillingDegree:         NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness, one childCare and one vacation Day",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Sickness, UserID: "1", TrackedAt: Date(2015, 1, 5, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: ChildCare, UserID: "1", TrackedAt: Date(2015, 1, 6, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(8),
				CumulatedVacationInterestHours: NewFloat(8),
				SicknessInterestHours:          NewFloat(8),
				CumulatedSicknessInterestHours: NewFloat(8),
				ChildCareHours:                 NewFloat(8),
				CumulatedChildCareHours:        NewFloat(8),
				BillableHours:                  NewFloat(16),
				CumulatedBillableHours:         NewFloat(16),
				OfficeHours:                    NewFloat(24),
				CumulatedOfficeHours:           NewFloat(24),
				NonbillableHours:               NewFloat(8),
				CumulatedNonbillableHours:      NewFloat(8),
				BillingDegree:                  NewFloat(16).Div(NewFloat(24)),
				CumulatedBillingDegree:         NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness and one vacation Day, but not in timeframe",
			period:      Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 2, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 2, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 2, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 2, 4, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Sickness, UserID: "1", TrackedAt: Date(2015, 2, 5, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(0),
				CumulatedVacationInterestHours: NewFloat(0),
				SicknessInterestHours:          NewFloat(0),
				CumulatedSicknessInterestHours: NewFloat(0),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(0),
				CumulatedBillableHours:         NewFloat(0),
				OfficeHours:                    NewFloat(0),
				CumulatedOfficeHours:           NewFloat(0),
				NonbillableHours:               NewFloat(0),
				CumulatedNonbillableHours:      NewFloat(0),
				BillingDegree:                  NewFloat(0),
				CumulatedBillingDegree:         NewFloat(0),
			},
		},
		{
			description: "Two billable, one nonbillabe, one sickness and one vacation Day, and the same one month ago",
			period:      Period{Timeframe{StartDate: Date(2015, 2, 1, time.Local), EndDate: Date(2015, 2, 25, time.Local)}, 10},
			userId:      "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Sickness, UserID: "1", TrackedAt: Date(2015, 1, 5, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 2, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 2, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 2, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 2, 4, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Sickness, UserID: "1", TrackedAt: Date(2015, 2, 5, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                         "1",
				Period:                         Period{Timeframe{StartDate: Date(2015, 2, 1, time.Local), EndDate: Date(2015, 2, 25, time.Local)}, 10},
				VacationInterestHours:          NewFloat(8),
				CumulatedVacationInterestHours: NewFloat(16),
				SicknessInterestHours:          NewFloat(8),
				CumulatedSicknessInterestHours: NewFloat(16),
				ChildCareHours:                 NewFloat(0),
				CumulatedChildCareHours:        NewFloat(0),
				BillableHours:                  NewFloat(16),
				CumulatedBillableHours:         NewFloat(32),
				OfficeHours:                    NewFloat(24),
				CumulatedOfficeHours:           NewFloat(48),
				NonbillableHours:               NewFloat(8),
				CumulatedNonbillableHours:      NewFloat(16),
				BillingDegree:                  NewFloat(16).Div(NewFloat(24)),
				CumulatedBillingDegree:         NewFloat(32).Div(NewFloat(48)),
			},
		},
	}
	for _, test := range tests {
		res := NewBillingPeriodUserEntry(test.period, test.userId, test.trackedEntries)

		// TODO: reflect.DeepEqual won't work with big.Float
		if fmt.Sprintf("%+v", test.output) != fmt.Sprintf("%+v", res) {
			t.Logf("TrackedHours: %s\n", test.description)
			t.Logf("Expected billing period to equal\n%+v\n\tgot:\n%+v\n", test.output, res)
			t.Fail()
		}
	}
}
