package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"pythagoreSynchroniser/logging"
)

const metadataDir = "data/metadata"

type InvoiceMetadata struct {
	InvoiceID int    `json:"invoice_id"`
	Reference string `json:"reference"`
	Token     string `json:"token"`
}

// LoadMetadata retourne toutes les métadonnées enregistrées.
func LoadMetadata() ([]InvoiceMetadata, error) {
	entries, err := os.ReadDir(metadataDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var list []InvoiceMetadata
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(metadataDir, e.Name()))
		if err != nil {
			return nil, err
		}
		var m InvoiceMetadata
		if err := json.Unmarshal(data, &m); err != nil {
			logging.Warnf("fichier metadata %s invalide: %v", e.Name(), err)
			continue
		}
		list = append(list, m)
	}
	return list, nil
}

// AppendMetadata ajoute des métadonnées à l'historique local.
func AppendMetadata(meta InvoiceMetadata) error {
	if err := os.MkdirAll(metadataDir, 0755); err != nil {
		return err
	}
	b, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	path := filepath.Join(metadataDir, fmt.Sprintf("%d.json", meta.InvoiceID))
	return os.WriteFile(path, b, 0644)
}
