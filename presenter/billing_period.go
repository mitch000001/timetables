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
	StartDate string               `json:"start_date"`
	EndDate   string               `json:"end_date"`
	Entries   []BillingPeriodEntry `json:"entries"`
}

type BillingPeriodEntry struct {
	Name          string `json:"name"`
	BillingDegree string `json:"billing_degree"`
	WorkingDegree string `json:"working_degree"`
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
	BillableDaysDelta             FormattedDelta `json:"billable_days_delta"`
	CumulatedBillableDaysDelta    FormattedDelta `json:"cumulated_billable_days_delta"`
	NonbillableDaysDelta          FormattedDelta `json:"nonbillable_days_delta"`
	CumulatedNonbillableDaysDelta FormattedDelta `json:"cumulated_nonbillable_days_delta"`
	VacationDaysDelta             FormattedDelta `json:"vacation_days_delta"`
	CumulatedVacationDaysDelta    FormattedDelta `json:"cumulated_vacation_days_delta"`
	SicknessDaysDelta             FormattedDelta `json:"sickness_days_delta"`
	CumulatedSicknessDaysDelta    FormattedDelta `json:"cumulated_sickness_days_delta"`
	ChildCareDaysDelta            FormattedDelta `json:"child_care_days_delta"`
	CumulatedChildCareDaysDelta   FormattedDelta `json:"cumulated_child_care_days_delta"`
	OfficeDaysDelta               FormattedDelta `json:"office_days_delta"`
	CumulatedOfficeDaysDelta      FormattedDelta `json:"cumulated_office_days_delta"`
	BillingDegreeDelta            FormattedDelta `json:"billing_degree_delta"`
	CumulatedBillingDegreeDelta   FormattedDelta `json:"cumulated_billing_degree_delta"`
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
	billingDelta := deltaConverter.Convert(b.model.TrackedDays, b.model.ForecastedDays)
	entry.FormattedBillingDelta = BillingDeltaFormatter{}.Format(billingDelta, b.DayPrecision)
	return entry
}
