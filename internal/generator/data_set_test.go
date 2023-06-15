package generator

import (
	"reflect"
	"testing"
	"time"

	"github.com/willsmile/goldman-go/internal/date"
)

func TestGenerate(t *testing.T) {
	ds := DataSet{
		"Weekday": Data{"12:00~13:00"},
		"Monday":  Data{"16:00~17:00"},
	}
	m1, _ := time.Parse(time.DateOnly, "2023-05-01")
	m2, _ := time.Parse(time.DateOnly, "2023-05-02")
	m3, _ := time.Parse(time.DateOnly, "2023-05-03")
	drg := date.DateRange{
		Start: m1,
		End:   m3,
	}
	s1 := Schedule{Date: m1, Option: "12:00~13:00"}
	s2 := Schedule{Date: m1, Option: "16:00~17:00"}
	s3 := Schedule{Date: m2, Option: "12:00~13:00"}
	s4 := Schedule{Date: m3, Option: "12:00~13:00"}
	expected := []Schedule{s1, s2, s3, s4}

	result := ds.Generate(drg)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Generate a slice of Schedule, expected %s, got %s", expected, result)
	}
}

func TestGenerate_ByEmptyDataSet(t *testing.T) {
	ds := DataSet{}
	m1, _ := time.Parse(time.DateOnly, "2023-05-01")
	m3, _ := time.Parse(time.DateOnly, "2023-05-03")
	drg := date.DateRange{
		Start: m1,
		End:   m3,
	}

	result := ds.Generate(drg)

	if len(result) != 0 {
		t.Fatalf("Generate a slice of Schedule, expected none Schedule, got %s", result)
	}
}
