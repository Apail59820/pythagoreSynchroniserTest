package config

import (
	"encoding/json"
	"os"
	"testing"
)

func TestAppendMetadata(t *testing.T) {
	dir := t.TempDir()
	wd, _ := os.Getwd()
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {

		}
	}(wd)
	err := os.Chdir(dir)
	if err != nil {
		return
	}

	m1 := InvoiceMetadata{InvoiceID: 1, Reference: "r1", Token: "t1"}
	if err := AppendMetadata(m1); err != nil {
		t.Fatalf("append1: %v", err)
	}
	m2 := InvoiceMetadata{InvoiceID: 2, Reference: "r2", Token: "t2"}
	if err := AppendMetadata(m2); err != nil {
		t.Fatalf("append2: %v", err)
	}

	data, err := os.ReadFile(metadataFile)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	var list []InvoiceMetadata
	if err := json.Unmarshal(data, &list); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(list) != 2 || list[0] != m1 || list[1] != m2 {
		t.Fatalf("unexpected list: %+v", list)
	}
}
