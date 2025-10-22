package utils

import (
	"fmt"
	"net/smtp"
	"portfolio/models"
)

// Envoyer l'email via SMTP
func SendEmail(msg models.ContactMessage) error {
	//Chargement des données SMTP
	smtpConfig := GetSMTPConfig()

	// Configuration de l'authentification SMTP
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	// Construction du sujet et du corps de l'email
	subject := fmt.Sprintf("Portfolio - Nouveau message de %s", msg.Name)

	body := fmt.Sprintf(
		"Vous avez reçu un nouveau message depuis votre portfolio:\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"De: %s\n"+
			"Email: %s\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"+
			"Message:\n\n%s\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"Répondre à: %s",
		msg.Name, msg.Email, msg.Message, msg.Email,
	)

	// Construction des headers de l'email
	headers := make(map[string]string)
	headers["From"] = smtpConfig.Username
	headers["To"] = smtpConfig.RecipientEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=utf-8"
	headers["Reply-To"] = msg.Email

	// Assembler le message complet
	emailMsg := ""
	for k, v := range headers {
		emailMsg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	emailMsg += "\r\n" + body

	// Envoyer l'email via SMTP
	err := smtp.SendMail(
		smtpConfig.Host+":"+smtpConfig.Port,
		auth,
		smtpConfig.Username,
		[]string{smtpConfig.RecipientEmail},
		[]byte(emailMsg),
	)

	return err
}
