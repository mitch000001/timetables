package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/interaction"
)

func TestNewBillingPeriodPresenter(t *testing.T) {
	var presenter BillingPeriodPresenter
	var billingPeriod = interaction.BillingPeriod{
		StartDate: timetables.Date(2015, 1, 1, time.Local),
		EndDate:   timetables.Date(2015, 1, 1, time.Local),
		Entries: []interaction.BillingPeriodEntry{
			interaction.BillingPeriodEntry{
				User: interaction.User{
					FirstName: "Max",
					LastName:  "Muster",
				},
				TrackedDays: interaction.Days{
					BillableDays:    timetables.NewRat(8),
					NonbillableDays: timetables.NewRat(8),
					VacationDays:    timetables.NewRat(8),
					SicknessDays:    timetables.NewRat(8),
					ChildCareDays:   timetables.NewRat(8),
					OfficeDays:      timetables.NewRat(8),
					BillingDegree:   timetables.NewRat(8),
				},
				EstimatedDays: interaction.Days{
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
		DateFormat: DefaultDateFormat,
		model:      billingPeriod,
	}

	if !reflect.DeepEqual(expectedPresenter, presenter) {
		t.Logf("Expected presenter to equal\n%+#v\n\tgot:\n%+#v\n", expectedPresenter, presenter)
		t.Fail()
	}
}

func TestBillingPeriodPresenterPresent(t *testing.T) {
	presenter := BillingPeriodPresenter{
		DateFormat: DefaultDateFormat,
		model: interaction.BillingPeriod{
			StartDate: timetables.Date(2015, 1, 1, time.Local),
			EndDate:   timetables.Date(2015, 1, 22, time.Local),
			Entries: []interaction.BillingPeriodEntry{
				interaction.BillingPeriodEntry{
					User: interaction.User{
						FirstName: "Max",
						LastName:  "Muster",
					},
					TrackedDays: interaction.Days{
						BillableDays:    timetables.NewRat(8),
						NonbillableDays: timetables.NewRat(8),
						VacationDays:    timetables.NewRat(8),
						SicknessDays:    timetables.NewRat(8),
						ChildCareDays:   timetables.NewRat(8),
						OfficeDays:      timetables.NewRat(8),
						BillingDegree:   timetables.NewRat(8),
					},
					EstimatedDays: interaction.Days{
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
		},
	}

	var viewModel BillingPeriod

	viewModel = presenter.Present()

	expectedViewModel := BillingPeriod{
		StartDate: "01.01.2015",
		EndDate:   "22.01.2015",
		Entries: []BillingPeriodEntry{
			BillingPeriodEntry{
				Name: "Max",
				FormattedBillingDelta: FormattedBillingDelta{
					BillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					NonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
					VacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					SicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					ChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
					OfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
					BillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
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
				Name: "Max",
				FormattedBillingDelta: FormattedBillingDelta{
					BillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					NonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
					VacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					SicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
					ChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
					OfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
					BillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
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
		model:        billingPeriodEntry,
		DayPrecision: 2,
	}

	if !reflect.DeepEqual(expectedPresenter, presenter) {
		t.Logf("Expected presenter to equal\n%+#v\n\tgot:\n%+#v\n", expectedPresenter, presenter)
		t.Fail()
	}
}

func TestBillingPeriodEntryPresenterPresent(t *testing.T) {
	var billingPeriodEntry = interaction.BillingPeriodEntry{
		User: interaction.User{
			FirstName: "Max",
			LastName:  "Muster",
		},
		TrackedDays: interaction.Days{
			BillableDays:    timetables.NewRat(8),
			NonbillableDays: timetables.NewRat(8),
			VacationDays:    timetables.NewRat(8),
			SicknessDays:    timetables.NewRat(8),
			ChildCareDays:   timetables.NewRat(8),
			OfficeDays:      timetables.NewRat(8),
			BillingDegree:   timetables.NewRat(8),
		},
		EstimatedDays: interaction.Days{
			BillableDays:    timetables.NewRat(7),
			NonbillableDays: timetables.NewRat(7),
			VacationDays:    timetables.NewRat(7),
			SicknessDays:    timetables.NewRat(7),
			ChildCareDays:   timetables.NewRat(7),
			OfficeDays:      timetables.NewRat(7),
			BillingDegree:   timetables.NewRat(7),
		},
	}
	presenter := BillingPeriodEntryPresenter{
		model:        billingPeriodEntry,
		DayPrecision: 2,
	}

	var viewModel BillingPeriodEntry

	viewModel = presenter.Present()

	expectedViewModel := BillingPeriodEntry{
		Name: "Max",
		FormattedBillingDelta: FormattedBillingDelta{
			BillableDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
			NonbillableDaysDelta: FormattedDelta{"8.00", "7.00", "1.00"},
			VacationDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
			SicknessDaysDelta:    FormattedDelta{"8.00", "7.00", "1.00"},
			ChildCareDaysDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
			OfficeDaysDelta:      FormattedDelta{"8.00", "7.00", "1.00"},
			BillingDegreeDelta:   FormattedDelta{"8.00", "7.00", "1.00"},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}

	// Changing DayPrecision

	presenter.DayPrecision = 5

	viewModel = presenter.Present()

	expectedViewModel = BillingPeriodEntry{
		Name: "Max",
		FormattedBillingDelta: FormattedBillingDelta{
			BillableDaysDelta:    FormattedDelta{"8.00000", "7.00000", "1.00000"},
			NonbillableDaysDelta: FormattedDelta{"8.00000", "7.00000", "1.00000"},
			VacationDaysDelta:    FormattedDelta{"8.00000", "7.00000", "1.00000"},
			SicknessDaysDelta:    FormattedDelta{"8.00000", "7.00000", "1.00000"},
			ChildCareDaysDelta:   FormattedDelta{"8.00000", "7.00000", "1.00000"},
			OfficeDaysDelta:      FormattedDelta{"8.00000", "7.00000", "1.00000"},
			BillingDegreeDelta:   FormattedDelta{"8.00000", "7.00000", "1.00000"},
		},
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}
}
