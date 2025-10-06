package ch14l6

import (
	"crypto/rand"
	"crypto/sha256"
)

// salt 생성 함수
func generateSalt(length int) ([]byte, error) {
	// ?

	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// salt와 password를 같이 해싱하는 함수
func hashPassword(password, salt []byte) []byte {
	// ?

	h := sha256.New()
	h.Write(password)
	h.Write(salt)

	return h.Sum(nil)
}
