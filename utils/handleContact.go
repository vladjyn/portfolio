package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"portfolio/models"
)

// Handler pour traiter les requêtes du formulaire de contact
func HandleContact(w http.ResponseWriter, r *http.Request) {
	// Gérer les requêtes OPTIONS pour CORS (preflight)
	if r.Method == http.MethodOptions {
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Accepter uniquement les requêtes POST
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Configurer les headers CORS
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json")

	// Décoder les données JSON du formulaire
	var msg models.ContactMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Printf("Erreur de décodage JSON: %v", err)
		RespondWithError(w, "Erreur de lecture des données", http.StatusBadRequest)
		return
	}

	// Valider les données reçues
	if err := ValidateMessage(msg); err != nil {
		log.Printf("Validation échouée: %v", err)
		RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Envoyer l'email
	err = SendEmail(msg)
	if err != nil {
		log.Printf("Erreur d'envoi d'email: %v", err)
		RespondWithError(w, "Erreur lors de l'envoi du message. Veuillez réessayer.", http.StatusInternalServerError)
		return
	}

	// Réponse de succès
	log.Printf("Message envoyé avec succès de %s (%s)", msg.Name, msg.Email)
	RespondWithSuccess(w, "Message envoyé avec succès!")
}

// Configurer les headers CORS
func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Répondre avec une erreur JSON
func RespondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}

// Répondre avec un succès JSON
func RespondWithSuccess(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": message,
	})
}
