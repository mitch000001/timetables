package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/date"
	"github.com/mitch000001/timetables/interaction"
)

func TestNewBillingPeriodPresenter(t *testing.T) {
	var presenter BillingPeriodPresenter
	var billingPeriod = interaction.BillingPeriod{
		StartDate: date.Date(2015, 1, 1, time.Local),
		EndDate:   date.Date(2015, 1, 1, time.Local),
		Entries: []interaction.BillingPeriodEntry{
			interaction.BillingPeriodEntry{
				User: interaction.User{
					FirstName: "Max",
					LastName:  "Muster",
				},
				TrackedDays: interaction.PeriodData{
					BillableDays:    timetables.NewRat(8),
					NonbillableDays: timetables.NewRat(8),
					VacationDays:    timetables.NewRat(8),
					SicknessDays:    timetables.NewRat(8),
					ChildCareDays:   timetables.NewRat(8),
					OfficeDays:      timetables.NewRat(8),
					BillingDegree:   timetables.NewRat(8),
				},
				ForecastedDays: interaction.PeriodData{
					BillableDays:    timetables.NewRat(7),
					NonbillableDays: timetables.NewRat(7),
					VacationDays:    timetables.NewRat(7),
					SicknessDays:    timetables.NewRat(7),
					ChildCareDays:   timetables.NewRat(7),
					OfficeDays:      timetables.NewRat(7),
					BillingDegree:   timetables.NewRat(7),
				},
			},
		},
	}

	presenter = NewBillingPeriodPresenter(billingPeriod)

	expectedPresenter := BillingPeriodPresenter{
		DateFormat:             DefaultDateFormat,
		DayPrecision:           DefaultDayPrecision,
		BillingDegreePrecision: DefaultBillingDegreePrecision,
		WorkingDegreePrecision: DefaultWorkingDegreePrecision,
		model: billingPeriod,
	}

	if !reflect.DeepEqual(expectedPresenter, presenter) {
		t.Logf("Expected presenter to equal\n%+#v\n\tgot:\n%+#v\n", expectedPresenter, presenter)
		t.Fail()
	}
}

