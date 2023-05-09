package main

import (
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func NewDateRange(s string, e string, n int) (DateRange, error) {
	var (
		start time.Time
		end   time.Time
		err   error
	)

	if s == "" || (e == "" && n == 0) {
		return DateRange{}, ErrInvalidArgument
	}

	start, err = time.Parse(time.DateOnly, s)
	if err != nil {
		return DateRange{}, ErrInvalidDateFormat
	}

	if e != "" {
		end, err = time.Parse(time.DateOnly, e)
	} else if n > 0 {
		end = start.AddDate(0, 0, n)
	}
	if err != nil {
		return DateRange{}, ErrInvalidDateFormat
	}

	if !start.Before(end) {
		return DateRange{}, ErrInvalidDateRange
	}

	return DateRange{Start: start, End: end}, nil
}

func (drg DateRange) DateList() []time.Time {
	var list []time.Time

	for date := drg.Start; !date.After(drg.End); date = date.AddDate(0, 0, 1) {
		list = append(list, date)
	}

	return list
}