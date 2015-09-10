package timetables

func NewBillingPeriod(period Period) BillingPeriod {
	var billingPeriod = BillingPeriod{
		period:      period,
		userEntries: make([]BillingPeriodUserEntry, 0),
	}
	return billingPeriod
}

type BillingPeriod struct {
	ID          string
	period      Period
	userEntries []BillingPeriodUserEntry
}

func (b *BillingPeriod) AddUserEntry(userId string, trackedHours TrackedHours) {
	var entry = NewBillingPeriodUserEntry(b.period, userId, trackedHours)
	b.userEntries = append(b.userEntries, entry)
}
