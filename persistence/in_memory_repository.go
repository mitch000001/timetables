package persistence

import (
	"fmt"

	"github.com/mitch000001/timetables"
)

func NewInMemoryBillingPeriodRepository() BillingPeriodRepository {
	return InMemoryBillingPeriodRepository{
		store: make(map[ID]timetables.BillingPeriod),
	}
}

type InMemoryBillingPeriodRepository struct {
	store map[ID]timetables.BillingPeriod
}

func (i InMemoryBillingPeriodRepository) Save(billingPeriod timetables.BillingPeriod) (ID, error) {
	id := fmt.Sprintf("%p", &billingPeriod)
	i.store[idString(id)] = billingPeriod
	return idString(id), nil
}

func (i InMemoryBillingPeriodRepository) Load(id ID) (timetables.BillingPeriod, error) {
	var res timetables.BillingPeriod
	res, ok := i.store[id]
	if ok {
		return res, nil
	}
	return res, NotFoundErr
}

func (i InMemoryBillingPeriodRepository) Update(id ID, period timetables.BillingPeriod) error {
	_, ok := i.store[id]
	if !ok {
		return NotFoundErr
	}
	i.store[id] = period
	return nil
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{
		store: make(map[string]interface{}),
	}
}

type InMemoryRepository struct {
	store map[string]interface{}
}

func (i InMemoryRepository) Load(key string) (interface{}, error) {
	res, ok := i.store[key]
	if ok {
		return res, nil
	}
	return nil, NotFoundErr
}

func (i InMemoryRepository) Store(key string, data interface{}) error {
	i.store[key] = data
	return nil
}
