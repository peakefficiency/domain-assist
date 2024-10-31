package main

import (
	"log"
	"net/http"

	handler "domain-assist/api"
	"domain-assist/utils"
)

func main() {
	// Register all handlers automatically
	utils.RegisterHandlers(
		handler.HelloHandler,
		handler.WhoisHandler,
	)

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
