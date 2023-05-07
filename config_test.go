package main

import (
	"errors"
	"io/fs"
	"testing"
)

func TestLoadConfig_NotExistingPath(t *testing.T) {
	_, err := LoadConfig("not-existing-path")
	if err != nil {
		var pathError *fs.PathError
		if !errors.As(err, &pathError) {
			t.Fatalf("LoadConfig(\"not-existing-path\"), expected fs.PathError error, got %s", err.Error())
		}
	} else {
		t.Fatalf("LoadConfig(\"not-existing-path\"), expected error, got none")
	}
}

func TestLoadConfig_ValidPath(t *testing.T) {
	_, err := LoadConfig("./testdata/config.yml")
	if err != nil {
		t.Fatalf("LoadConfig(\"./testdata/config.yml\"), expected none error, got %s", err)
	}
}
