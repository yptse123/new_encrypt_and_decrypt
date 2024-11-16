package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// encryptAES encrypts the plaintext using AES in CFB mode
func EncryptAES(plaintext, key []byte) ([]byte, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Allocate space for the ciphertext, which includes space for the IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// The IV is stored in the first blockSize bytes of the ciphertext
	iv := ciphertext[:aes.BlockSize]

	// Generate a random IV
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Create a CFB encrypter and perform encryption
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}
