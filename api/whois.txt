package handler

import (
	"fmt"
	"net/http"

	"github.com/likexian/whois"
)

func WhoisHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: domain parameter is required")
		return
	}

	result, err := whois.Whois(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, result)
}
