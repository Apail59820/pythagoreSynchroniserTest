package config

import (
	"os"
	"testing"
)

func TestLoadLastIDNoFile(t *testing.T) {
	dir := t.TempDir()
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(dir)
	if id := LoadLastID(); id != 0 {
		t.Fatalf("expected 0, got %d", id)
	}
}

func TestSaveAndLoadLastID(t *testing.T) {
	dir := t.TempDir()
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(dir)

	if err := SaveLastID(42); err != nil {
		t.Fatalf("save: %v", err)
	}
	if id := LoadLastID(); id != 42 {
		t.Fatalf("expected 42, got %d", id)
	}
}
