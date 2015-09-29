package presenter

import (
	"testing"

	"github.com/mitch000001/timetables"
)

func TestDeltaDelta(t *testing.T) {
	tests := []struct {
		tracked   *timetables.Rat
		estimated *timetables.Rat
		delta     *timetables.Rat
	}{
		{
			tracked:   timetables.NewRat(8),
			estimated: timetables.NewRat(5),
			delta:     timetables.NewRat(3),
		},
		{
			tracked:   timetables.NewRat(5),
			estimated: timetables.NewRat(8),
			delta:     timetables.NewRat(-3),
		},
	}

	for _, test := range tests {
		d := Delta{test.tracked, test.estimated}

		delta := d.Delta()

		if test.delta.Cmp(delta) != 0 {
			t.Logf("Expected Delta to return\n%+v\n\tgot:\n%+v\n", test.delta, delta)
			t.Fail()
		}
	}
}
