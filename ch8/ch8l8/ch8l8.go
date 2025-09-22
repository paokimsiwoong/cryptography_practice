package ch8l8

import "crypto/rand"

// var count = 0

func generateIV(length int) ([]byte, error) {
	// count++

	// return []byte{byte(count)}, nil
	iv := make([]byte, length)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}

	return iv, nil
}
