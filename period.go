package timetables

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mitch000001/timetables/date"
)

type PeriodProvider interface {
	Period() Period
}

func NewPeriod(timeframe date.Timeframe, businessDays float64) Period {
	if businessDays <= 0 {
		businessDays = float64(timeframe.Days())
	}
	return Period{
		Timeframe:    timeframe,
		BusinessDays: businessDays,
	}
}

type Period struct {
	Timeframe    date.Timeframe
	BusinessDays float64
}

func (p Period) MarshalText() ([]byte, error) {
	marshaledTimeframe, err := p.Timeframe.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("Error while marshaling Timeframe: %v", err)
	}
	marshaled := fmt.Sprintf("{%s}:{%g}", marshaledTimeframe, p.BusinessDays)
	return []byte(marshaled), nil
}

func (p *Period) UnmarshalText(value []byte) error {
	parts := bytes.Split(value, []byte(":"))
	if len(parts) != 3 {
		return fmt.Errorf("Malformed marshaled value")
	}
	timeframe := bytes.Join(parts[0:2], []byte(":"))
	timeframe = bytes.Trim(timeframe, "{}")
	err := p.Timeframe.UnmarshalText(timeframe)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling Timeframe: %v", err)
	}
	businessDays := bytes.Trim(parts[2], "{}")
	p.BusinessDays, err = strconv.ParseFloat(string(businessDays), 64)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling BusinessDays: %v", err)
	}
	return nil
}
