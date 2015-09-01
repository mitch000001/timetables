package timetables

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/go-harvest/harvest"
)

func TestNewPeriod(t *testing.T) {
	p := NewPeriod(harvest.NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local), 25)

	expectedTimeframe := harvest.Timeframe{
		StartDate: harvest.Date(2015, 1, 1, time.Local),
		EndDate:   harvest.Date(2015, 2, 1, time.Local),
	}

	if !reflect.DeepEqual(expectedTimeframe, p.Timeframe) {
		t.Logf("Expected period timeframe to equal\n%s\n\tgot\n%s\n", expectedTimeframe, p.Timeframe)
		t.Fail()
	}
}
