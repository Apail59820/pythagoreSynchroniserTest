package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

// Load charge les variables d'environnement depuis un fichier .env s'il existe.
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, on utilise les variables d'environnement système")
	}
}

// SyncInterval retourne la durée entre deux synchronisations.
// La valeur peut être définie via la variable SYNC_INTERVAL en secondes.
func SyncInterval() time.Duration {
	if v := os.Getenv("SYNC_INTERVAL"); v != "" {
		if secs, err := strconv.Atoi(v); err == nil && secs > 0 {
			return time.Duration(secs) * time.Second
		}
		log.Printf("intervalle invalide %q, utilisation de la valeur par défaut", v)
	}
	return 10 * time.Second
}

// StartDate lit la date de début dans la variable START_DATE.
// Le format attendu est YYYY-MM-DD.
func StartDate() (time.Time, error) {
	v := os.Getenv("START_DATE")
	if v == "" {
		return time.Time{}, fmt.Errorf("START_DATE non definie")
	}
	t, err := time.Parse("2006-01-02", v)
	if err != nil {
		return time.Time{}, fmt.Errorf("START_DATE invalide: %w", err)
	}
	return t, nil
}

// EndDate lit la date de fin dans la variable END_DATE.
// Le format attendu est YYYY-MM-DD.
func EndDate() (time.Time, error) {
	v := os.Getenv("END_DATE")
	if v == "" {
		return time.Time{}, fmt.Errorf("END_DATE non definie")
	}
	t, err := time.Parse("2006-01-02", v)
	if err != nil {
		return time.Time{}, fmt.Errorf("END_DATE invalide: %w", err)
	}
	return t, nil
}
