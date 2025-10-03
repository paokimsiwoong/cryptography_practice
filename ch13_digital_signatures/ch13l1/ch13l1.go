package ch13l1

import (
	"crypto/sha256"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
	// ?

	h := sha256.New()
	_, err := h.Write([]byte(message))
	if err != nil {
		panic("으악")
	}

	return fmt.Sprintf("%x", h.Sum(nil)) == checksum
}
