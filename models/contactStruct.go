package models

// Structure pour recevoir les données du formulaire de contact
type ContactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Configuration SMTP
type SMTPConfig struct {
	Host           string
	Port           string
	Username       string
	Password       string
	RecipientEmail string
}
