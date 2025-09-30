package ch11l1

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// decrypt decrypts the given ciphertext using RSA-OAEP and the provided private key.
func decrypt(privKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// genKeys generates a new RSA key pair (public and private keys).
func genKeys() (pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return &privateKey.PublicKey, privateKey, nil
}
