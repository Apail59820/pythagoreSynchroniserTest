package config

import (
	"encoding/json"
	"os"
)

const metadataFile = "invoice_metadata.json"

type InvoiceMetadata struct {
	InvoiceID int    `json:"invoice_id"`
	Reference string `json:"reference"`
	Token     string `json:"token"`
}

// LoadMetadata retourne toutes les métadonnées enregistrées.
func LoadMetadata() ([]InvoiceMetadata, error) {
	b, err := os.ReadFile(metadataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var list []InvoiceMetadata
	if err := json.Unmarshal(b, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// AppendMetadata ajoute des métadonnées à l'historique local.
func AppendMetadata(meta InvoiceMetadata) error {
	var list []InvoiceMetadata
	if b, err := os.ReadFile(metadataFile); err == nil {
		json.Unmarshal(b, &list)
	}
	list = append(list, meta)
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(metadataFile, data, 0644)
}
