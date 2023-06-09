package generator

import "strings"

const (
	LabelDate = "%{date}"
	LabelTime = "%{time}"
	LabelWday = "%{wday}"
)

type Format struct {
	Schedule string            `yaml:"schedule"`
	Date     string            `yaml:"date"`
	Wday     map[string]string `yaml:"wday"`
}

func (f Format) ScheduleFormat() string {
	if f.Schedule != "" {
		return f.Schedule
	} else {
		return defaultFormat().Schedule
	}
}

func (f Format) DateLayout() string {
	if f.Date != "" {
		return f.Date
	} else {
		return defaultFormat().Date
	}
}

func (f Format) WdayAlias() map[string]string {
	if len(f.Wday) != 0 {
		return f.Wday
	} else {
		return defaultFormat().Wday
	}
}

func NewReplacer(s Schedule, f Format) *strings.Replacer {
	layout := f.DateLayout()
	alias := f.WdayAlias()
	dateValue := s.Date.Format(layout)
	wdayValue := alias[s.Date.Weekday().String()]
	timeValue := s.Option
	oldnew := []string{LabelDate, dateValue, LabelTime, timeValue, LabelWday, wdayValue}

	return strings.NewReplacer(oldnew...)
}

func defaultFormat() *Format {
	schedule := "%{date}(%{wday}) %{time}"
	date := "2006-01-02"
	wday := map[string]string{
		"Monday":    "Mon",
		"Tuesday":   "Tue",
		"Wednesday": "Wed",
		"Thursday":  "Thu",
		"Friday":    "Fri",
		"Saturday":  "Sat",
		"Sunday":    "Sun",
	}
	return &Format{Schedule: schedule, Date: date, Wday: wday}
}
