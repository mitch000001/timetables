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
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "1",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, IsBilled: true, IsClosed: true},
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, IsBilled: true, IsClosed: true},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 1, IsBilled: false, IsClosed: true},
				},
			},
			output: BillingPeriod{
				UserID:                        "1",
				Timeframe:                     harvest.NewTimeframe(2015, 1, 1, 2015, 25, 1, time.Local),
				BusinessDays:                  NewFloat(10),
				CumulatedBusinessDays:         NewFloat(10),
				VacationInterest:              NewFloat(0),
				CumulatedVacationInterest:     NewFloat(0),
				SicknessInterest:              NewFloat(0),
				CumulatedSicknessInterest:     NewFloat(0),
				BilledDays:                    NewFloat(2),
				CumulatedBilledDays:           NewFloat(2),
				OfficeDays:                    NewFloat(3),
				CumulatedOfficeDays:           NewFloat(3),
				OverheadAndSlacktime:          NewFloat(1),
				CumulatedOverheadAndSlacktime: NewFloat(1),
				BillingDegree:                 NewFloat(2).Div(NewFloat(3)),
				CumulatedBillingDegree:        NewFloat(2).Div(NewFloat(3)),
			},
		},
		{
			period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)}, 10},
			trackedEntries: TrackedEntries{
				billingConfig: BillingConfig{
					UserID:         "2",
					VacationTaskID: 10,
					SicknessTaskID: 13,
				},
				billableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, IsBilled: true, IsClosed: true},
					&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 2, IsBilled: true, IsClosed: true},
				},
				nonbillableEntries: []*harvest.DayEntry{
					&harvest.DayEntry{Hours: 8, TaskId: 12, UserId: 2, IsBilled: false, IsClosed: true},
				},
			},
			output: BillingPeriod{
				UserID:                        "2",
				Timeframe:                     harvest.NewTimeframe(2015, 1, 1, 2015, 25, 1, time.Local),
				BusinessDays:                  NewFloat(10),
				CumulatedBusinessDays:         NewFloat(10),
				VacationInterest:              NewFloat(0),
				CumulatedVacationInterest:     NewFloat(0),
				SicknessInterest:              NewFloat(0),
				CumulatedSicknessInterest:     NewFloat(0),
				BilledDays:                    NewFloat(2),
				CumulatedBilledDays:           NewFloat(2),
				OfficeDays:                    NewFloat(3),
				CumulatedOfficeDays:           NewFloat(3),
				OverheadAndSlacktime:          NewFloat(1),
				CumulatedOverheadAndSlacktime: NewFloat(1),
				BillingDegree:                 NewFloat(2).Div(NewFloat(3)),
				CumulatedBillingDegree:        NewFloat(2).Div(NewFloat(3)),
			},
		},
		//{
		//period: Period{harvest.Timeframe{StartDate: harvest.Date(2015, 1, 1, time.Local), EndDate: harvest.Date(2015, 25, 1, time.Local)}, 10},
		//billingConfig: BillingConfig{
		//UserID:         "1",
		//VacationTaskID: 15,
		//},
		//entries: []*harvest.DayEntry{
		//&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, IsBilled: true, IsClosed: true},
		//&harvest.DayEntry{Hours: 8, TaskId: 5, UserId: 1, IsBilled: true, IsClosed: true},
		//&harvest.DayEntry{Hours: 8, TaskId: 15, UserId: 1, IsBilled: false, IsClosed: true},
		//},
		//output: BillingPeriod{
		//UserID:                        "1",
		//Timeframe:                     harvest.NewTimeframe(2015, 1, 1, 2015, 25, 1, time.Local),
		//BusinessDays:                  NewFloat(10),
		//CumulatedBusinessDays:         NewFloat(10),
		//VacationInterest:              NewFloat(1),
		//CumulatedVacationInterest:     NewFloat(0),
		//SicknessInterest:              NewFloat(0),
		//CumulatedSicknessInterest:     NewFloat(0),
		//BilledDays:                    NewFloat(2),
		//CumulatedBilledDays:           NewFloat(2),
		//OfficeDays:                    NewFloat(2),
		//CumulatedOfficeDays:           NewFloat(2),
		//OverheadAndSlacktime:          NewFloat(0),
		//CumulatedOverheadAndSlacktime: NewFloat(0),
		//BillingDegree:                 NewFloat(1).Div(NewFloat(2)),
		//CumulatedBillingDegree:        NewFloat(1).Div(NewFloat(2)),
		//},
		//},
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
