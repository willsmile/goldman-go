package generator

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/willsmile/goldman-go/internal/date"
)

type DataSet map[string]Data
type Data []string

func (ds DataSet) Generate(drg date.DateRange) []Schedule {
	var result []Schedule

	if len(ds) == 0 {
		return result
	}

	list := drg.DateList()

	for _, date := range list {
		data := ds.selectByDate(date).sortByStartHour()
		for _, d := range data {
			s := Schedule{Date: date, Option: d}
			result = append(result, s)
		}
	}

	return result
}

func (ds DataSet) selectByDate(t time.Time) Data {
	var data Data
	wday := t.Weekday().String()
	data = append(data, ds[wday]...)
	data = append(data, ds["Everyday"]...)
	if isWeekend(t) {
		data = append(data, ds["Weekend"]...)
	} else {
		data = append(data, ds["Weekday"]...)
	}

	return data
}

func (d Data) sortByStartHour() Data {
	sort.Slice(d, func(i, j int) bool {
		return hourOfTime(d[i]) < hourOfTime(d[j])
	})

	return d
}

func hourOfTime(s string) int {
	hourStr := strings.Split(s, ":")[0]
	hour, _ := strconv.Atoi(hourStr)

	return hour
}

func isWeekend(t time.Time) bool {
	if t.Weekday() == time.Sunday || t.Weekday() == time.Saturday {
		return true
	} else {
		return false
	}
}
