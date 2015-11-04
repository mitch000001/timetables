package presenter

import "github.com/mitch000001/timetables/interaction"

const (
	DefaultDayPrecision           = 2
	DefaultWorkingDegreePrecision = 2
	DefaultBillingDegreePrecision = 2
	DefaultDateFormat             = "02.01.2006"
)

func NewBillingPeriodPresenter(billingPeriod interaction.BillingPeriod) BillingPeriodPresenter {
	return BillingPeriodPresenter{
		model:                  billingPeriod,
		DateFormat:             DefaultDateFormat,
		DayPrecision:           DefaultDayPrecision,
		BillingDegreePrecision: DefaultBillingDegreePrecision,
		WorkingDegreePrecision: DefaultWorkingDegreePrecision,
	}
}

type BillingPeriodPresenter struct {
	model                  interaction.BillingPeriod
	DateFormat             string
	DayPrecision           int
	BillingDegreePrecision int
	WorkingDegreePrecision int
}

func (b BillingPeriodPresenter) Present() BillingPeriod {
	var entries []BillingPeriodEntry
	for _, entry := range b.model.Entries {
		entryPresenter := NewBillingPeriodEntryPresenter(entry)
		entryPresenter.DayPrecision = b.DayPrecision
		entryPresenter.BillingDegreePrecision = b.BillingDegreePrecision
		entryPresenter.WorkingDegreePrecision = b.WorkingDegreePrecision
		entries = append(entries, entryPresenter.Present())
	}
	return BillingPeriod{
		StartDate: b.model.StartDate.Format(b.DateFormat),
		EndDate:   b.model.EndDate.Format(b.DateFormat),
		Entries:   entries,
	}
}

type BillingPeriod struct {
	StartDate string
	EndDate   string
	Entries   []BillingPeriodEntry
}

type BillingPeriodEntry struct {
	Name          string
	BillingDegree string
	WorkingDegree string
	FormattedBillingDelta
}

type BillingDelta struct {
	BillableDaysDelta             Delta
	CumulatedBillableDaysDelta    Delta
	NonbillableDaysDelta          Delta
	CumulatedNonbillableDaysDelta Delta
	VacationDaysDelta             Delta
	CumulatedVacationDaysDelta    Delta
	SicknessDaysDelta             Delta
	CumulatedSicknessDaysDelta    Delta
	ChildCareDaysDelta            Delta
	CumulatedChildCareDaysDelta   Delta
	OfficeDaysDelta               Delta
	CumulatedOfficeDaysDelta      Delta
	BillingDegreeDelta            Delta
	CumulatedBillingDegreeDelta   Delta
}

type FormattedBillingDelta struct {
	BillableDaysDelta             FormattedDelta
	CumulatedBillableDaysDelta    FormattedDelta
	NonbillableDaysDelta          FormattedDelta
	CumulatedNonbillableDaysDelta FormattedDelta
	VacationDaysDelta             FormattedDelta
	CumulatedVacationDaysDelta    FormattedDelta
	SicknessDaysDelta             FormattedDelta
	CumulatedSicknessDaysDelta    FormattedDelta
	ChildCareDaysDelta            FormattedDelta
	CumulatedChildCareDaysDelta   FormattedDelta
	OfficeDaysDelta               FormattedDelta
	CumulatedOfficeDaysDelta      FormattedDelta
	BillingDegreeDelta            FormattedDelta
	CumulatedBillingDegreeDelta   FormattedDelta
}

func NewBillingPeriodEntryPresenter(entry interaction.BillingPeriodEntry) BillingPeriodEntryPresenter {
	return BillingPeriodEntryPresenter{
		model:                  entry,
		DayPrecision:           DefaultDayPrecision,
		WorkingDegreePrecision: DefaultWorkingDegreePrecision,
		BillingDegreePrecision: DefaultBillingDegreePrecision,
	}
}

type BillingPeriodEntryPresenter struct {
	model                  interaction.BillingPeriodEntry
	DayPrecision           int
	WorkingDegreePrecision int
	BillingDegreePrecision int
}

func (b BillingPeriodEntryPresenter) Present() BillingPeriodEntry {
	var entry BillingPeriodEntry
	entry.Name = b.model.User.FirstName
	entry.BillingDegree = b.model.User.BillingDegree.FloatString(b.BillingDegreePrecision)
	entry.WorkingDegree = b.model.User.WorkingDegree.FloatString(b.WorkingDegreePrecision)
	deltaConverter := DeltaConverter{}
	billingDelta := deltaConverter.Convert(b.model.TrackedDays, b.model.EstimatedDays)
	entry.FormattedBillingDelta = BillingDeltaFormatter{}.Format(billingDelta, b.DayPrecision)
	return entry
}
