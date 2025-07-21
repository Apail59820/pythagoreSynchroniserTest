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
