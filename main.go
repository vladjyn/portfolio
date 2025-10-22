package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"portfolio/utils"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API endpoint pour le formulaire de contact
	http.HandleFunc("/api/contact", utils.HandleContact)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5280"
	}

	fmt.Printf("\nServeur portfolio démarré avec succès!\n")
	fmt.Printf("Accédez au site: http://localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
