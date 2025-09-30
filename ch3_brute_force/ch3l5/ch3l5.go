package ch3l5

import "errors"

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	// ?

	// go 1.22부터 for i := range (1 << 24) {} 가능
	for i := 0; i < (1 << 24); i++ {
		b := intToBytes(i)
		c := crypt(encrypted, b)
		if decrypted == string(c) {
			return b, nil
		}
	}

	return []byte{}, errors.New("key not found")
}

// main.go에서 사용가능하도록 public화
func FindKey(encrypted []byte, decrypted string) ([]byte, error) {
	// ?

	// go 1.22부터 for i := range (1 << 24) {} 가능
	for i := 0; i < (1 << 10); i++ {
		b := intToBytes(i)
		c := crypt(encrypted, b)
		if decrypted == string(c) {
			return b, nil
		}
	}

	return []byte{}, errors.New("key not found")
}
