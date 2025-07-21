package config

import (
	"encoding/json"
	"os"
)

const stateFile = "sync_state.json"

type syncState struct {
	LastID int `json:"last_id"`
}

// LoadLastID lit l'identifiant de la derniere facture traitee depuis le fichier de statut.
func LoadLastID() int {
	b, err := os.ReadFile(stateFile)
	if err != nil {
		return 0
	}
	var s syncState
	if err := json.Unmarshal(b, &s); err != nil {
		return 0
	}
	return s.LastID
}

// SaveLastID enregistre l'identifiant de la derniere facture traitee.
func SaveLastID(id int) error {
	s := syncState{LastID: id}
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, b, 0644)
}
