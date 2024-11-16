package main

import (
	"fmt"
	"log"
	"net/http"

	"new_crypt/handler"
)

func main() {
	// handle encrypt requests
	http.HandleFunc("/encrypt", corsMiddleware(handler.EncryptHandler))
	// handle decrypt requests
	http.HandleFunc("/decrypt", corsMiddleware(handler.DecryptHandler))

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// corsMiddleware is a middleware that adds CORS headers to the response
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// Return a new handler function
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}
