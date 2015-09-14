package timetables

import (
	"bytes"
	"fmt"
)

func NewBillingPeriod(period Period) BillingPeriod {
	var billingPeriod = BillingPeriod{
		Period:      period,
		userEntries: make([]BillingPeriodUserEntry, 0),
	}
	return billingPeriod
}

type BillingPeriod struct {
	ID          string
	Period      Period
	userEntries []BillingPeriodUserEntry
}

func (b *BillingPeriod) AddUserEntry(userId string, trackedHours TrackedHours) {
	var entry = NewBillingPeriodUserEntry(b.Period, userId, trackedHours)
	b.userEntries = append(b.userEntries, entry)
}

func (b BillingPeriod) MarshalText() ([]byte, error) {
	marshaledPeriod, err := b.Period.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("Error while marshaling period: %v", err)
	}
	marshaled := fmt.Sprintf("{%s}:{%s}:[]", b.ID, marshaledPeriod)
	return []byte(marshaled), nil
}

func (b *BillingPeriod) UnmarshalText(value []byte) error {
	parts := bytes.SplitN(value, []byte(":"), 2)
	id := bytes.Trim(parts[0], "{}")
	b.ID = string(id)
	idx := bytes.Index(parts[1], []byte(":["))
	periodBytes := bytes.Trim(parts[1][1:idx], "{}")
	err := b.Period.UnmarshalText(periodBytes)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling period: %v", err)
	}
	return nil
}
