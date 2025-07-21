package config

import (
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
