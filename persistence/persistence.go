package persistence

import "errors"

type Repository interface {
	Load(key string) (interface{}, error)
	Store(key string, data interface{}) error
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
