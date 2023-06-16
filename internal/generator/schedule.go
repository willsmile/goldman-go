package generator

import (
	"strings"
	"time"
)

type Schedule struct {
	Date   time.Time
	Option string
}

func (s Schedule) Format(f Format) string {
	layout := f.DateLayout()
	alias := f.WdayAlias()
	dateValue := s.Date.Format(layout)
	wdayValue := alias[s.Date.Weekday().String()]
	timeValue := s.Option

	dateReplaced := strings.Replace(f.ScheduleFormat(), LabelDate, dateValue, 1)
	timeReplaced := strings.Replace(dateReplaced, LabelTime, timeValue, 1)
	wdayReplaced := strings.Replace(timeReplaced, LabelWday, wdayValue, 1)

	return wdayReplaced
}
