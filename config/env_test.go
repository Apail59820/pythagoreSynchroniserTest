package config

import (
	"testing"
	"time"
)

func TestSyncIntervalDefault(t *testing.T) {
	t.Setenv("SYNC_INTERVAL", "")
	if d := SyncInterval(); d != 10*time.Second {
		t.Fatalf("expected 10s, got %v", d)
	}
}

func TestSyncIntervalValid(t *testing.T) {
	t.Setenv("SYNC_INTERVAL", "5")
	if d := SyncInterval(); d != 5*time.Second {
		t.Fatalf("expected 5s, got %v", d)
	}
}

func TestSyncIntervalInvalid(t *testing.T) {
	t.Setenv("SYNC_INTERVAL", "abc")
	if d := SyncInterval(); d != 10*time.Second {
		t.Fatalf("expected default 10s, got %v", d)
	}
}
