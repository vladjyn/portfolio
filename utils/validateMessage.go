package utils

import (
	"fmt"
	"portfolio/models"
	"strings"
)

// Valider les données du message
func ValidateMessage(msg models.ContactMessage) error {
	// Vérifier que les champs ne sont pas vides
	if strings.TrimSpace(msg.Name) == "" {
		return fmt.Errorf("le nom est requis")
	}
	if strings.TrimSpace(msg.Email) == "" {
		return fmt.Errorf("l'email est requis")
	}
	if strings.TrimSpace(msg.Message) == "" {
		return fmt.Errorf("le message est requis")
	}

	// Validation basique de l'email
	if !strings.Contains(msg.Email, "@") || !strings.Contains(msg.Email, ".") {
		return fmt.Errorf("adresse email invalide")
	}

	// Vérifier la longueur maximale
	if len(msg.Name) > 60 {
		return fmt.Errorf("le nom est trop long (maximum 60 caractères)")
	}
	if len(msg.Email) > 60 {
		return fmt.Errorf("l'email est trop long (maximum 60 caractères)")
	}
	if len(msg.Message) > 5000 {
		return fmt.Errorf("le message est trop long (maximum 5000 caractères)")
	}

	return nil
}
