package main

type Format struct {
	Date string            `yaml:"date"`
	Wday map[string]string `yaml:"wday"`
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

func defaultFormat() *Format {
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
	return &Format{Date: date, Wday: wday}
}