func TestBillingPeriodPresenterPresent(t *testing.T) {
	presenter := NewBillingPeriodPresenter(interaction.BillingPeriod{
		StartDate: date.Date(2015, 1, 1, time.Local),
		EndDate:   date.Date(2015, 1, 22, time.Local),
		Entries: []interaction.BillingPeriodEntry{
			interaction.BillingPeriodEntry{
				User: interaction.User{
					FirstName:     "Max",
					LastName:      "Muster",
					BillingDegree: timetables.NewRat(0.8),
					WorkingDegree: timetables.NewRat(1),
				},
				TrackedDays: interaction.PeriodData{
					BillableDays:             timetables.NewRat(8),
					CumulatedBillableDays:    timetables.NewRat(8),
					NonbillableDays:          timetables.NewRat(8),
					CumulatedNonbillableDays: timetables.NewRat(8),
					VacationDays:             timetables.NewRat(8),
					CumulatedVacationDays:    timetables.NewRat(8),
					SicknessDays:             timetables.NewRat(8),
					CumulatedSicknessDays:    timetables.NewRat(8),
					ChildCareDays:            timetables.NewRat(8),
					CumulatedChildCareDays:   timetables.NewRat(8),
					OfficeDays:               timetables.NewRat(8),
					CumulatedOfficeDays:      timetables.NewRat(8),
					BillingDegree:            timetables.NewRat(8),
					CumulatedBillingDegree:   timetables.NewRat(8),
				},
				ForecastedDays: interaction.PeriodData{
					BillableDays:             timetables.NewRat(7),
					CumulatedBillableDays:    timetables.NewRat(7),
					NonbillableDays:          timetables.NewRat(7),
					CumulatedNonbillableDays: timetables.NewRat(7),
					VacationDays:             timetables.NewRat(7),
					CumulatedVacationDays:    timetables.NewRat(7),
					SicknessDays:             timetables.NewRat(7),
					CumulatedSicknessDays:    timetables.NewRat(7),
					ChildCareDays:            timetables.NewRat(7),
					CumulatedChildCareDays:   timetables.NewRat(7),
					OfficeDays:               timetables.NewRat(7),
					CumulatedOfficeDays:      timetables.NewRat(7),
					BillingDegree:            timetables.NewRat(7),
					CumulatedBillingDegree:   timetables.NewRat(7),
				},
			},
		},
	},
	)

	var viewModel BillingPeriod

	viewModel = presenter.Present()

	expectedViewModel := BillingPeriod{
		StartDate: "01.01.2015",
		EndDate:   "22.01.2015",
		Entries: []BillingPeriodEntry{
			BillingPeriodEntry{
				Name:          "Max",
				BillingDegree: "0.80",
				WorkingDegree: "1.00",
				FormattedBillingDelta: FormattedBillingDelta{
					BillableDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedBillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					NonbillableDaysDelta:          FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedNonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
					VacationDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedVacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					SicknessDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedSicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					ChildCareDaysDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
					OfficeDaysDelta:               FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedOfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
					BillingDegreeDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedBillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}

	// Changed DateFormat

	presenter.DateFormat = "02/01/2006"

	viewModel = presenter.Present()

	expectedViewModel = BillingPeriod{
		StartDate: "01/01/2015",
		EndDate:   "22/01/2015",
		Entries: []BillingPeriodEntry{
			BillingPeriodEntry{
				Name:          "Max",
				BillingDegree: "0.80",
				WorkingDegree: "1.00",
				FormattedBillingDelta: FormattedBillingDelta{
					BillableDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedBillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					NonbillableDaysDelta:          FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedNonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
					VacationDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedVacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					SicknessDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedSicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					ChildCareDaysDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
					OfficeDaysDelta:               FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedOfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
					BillingDegreeDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
					CumulatedBillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}

	// Changed DayPrecision, BillingDegreePrecision, WorkingDegreePrecision

	presenter.DayPrecision = 4
	presenter.BillingDegreePrecision = 3
	presenter.WorkingDegreePrecision = 5

	viewModel = presenter.Present()

	expectedViewModel = BillingPeriod{
		StartDate: "01/01/2015",
		EndDate:   "22/01/2015",
		Entries: []BillingPeriodEntry{
			BillingPeriodEntry{
				Name:          "Max",
				BillingDegree: "0.800",
				WorkingDegree: "1.00000",
				FormattedBillingDelta: FormattedBillingDelta{
					BillableDaysDelta:             FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedBillableDaysDelta:    FormattedDelta{"8.0000", "7.0000", "1.0000"},
					NonbillableDaysDelta:          FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedNonbillableDaysDelta: FormattedDelta{"8.0000", "7.0000", "1.0000"},
					VacationDaysDelta:             FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedVacationDaysDelta:    FormattedDelta{"8.0000", "7.0000", "1.0000"},
					SicknessDaysDelta:             FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedSicknessDaysDelta:    FormattedDelta{"8.0000", "7.0000", "1.0000"},
					ChildCareDaysDelta:            FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedChildCareDaysDelta:   FormattedDelta{"8.0000", "7.0000", "1.0000"},
					OfficeDaysDelta:               FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedOfficeDaysDelta:      FormattedDelta{"8.0000", "7.0000", "1.0000"},
					BillingDegreeDelta:            FormattedDelta{"8.0000", "7.0000", "1.0000"},
					CumulatedBillingDegreeDelta:   FormattedDelta{"8.0000", "7.0000", "1.0000"},
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}
}

func TestNewBillingPeriodEntryPresenter(t *testing.T) {
	var presenter BillingPeriodEntryPresenter
	var billingPeriodEntry = interaction.BillingPeriodEntry{
		User: interaction.User{
			FirstName: "Max",
		},
	}

	presenter = NewBillingPeriodEntryPresenter(billingPeriodEntry)

	expectedPresenter := BillingPeriodEntryPresenter{
		model:                  billingPeriodEntry,
		DayPrecision:           2,
		WorkingDegreePrecision: 2,
		BillingDegreePrecision: 2,
	}

	if !reflect.DeepEqual(expectedPresenter, presenter) {
		t.Logf("Expected presenter to equal\n%+#v\n\tgot:\n%+#v\n", expectedPresenter, presenter)
		t.Fail()
	}
}

func TestBillingPeriodEntryPresenterPresent(t *testing.T) {
	var billingPeriodEntry = interaction.BillingPeriodEntry{
		User: interaction.User{
			FirstName:     "Max",
			LastName:      "Muster",
			WorkingDegree: timetables.NewRat(1),
			BillingDegree: timetables.NewRat(0.8),
		},
		TrackedDays: interaction.PeriodData{
			BillableDays:             timetables.NewRat(8),
			CumulatedBillableDays:    timetables.NewRat(8),
			NonbillableDays:          timetables.NewRat(8),
			CumulatedNonbillableDays: timetables.NewRat(8),
			VacationDays:             timetables.NewRat(8),
			CumulatedVacationDays:    timetables.NewRat(8),
			SicknessDays:             timetables.NewRat(8),
			CumulatedSicknessDays:    timetables.NewRat(8),
			ChildCareDays:            timetables.NewRat(8),
			CumulatedChildCareDays:   timetables.NewRat(8),
			OfficeDays:               timetables.NewRat(8),
			CumulatedOfficeDays:      timetables.NewRat(8),
			BillingDegree:            timetables.NewRat(8),
			CumulatedBillingDegree:   timetables.NewRat(8),
		},
		ForecastedDays: interaction.PeriodData{
			BillableDays:             timetables.NewRat(7),
			CumulatedBillableDays:    timetables.NewRat(7),
			NonbillableDays:          timetables.NewRat(7),
			CumulatedNonbillableDays: timetables.NewRat(7),
			VacationDays:             timetables.NewRat(7),
			CumulatedVacationDays:    timetables.NewRat(7),
			SicknessDays:             timetables.NewRat(7),
			CumulatedSicknessDays:    timetables.NewRat(7),
			ChildCareDays:            timetables.NewRat(7),
			CumulatedChildCareDays:   timetables.NewRat(7),
			OfficeDays:               timetables.NewRat(7),
			CumulatedOfficeDays:      timetables.NewRat(7),
			BillingDegree:            timetables.NewRat(7),
			CumulatedBillingDegree:   timetables.NewRat(7),
		},
	}
	presenter := BillingPeriodEntryPresenter{
		model:                  billingPeriodEntry,
		DayPrecision:           2,
		WorkingDegreePrecision: 2,
		BillingDegreePrecision: 2,
	}

	var viewModel BillingPeriodEntry

	viewModel = presenter.Present()

	expectedViewModel := BillingPeriodEntry{
		Name:          "Max",
		BillingDegree: "0.80",
		WorkingDegree: "1.00",
		FormattedBillingDelta: FormattedBillingDelta{
			BillableDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedBillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
			NonbillableDaysDelta:          FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedNonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
			VacationDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedVacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
			SicknessDaysDelta:             FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedSicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
			ChildCareDaysDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
			OfficeDaysDelta:               FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedOfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
			BillingDegreeDelta:            FormattedDelta{"8.00", "7.00", "1.00"},
			CumulatedBillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}

	// Changing DayPrecision, WorkingDegreePrecision, BillingDegreePrecision

	presenter.DayPrecision = 5
	presenter.WorkingDegreePrecision = 4
	presenter.BillingDegreePrecision = 3

	viewModel = presenter.Present()

	expectedViewModel = BillingPeriodEntry{
		Name:          "Max",
		BillingDegree: "0.800",
		WorkingDegree: "1.0000",
		FormattedBillingDelta: FormattedBillingDelta{
			BillableDaysDelta:             FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedBillableDaysDelta:    FormattedDelta{"8.00000", "7.00000", "1.00000"},
			NonbillableDaysDelta:          FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedNonbillableDaysDelta: FormattedDelta{"8.00000", "7.00000", "1.00000"},
			VacationDaysDelta:             FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedVacationDaysDelta:    FormattedDelta{"8.00000", "7.00000", "1.00000"},
			SicknessDaysDelta:             FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedSicknessDaysDelta:    FormattedDelta{"8.00000", "7.00000", "1.00000"},
			ChildCareDaysDelta:            FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedChildCareDaysDelta:   FormattedDelta{"8.00000", "7.00000", "1.00000"},
			OfficeDaysDelta:               FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedOfficeDaysDelta:      FormattedDelta{"8.00000", "7.00000", "1.00000"},
			BillingDegreeDelta:            FormattedDelta{"8.00000", "7.00000", "1.00000"},
			CumulatedBillingDegreeDelta:   FormattedDelta{"8.00000", "7.00000", "1.00000"},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}
}
