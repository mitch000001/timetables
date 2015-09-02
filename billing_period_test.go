package timetables

import (
	"fmt"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

func TestCreateBillingPeriod(t *testing.T) {
	tests := []struct {
		period         Period
		trackedEntries TrackedEntries
		output         BillingPeriod
	}{
		{
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 1, 25, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "1",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 1, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriod{
				UserID:                    "1",
				Timeframe:                 harvest.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
				BusinessDays:              NewFloat(10),
				VacationInterestHours:     NewFloat(0),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(16),
				OfficeHours:               NewFloat(24),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 1, 25, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "2",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 6, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
					&harvest.DayEntry{Hours: 6, TaskId: 5, UserId: 2, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 6, TaskId: 12, UserId: 2, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
				},
			},
			output: BillingPeriod{
				UserID:                    "2",
				Timeframe:                 harvest.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
				BusinessDays:              NewFloat(10),
				VacationInterestHours:     NewFloat(0),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(12),
				OfficeHours:               NewFloat(18),
				OverheadAndSlacktimeHours: NewFloat(6),
				BillingDegree:             NewFloat(12).Div(NewFloat(18)),
			},
		},
		{
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 1, 25, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "1",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 1, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 10, UserId: 1, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
				},
			},
			output: BillingPeriod{
				UserID:                    "1",
				Timeframe:                 harvest.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
				BusinessDays:              NewFloat(10),
				VacationInterestHours:     NewFloat(8),
				SicknessInterestHours:     NewFloat(0),
				BilledHours:               NewFloat(16),
				OfficeHours:               NewFloat(24),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 1, 25, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "1",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 1, 1, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 1, 2, time.Local)},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 1, SpentAt: harvest.Date(2015, 1, 3, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 10, UserId: 1, SpentAt: harvest.Date(2015, 1, 4, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 13, UserId: 1, SpentAt: harvest.Date(2015, 1, 5, time.Local)},
				},
			},
			output: BillingPeriod{
				UserID:                    "1",
				Timeframe:                 harvest.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
				BusinessDays:              NewFloat(10),
				VacationInterestHours:     NewFloat(8),
				SicknessInterestHours:     NewFloat(8),
				BilledHours:               NewFloat(16),
				OfficeHours:               NewFloat(24),
				OverheadAndSlacktimeHours: NewFloat(8),
				BillingDegree:             NewFloat(16).Div(NewFloat(24)),
			},
		},
		{
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 1, 25, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "1",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 2, 1, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, SpentAt: harvest.Date(2015, 2, 2, time.Local)},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 1, SpentAt: harvest.Date(2015, 2, 3, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 10, UserId: 1, SpentAt: harvest.Date(2015, 2, 4, time.Local)},
					&harvest.DayEntry{Hours: 8, TaskId: 13, UserId: 1, SpentAt: harvest.Date(2015, 2, 5, time.Local)},
				},
			},
			output: BillingPeriod{
				UserID:                    "1",
				Timeframe:                 harvest.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
				BusinessDays:              NewFloat(10),
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
		res, _ := CreateBillingPeriod(test.period, test.trackedEntries)

		// TODO: reflect.DeepEqual won't work with big.Float
		if fmt.Sprintf("%#v", test.output) != fmt.Sprintf("%#v", res) {
			t.Logf("Expected estimation period to equal\n%#v\n\tgot:\n%#v\n", test.output, res)
			t.Fail()
		}
	}
}
