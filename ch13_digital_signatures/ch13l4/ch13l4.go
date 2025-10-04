package ch13l4

import (
	"crypto/sha256"
	"fmt"
)

// func macMatches(message, key, checksum string) bool {
// 	h := sha256.New()
// 	h.Write([]byte(message))
// 	return checksum == fmt.Sprintf("%x", h.Sum(nil))
// }

func macMatches(message, key, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message + key))
	// mac == hash(message + key) 계산 후 checksum과 동일 여부 확인
	return checksum == fmt.Sprintf("%x", h.Sum(nil))
}
