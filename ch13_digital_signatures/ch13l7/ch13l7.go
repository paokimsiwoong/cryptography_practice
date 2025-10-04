package ch13l7

import (
	"crypto/sha256"
	"fmt"
)

// sha256(keyFirstHalf + sha256(keySecondHalf + message)) 형태의 toy hmac 계산 함수
func hmac(message, key string) string {
	// ?
	half := len(key) / 2

	firstHalf := key[:half]
	secondHalf := key[half:]

	h := sha256.New()
	h.Write([]byte(secondHalf + message))
	firstRound := h.Sum(nil)

	// .Write로 쌓인 데이터를 리셋
	h.Reset()
	// @@@ 리셋하지 않으면 hash(firstHalf + firstRound)가 되지 않고
	// @@@ hash(secondHalf + message + firstHalf + firstRound)가 된다
	h.Write(append([]byte(firstHalf), firstRound...))
	secondRound := h.Sum(nil)

	return fmt.Sprintf("%x", secondRound)
}
