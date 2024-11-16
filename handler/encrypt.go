package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"new_crypt/encrypt"
)

// Handler for encryption requests
func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Plaintext   string `json:"plaintext"`
		AESKey      string `json:"aesKey"`
		BlowfishKey string `json:"blowfishKey"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate AES key length (16, 24, 32 bytes)
	aesKey := []byte(request.AESKey)
	if len(aesKey) != 16 && len(aesKey) != 24 && len(aesKey) != 32 {
		http.Error(w, "Invalid AES key length. Must be 16, 24, or 32 bytes.", http.StatusBadRequest)
		return
	}

	// Validate Blowfish key length (1-56 bytes)
	blowfishKey := []byte(request.BlowfishKey)
	if len(blowfishKey) < 1 || len(blowfishKey) > 56 {
		http.Error(w, "Invalid Blowfish key length. Must be between 1 and 56 bytes.", http.StatusBadRequest)
		return
	}

	// Step 1: Encrypt the plaintext using AES
	encryptedAES, err := encrypt.EncryptAES([]byte(request.Plaintext), aesKey)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	// Step 2: Encrypt the AES encrypted data using Blowfish
	encryptedBlowfish, err := encrypt.EncryptBlowfish(encryptedAES, blowfishKey)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	// Encode the final encrypted data in base64
	response := struct {
		EncryptedData string `json:"encryptedData"`
	}{
		EncryptedData: base64.StdEncoding.EncodeToString(encryptedBlowfish),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
