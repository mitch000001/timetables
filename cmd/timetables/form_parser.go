package main

import (
	"net/url"
	"strconv"
)

type FormParser struct {
	form   url.Values
	errors []error
}

func NewFormParser(form url.Values) *FormParser {
	return &FormParser{
		errors: make([]error, 0),
		form:   form,
	}
}

func (f *FormParser) SetForm(form url.Values) {
	f.form = form
	f.errors = make([]error, 0)
}

func (f *FormParser) ResetErrors() {
	f.errors = make([]error, 0)
}

func (f *FormParser) GetErrors() []error {
	return f.errors
}

func (f *FormParser) Float64(key string) float64 {
	value := f.form.Get(key)
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		f.errors = append(f.errors, err)
	}
	return num
}

func (f *FormParser) Int(key string) int {
	value := f.form.Get(key)
	num, err := strconv.Atoi(value)
	if err != nil {
		f.errors = append(f.errors, err)
	}
	return num
}
