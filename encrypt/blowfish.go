package encrypt

import (
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/blowfish"
)

// encryptBlowfish encrypts the plaintext using Blowfish in CFB mode
func EncryptBlowfish(plaintext, key []byte) ([]byte, error) {
	// Create a new Blowfish cipher block using the provided key
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Allocate space for the ciphertext, which includes space for the IV
	// Blowfish block size is 8 bytes
	ciphertext := make([]byte, blowfish.BlockSize+len(plaintext))
	// The IV is stored in the first blockSize bytes of the ciphertext
	iv := ciphertext[:blowfish.BlockSize]

	// Generate a random IV
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Create a CFB encrypter and perform encryption
	encrypt := cipher.NewCFBEncrypter(block, iv)
	encrypt.XORKeyStream(ciphertext[blowfish.BlockSize:], plaintext)

	return ciphertext, nil
}
