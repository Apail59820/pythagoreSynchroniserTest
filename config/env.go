package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aucun fichier .env trouvé, on utilise les variables d'environnement système")
	}
}
