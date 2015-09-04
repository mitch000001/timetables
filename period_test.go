package timetables

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPeriod(t *testing.T) {
	p := NewPeriod(NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), 25)

	expectedTimeframe := Timeframe{
		StartDate: Date(2015, 1, 1, time.Local),
		EndDate:   Date(2015, 2, 1, time.Local),
	}

	if !reflect.DeepEqual(expectedTimeframe, p.Timeframe) {
		t.Logf("Expected period timeframe to equal\n%s\n\tgot\n%s\n", expectedTimeframe, p.Timeframe)
		t.Fail()
	}

	if p.BusinessDays != 25.0 {
		t.Logf("Expected period BusinessDays to equal 25, got %d\n", p.BusinessDays)
		t.Fail()
	}

	// invalid businessDays

	p = NewPeriod(NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), 0)

	expectedBusinessDays := 32

	if float64(expectedBusinessDays) != p.BusinessDays {
		t.Logf("Expected period BusinessDays to equal %f, got %f\n", float64(expectedBusinessDays), p.BusinessDays)
		t.Fail()
	}

	p = NewPeriod(NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), -10)

	expectedBusinessDays = 32

	if float64(expectedBusinessDays) != p.BusinessDays {
		t.Logf("Expected period BusinessDays to equal %f, got %f\n", float64(expectedBusinessDays), p.BusinessDays)
		t.Fail()
	}
}
