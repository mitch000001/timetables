package persistence

import (
	"errors"

	"github.com/mitch000001/timetables"
)

type Repository interface {
	Load(key string) (interface{}, error)
	Store(data interface{}) (string, error)
}

type BillingPeriodRepository interface {
	Load(id ID) (timetables.BillingPeriod, error)
	Save(period timetables.BillingPeriod) (ID, error)
	Update(id ID, period timetables.BillingPeriod) error
}

type ID interface {
	String() string
}

type NotFound interface {
	NotFound() bool
}

var NotFoundErr = errors.New("Not found")

func IsNotFound(err error) bool {
	if nf, ok := err.(NotFound); ok {
		return nf.NotFound()
	}
	return err == NotFoundErr
}

type idString string

func (i idString) String() string {
	return string(i)
}
