package decrypt

import (
	"crypto/cipher"

	"golang.org/x/crypto/blowfish"
)

// decryptBlowfish decrypts the ciphertext using Blowfish in CFB mode
func DecryptBlowfish(ciphertext, key []byte) ([]byte, error) {
	// Create a new Blowfish cipher block using the provided key
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the IV from the beginning of the ciphertext
	iv := ciphertext[:blowfish.BlockSize]
	// The actual ciphertext follows the IV
	ciphertext = ciphertext[blowfish.BlockSize:]

	// Create a CFB decrypter and perform decryption
	decrypt := cipher.NewCFBDecrypter(block, iv)
	decrypt.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
