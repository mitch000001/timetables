package timetables

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateBillingPeriodUserEntry(t *testing.T) {
	tests := []struct {
		period         Period
		userId         string
		trackedEntries TrackedHours
		output         BillingPeriodUserEntry
	}{
		{
			period: Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId: "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                    "1",
				Period:                    Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:     NewFloat(0),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(16),
				OfficeHours:               NewFloat(24),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			period: Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId: "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "2", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                    "1",
				Period:                    Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:     NewFloat(0),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(8),
				OfficeHours:               NewFloat(16),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(8).Div(NewFloat(16)),
			},
		},
		{
			period: Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId: "2",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(6), Type: Billable, UserID: "2", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(6), Type: Billable, UserID: "2", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(6), Type: NonBillable, UserID: "2", TrackedAt: Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                    "2",
				Period:                    Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:     NewFloat(0),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(12),
				OfficeHours:               NewFloat(18),
				OverheadAndSlacktimeHours: NewFloat(6),
				BillingDegree:             NewFloat(12).Div(NewFloat(18)),
			},
		},
		{
			period: Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId: "1",
			trackedEntries: TrackedHours{
				entries: []TrackingEntry{
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 1, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Billable, UserID: "1", TrackedAt: Date(2015, 1, 2, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: NonBillable, UserID: "1", TrackedAt: Date(2015, 1, 3, time.Local)},
					TrackingEntry{Hours: NewFloat(8), Type: Vacation, UserID: "1", TrackedAt: Date(2015, 1, 4, time.Local)},
				},
			},
			output: BillingPeriodUserEntry{
				UserID:                    "1",
				Period:                    Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:     NewFloat(8),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(16),
				OfficeHours:               NewFloat(24),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			period: Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId: "1",
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
				UserID:                    "1",
				Period:                    Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:     NewFloat(8),
				SicknessInterestHours:     NewFloat(8),
				BilledHours:               NewFloat(16),
				OfficeHours:               NewFloat(24),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			period: Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
			userId: "1",
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
				UserID:                    "1",
				Period:                    Period{Timeframe{StartDate: Date(2015, 1, 1, time.Local), EndDate: Date(2015, 1, 25, time.Local)}, 10},
				VacationInterestHours:     NewFloat(0),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(0),
				OfficeHours:               NewFloat(0),
				OverheadAndSlacktimeHours: NewFloat(0),
				BillingDegree:             NewFloat(0),
			},
		},
	}
	for _, test := range tests {
		res, _ := CreateBillingPeriodUserEntry(test.period, test.userId, test.trackedEntries)

		// TODO: reflect.DeepEqual won't work with big.Float
		if fmt.Sprintf("%#v", test.output) != fmt.Sprintf("%#v", res) {
			t.Logf("Expected billing period to equal\n%#v\n\tgot:\n%#v\n", test.output, res)
			t.Fail()
		}
	}
}
