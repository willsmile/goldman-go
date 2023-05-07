package main

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidArgument is returned when the argument is invalid
	ErrInvalidArgument = errors.New("the argument is invalid")
	// ErrInvalidDateFormat is returned when the format of date argument is invalid
	ErrInvalidDateFormat = errors.New("the format of date argument is invalid")
	// ErrInvalidDateRange is returned when the range of date argument is invalid
	ErrInvalidDateRange = errors.New("the range of date argument is invalid")
	// ErrInvalidDay is returned when the day argument is invalid
	ErrInvalidDay = errors.New("the day argument is invalid")
)

type LoadConfigError struct {
	msg string
	err error
}

func (e *LoadConfigError) Error() string {
	return fmt.Sprintf("cannot load config file: %s (%s)", e.msg, e.err.Error())
}

func (e *LoadConfigError) Unwrap() error {
	return e.err
}
