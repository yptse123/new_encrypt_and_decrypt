package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

// decryptAES decrypts the ciphertext using AES in CFB mode
func DecryptAES(ciphertext, key []byte) ([]byte, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the IV from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	// The actual ciphertext follows the IV
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a CFB decrypter and perform decryption
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
