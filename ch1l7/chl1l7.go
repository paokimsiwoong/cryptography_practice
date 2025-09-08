package ch1l7

import (
	"crypto/aes"
	"crypto/cipher"
)

//	func keyToCipher(key string) (cipher.Block, error) {
//		return aes.NewCipher(key)
//	}
func keyToCipher(key string) (cipher.Block, error) {
	return aes.NewCipher([]byte(key))
}
