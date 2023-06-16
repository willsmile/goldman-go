package generator

import (
	"time"
)

type Schedule struct {
	Date   time.Time
	Option string
}

func (s Schedule) Format(f Format) string {
	replacer := NewReplacer(s, f)

	return replacer.Replace(f.ScheduleFormat())
}
