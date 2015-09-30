package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables/date"
)

func TestNewPeriod(t *testing.T) {
	p := NewPeriod(date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), 25)

	expectedTimeframe := date.Timeframe{
		StartDate: date.Date(2015, 1, 1, time.Local),
		EndDate:   date.Date(2015, 2, 1, time.Local),
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

	p = NewPeriod(date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), 0)

	expectedBusinessDays := 32

	if float64(expectedBusinessDays) != p.BusinessDays {
		t.Logf("Expected period BusinessDays to equal %f, got %f\n", float64(expectedBusinessDays), p.BusinessDays)
		t.Fail()
	}

	p = NewPeriod(date.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), -10)

	expectedBusinessDays = 32

	if float64(expectedBusinessDays) != p.BusinessDays {
		t.Logf("Expected period BusinessDays to equal %f, got %f\n", float64(expectedBusinessDays), p.BusinessDays)
		t.Fail()
	}
}

func TestPeriodMarshalText(t *testing.T) {
	period := Period{
		Timeframe:    date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}

	expected := "{{2015-01-01}:{2015-01-25}}:{20}"

	var marshaled []byte
	var err error

	marshaled, err = period.MarshalText()

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	if expected != string(marshaled) {
		t.Logf("Expected marshaled value to equal\n%q\n\tgot:\n%q\n", expected, marshaled)
		t.Fail()
	}
}

func TestPeriodUnmarshalText(t *testing.T) {
	expectedPeriod := Period{
		Timeframe:    date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local),
		BusinessDays: 20,
	}
	marshaled := "{2015-01-01:2015-01-25}:{20}"

	var err error

	period := Period{}

	err = period.UnmarshalText([]byte(marshaled))

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedPeriod, period) {
		t.Logf("Expected period to equal\n%+v\n\tgot\n%+v\n", expectedPeriod, period)
		t.Fail()
	}
}
