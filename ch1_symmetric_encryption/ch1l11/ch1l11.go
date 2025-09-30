package ch1l11

import (
	"fmt"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	randReader := rand.New(rand.NewSource(0))

	// ?
	key := make([]byte, length)
	_, err := randReader.Read(key)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", key), nil
}
