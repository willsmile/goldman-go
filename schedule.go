package main

import (
	"fmt"
	"time"
)

type Schedule struct {
	Date   time.Time
	Option string
}

func (s Schedule) Format(f Format) string {
	layout := f.DateLayout()
	alias := f.WdayAlias()
	wday := s.Date.Weekday().String()

	return fmt.Sprintf("%s(%s) %s", s.Date.Format(layout), alias[wday], s.Option)
}
