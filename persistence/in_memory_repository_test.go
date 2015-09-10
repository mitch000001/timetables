package persistence

import (
	"reflect"
	"testing"
)

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
