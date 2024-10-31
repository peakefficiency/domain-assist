package handler

import (
	"fmt"
	"net/http"

	"github.com/likexian/whois"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Hardcoded domain for testing
	domain := "google.com"

	result, err := whois.Whois(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, result)
}
