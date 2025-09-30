package ch11l1

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// encrypt encrypts the given plaintext using RSA-OAEP and the provided public key.
func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	// ?

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, msg, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}
