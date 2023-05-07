package main

import (
	"errors"
	"testing"
	"time"
)

func TestNewDateRange_StartDateAndEndDate(t *testing.T) {
	r, err := NewDateRange("2023-05-01", "2023-05-07", 0)
	if err != nil {
		t.Fatalf("NewDateRange(\"2023-05-01\", \"2023-05-07\", 0), expected none error, got %s", err)
	}
	start, _ := time.Parse(time.DateOnly, "2023-05-01")
	if !r.Start.Equal(start) {
		t.Fatalf("Start date of Date Range should be %s but %s", start, r.Start)
	}
	end, _ := time.Parse(time.DateOnly, "2023-05-07")
	if !r.End.Equal(end) {
		t.Fatalf("End date of Date Range should be %s but %s", end, r.End)
	}
}

func TestNewDateRange_StartDateAndDay(t *testing.T) {
	r, err := NewDateRange("2023-05-01", "", 6)
	if err != nil {
		t.Fatalf("NewDateRange(\"2023-05-01\", \"\", 6), expected none error, got %s", err)
	}
	start, _ := time.Parse(time.DateOnly, "2023-05-01")
	if !r.Start.Equal(start) {
		t.Fatalf("Start date of Date Range should be %s but %s", start, r.Start)
	}
	end, _ := time.Parse(time.DateOnly, "2023-05-07")
	if !r.End.Equal(end) {
		t.Fatalf("End date of Date Range should be %s but %s", end, r.End)
	}
}

func TestNewDateRange_StartDateAndEndDateAndDay(t *testing.T) {
	r, err := NewDateRange("2023-05-01", "2023-05-07", 20)
	if err != nil {
		t.Fatalf("NewDateRange(\"2023-05-01\", \"2023-05-07\", 20), expected none error, got %s", err)
	}
	start, _ := time.Parse(time.DateOnly, "2023-05-01")
	if !r.Start.Equal(start) {
		t.Fatalf("Start date of Date Range should be %s but %s", start, r.Start)
	}
	end, _ := time.Parse(time.DateOnly, "2023-05-07")
	if !r.End.Equal(end) {
		t.Fatalf("End date of Date Range should be %s but %s", end, r.End)
	}
}

func TestNewDateRange_InvalidDateRange(t *testing.T) {
	_, err := NewDateRange("2023-05-01", "2023-04-30", 0)
	if !errors.Is(err, ErrInvalidDateRange) {
		t.Fatalf("NewDateRange(\"2023-05-01\", \"2023-04-30\", 0), expected %s, got %s", ErrInvalidDateRange, err)
	}
}

func TestNewDateRange_NoneOfArgument(t *testing.T) {
	_, err := NewDateRange("", "", 0)
	if !errors.Is(err, ErrInvalidArgument) {
		t.Fatalf("NewDateRange(\"\", \"\", 0), expected %s, got %s", ErrInvalidArgument, err)
	}
}

func TestNewDateRange_OnlyArgumentOfStartDate(t *testing.T) {
	_, err := NewDateRange("2023-05-01", "", 0)
	if !errors.Is(err, ErrInvalidArgument) {
		t.Fatalf("NewDateRange(\"2023-05-01\", \"\", 0), expected %s, got %s", ErrInvalidArgument, err)
	}
}

func TestNewDateRange_OnlyArgumentOfEndDate(t *testing.T) {
	_, err := NewDateRange("", "2023-05-01", 0)
	if !errors.Is(err, ErrInvalidArgument) {
		t.Fatalf("NewDateRange(\"\", \"2023-05-01\", 0), expected %s, got %s", ErrInvalidArgument, err)
	}
}

func TestNewDateRange_OnlyArgumentOfDay(t *testing.T) {
	_, err := NewDateRange("", "", 6)
	if !errors.Is(err, ErrInvalidArgument) {
		t.Fatalf("NewDateRange(\"\", \"\", 6), expected %s, got %s", ErrInvalidArgument, err)
	}
}

func TestNewDateRange_WithoutArgumentOfStartDate(t *testing.T) {
	_, err := NewDateRange("", "2023-05-01", 6)
	if !errors.Is(err, ErrInvalidArgument) {
		t.Fatalf("NewDateRange(\"\", \"2023-05-01\", 6), expected %s, got %s", ErrInvalidArgument, err)
	}
}

func TestNewDateRange_InvalidFormatOfStartDate(t *testing.T) {
	_, err := NewDateRange("Invalid Date Format", "Invalid Date Format", 0)
	if !errors.Is(err, ErrInvalidDateFormat) {
		t.Fatalf("NewDateRange(\"Invalid Date Format\", \"2023-05-01\", 0), expected %s, got %s", ErrInvalidDateFormat, err)
	}
}

func TestNewDateRange_InvalidFormatOfEndDate(t *testing.T) {
	_, err := NewDateRange("Invalid Date Format", "Invalid Date Format", 0)
	if !errors.Is(err, ErrInvalidDateFormat) {
		t.Fatalf("NewDateRange(\"2023-05-01\", \"Invalid Date Format\", 0), expected %s, got %s", ErrInvalidDateFormat, err)
	}
}

func TestDateList(t *testing.T) {
	r, _ := NewDateRange("2023-05-01", "2023-05-07", 0)
	list := r.DateList()
	if len(list) != 7 {
		t.Fatalf("")
	}
}
