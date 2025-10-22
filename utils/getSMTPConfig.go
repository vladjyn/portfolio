package utils

import (
	"log"
	"os"
	"portfolio/models"

	"github.com/joho/godotenv"
)

func GetSMTPConfig() models.SMTPConfig {
	// Charger le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Fichier .env non trouvé, utilisation des variables système")
	}
	// Essayer de lire depuis les variables d'environnement
	host := os.Getenv("SMTP_HOST")
	if host == "" {
		host = "smtp.gmail.com" // Serveur par défaut (Gmail)
	}

	port := os.Getenv("SMTP_PORT")
	if port == "" {
		port = "587" // Port STARTTLS standard
	}

	username := os.Getenv("SMTP_USERNAME")
	if username == "" {
		log.Println("<getSMTPconfig.go> username vide")
	}

	password := os.Getenv("SMTP_PASSWORD")
	if password == "" {
		log.Println("<getSMTPconfig.go> passeword vide")
	}

	recipient := os.Getenv("RECIPIENT_EMAIL")
	if recipient == "" {
		log.Println("<getSMTPconfig.go> recipient email vide")
	}

	return models.SMTPConfig{
		Host:           host,
		Port:           port,
		Username:       username,
		Password:       password,
		RecipientEmail: recipient,
	}
}
