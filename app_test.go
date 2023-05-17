package main

import (
	"os"
	"testing"
)

func TestAppRun_MinimumArgs(t *testing.T) {
	os.Setenv("GOLDMAN_GO_PATH", "./testdata/config.yml")
	args := []string{"./goldman-go", "g"}
	err := New().Run(args)
	if err != nil {
		t.Fatalf("App run with minimum args, expected none error, got %s", err)
	}
}

func TestAppRun_OftenUsedArgs(t *testing.T) {
	os.Setenv("GOLDMAN_GO_PATH", "./testdata/config.yml")
	args := []string{"./goldman-go", "-s", "2023-05-01", "-d", "10", "g"}
	err := New().Run(args)
	if err != nil {
		t.Fatalf("App run with often used args, expected none error, got %s", err)
	}
}

func TestAppRun_ConfigArgs(t *testing.T) {
	args := []string{"./goldman-go", "-c", "./testdata/config.yml", "g"}
	err := New().Run(args)
	if err != nil {
		t.Fatalf("App run with config args, expected none error, got %s", err)
	}
}
