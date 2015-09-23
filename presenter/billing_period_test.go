package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/interaction"
)

func TestBillingPeriodPresenterPresent(t *testing.T) {
	presenter := BillingPeriodPresenter{
		model: interaction.BillingPeriod{
			StartDate: timetables.Date(2015, 1, 1, time.Local),
		},
	}

	var viewModel BillingPeriod

	viewModel = presenter.Present()

	expectedViewModel := BillingPeriod{
		Entries: []BillingPeriodEntry{},
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
		model: billingPeriodEntry,
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
		TrackedDays: interaction.TrackedDays{
			BillableDays: timetables.NewFloat(8),
		},
	}
	presenter := BillingPeriodEntryPresenter{
		model: billingPeriodEntry,
	}

	var viewModel BillingPeriodEntry

	viewModel = presenter.Present()

	expectedViewModel := BillingPeriodEntry{
		Name: "Max",
	}

	if !reflect.DeepEqual(expectedViewModel, viewModel) {
		t.Logf("Expected viewModel to equal\n%+#v\n\tgot:\n%+#v\n", expectedViewModel, viewModel)
		t.Fail()
	}
}
