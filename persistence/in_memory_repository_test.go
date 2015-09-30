package persistence

import (
	"reflect"
	"testing"
	"time"

	"github.com/mitch000001/timetables"
	"github.com/mitch000001/timetables/date"
)

func TestNewInMemoryBillingPeriodRepository(t *testing.T) {
	inMemory := NewInMemoryBillingPeriodRepository()

	expectedRepository := InMemoryBillingPeriodRepository{
		store: make(map[ID]timetables.BillingPeriod),
	}

	if !reflect.DeepEqual(expectedRepository, inMemory) {
		t.Logf("Expected repository to equal\n%q\n\tgot\n%q\n", expectedRepository, inMemory)
		t.Fail()
	}
}

func TestInMemoryBillingPeriodRepositorySave(t *testing.T) {
	store := InMemoryBillingPeriodRepository{
		store: make(map[ID]timetables.BillingPeriod),
	}

	billingPeriod := timetables.NewBillingPeriod(timetables.Period{"1", date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 16})

	var err error
	var key ID

	key, err = store.Save(billingPeriod)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if key == nil {
		t.Logf("Expected key not to be nil\n")
		t.Fail()
	}

	if len(store.store) != 1 {
		t.Logf("Expected store to have one item, got %d\n", len(store.store))
		t.Fail()
	}

	if stored := store.store[key]; !reflect.DeepEqual(billingPeriod, stored) {
		t.Logf("Expected stored period to equal\n%q\n\tgot\n%q\n", billingPeriod, stored)
		t.Fail()
	}

	// Test for key difference on store
	var newKey ID
	newKey, err = store.Save(billingPeriod)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if reflect.DeepEqual(key, newKey) {
		t.Logf("Expected new key to be different than the old key\n")
		t.Fail()
	}
}

func TestInMemoryBillingPeriodRepositoryLoad(t *testing.T) {
	store := InMemoryBillingPeriodRepository{
		store: make(map[ID]timetables.BillingPeriod),
	}

	expectedEntry := timetables.NewBillingPeriod(timetables.Period{"1", date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 16})

	key, err := store.Save(expectedEntry)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	entry, err := store.Load(key)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedEntry, entry) {
		t.Logf("Expected loaded period to equal\n%q\n\tgot\n%q\n", expectedEntry, entry)
		t.Fail()
	}

	// Not found
	entry, err = store.Load(idString("foo"))

	expectedEntry = timetables.BillingPeriod{}

	if err != nil {
		if !IsNotFound(err) {
			t.Logf("Expected error to be a not found error, got %T:%v\n", err, err)
			t.Fail()
		}
	} else {
		t.Logf("Expected error not to be nil")
		t.Fail()
	}

	if !reflect.DeepEqual(expectedEntry, entry) {
		t.Logf("Expected loaded period to equal\n%q\n\tgot\n%q\n", expectedEntry, entry)
		t.Fail()
	}
}

func TestInMemoryBillingPeriodRepositoryUpdate(t *testing.T) {
	store := InMemoryBillingPeriodRepository{
		store: make(map[ID]timetables.BillingPeriod),
	}

	entry := timetables.NewBillingPeriod(timetables.Period{"1", date.NewTimeframe(2015, 1, 1, 2015, 1, 25, time.Local), 16})

	key, err := store.Save(entry)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	entry.AddUserEntry("1", timetables.NewTrackedHours([]timetables.TrackingEntry{
		timetables.TrackingEntry{Hours: timetables.NewRat(8), UserID: "1", TrackedAt: date.Date(2015, 1, 15, time.Local), Type: timetables.Billable},
	}))

	err = store.Update(key, entry)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	updatedEntry, err := store.Load(key)

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(entry, updatedEntry) {
		t.Logf("Expected updated period to equal\n%q\n\tgot\n%q\n", entry, updatedEntry)
		t.Fail()
	}

	// Not found
	err = store.Update(idString("foo"), entry)

	if err != nil {
		if !IsNotFound(err) {
			t.Logf("Expected error to be a not found error, got %T:%v\n", err, err)
			t.Fail()
		}
	} else {
		t.Logf("Expected error not to be nil")
		t.Fail()
	}
}

func TestNewInMemoryRepository(t *testing.T) {
	inMemory := NewInMemoryRepository()

	expectedRepository := InMemoryRepository{
		store: make(map[string]interface{}),
	}

	if !reflect.DeepEqual(expectedRepository, inMemory) {
		t.Logf("Expected repository to equal\n%q\n\tgot\n%q\n", expectedRepository, inMemory)
		t.Fail()
	}
}

func TestInMemoryRepositoryLoad(t *testing.T) {
	inMemory := InMemoryRepository{
		store: map[string]interface{}{"foo": 5},
	}

	entry, err := inMemory.Load("foo")

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	expectedEntry := 5

	if !reflect.DeepEqual(expectedEntry, entry) {
		t.Logf("Expected entry to equal\n%q\n\tgot\n%q\n", expectedEntry, entry)
		t.Fail()
	}

	// value is nil
	inMemory.store["xyz"] = nil

	entry, err = inMemory.Load("xyz")

	if err != nil {
		t.Logf("Expected error to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(nil, entry) {
		t.Logf("Expected entry to equal\n%q\n\tgot\n%q\n", nil, entry)
		t.Fail()
	}

	// No value found
	entry, err = inMemory.Load("bar")

	if entry != nil {
		t.Logf("Expected entry to be nil, got %v\n", entry)
		t.Fail()
	}

	if err == nil {
		t.Logf("Expected error not to be nil")
		t.Fail()
	}

	if !IsNotFound(err) {
		t.Logf("Expected error to be of type NotFound, got %T\n", err)
		t.Fail()
	}
}

func TestInMemoryRepositoryStore(t *testing.T) {
	inMemory := InMemoryRepository{
		store: make(map[string]interface{}),
	}

	valueToStore := 4

	err := inMemory.Store("foo", valueToStore)

	if err != nil {
		t.Logf("Expected err to be nil, got %T:%v\n", err, err)
		t.Fail()
	}

	value, ok := inMemory.store["foo"]

	if !ok {
		t.Logf("Expected value to be stored, was not\n")
		t.Fail()
	}

	if !reflect.DeepEqual(valueToStore, value) {
		t.Logf("Expected stored value to equal\n%q\n\tgot\n%q\n", valueToStore, value)
		t.Fail()
	}
}
