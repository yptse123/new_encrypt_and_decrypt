package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"new_crypt/decrypt"
)

// Handler for decryption requests
func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		EncryptedData string `json:"encryptedData"`
		AESKey        string `json:"aesKey"`
		BlowfishKey   string `json:"blowfishKey"`
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

	// Decode the encrypted data from base64
	encryptedData, err := base64.StdEncoding.DecodeString(request.EncryptedData)
	if err != nil {
		http.Error(w, "Invalid encrypted data", http.StatusBadRequest)
		return
	}

	// Step 1: Decrypt the data using Blowfish
	decryptedBlowfish, err := decrypt.DecryptBlowfish(encryptedData, blowfishKey)
	if err != nil {
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}

	// Step 2: Decrypt the Blowfish decrypted data using AES
	decryptedAES, err := decrypt.DecryptAES(decryptedBlowfish, aesKey)
	if err != nil {
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}

	response := struct {
		DecryptedText string `json:"decryptedText"`
	}{
		DecryptedText: string(decryptedAES),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
