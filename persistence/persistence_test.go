package persistence

import (
	"errors"
	"testing"
)

type customNotFoundErr string

func (c customNotFoundErr) Error() string { return string(c) }

func (c customNotFoundErr) NotFound() bool {
	return true
}

func TestIsNotFound(t *testing.T) {
	err := errors.New("NOT FOUND")

	isNotFound := IsNotFound(err)

	if isNotFound {
		t.Logf("Expected error not to be an IsNotFound error\n")
		t.Fail()
	}

	isNotFound = IsNotFound(NotFoundErr)

	if !isNotFound {
		t.Logf("Expected error to be an IsNotFound error\n")
		t.Fail()
	}

	isNotFound = IsNotFound(customNotFoundErr("NOT FOUND"))

	if !isNotFound {
		t.Logf("Expected error to be an IsNotFound error\n")
		t.Fail()
	}
}
