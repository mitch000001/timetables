package presenter

import "github.com/mitch000001/timetables"

type Delta struct {
	Tracked   *timetables.Rat
	Estimated *timetables.Rat
}

func (d Delta) Delta() *timetables.Rat {
	return d.Tracked.Sub(d.Estimated)
}

type FormattedDelta struct {
	Tracked   string
	Estimated string
	Delta     string
}
