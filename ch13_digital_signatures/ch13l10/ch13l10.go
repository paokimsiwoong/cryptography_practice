package ch13l10

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// 주어진 메시지와 private key로 signature 생성 뒤, message.signature 형태로 반환하는 함수
func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	// ?

	h := sha256.New()
	h.Write([]byte(message))
	hashed := h.Sum(nil)
	// @@@ sha256.Sum256([]byte(message))로 한줄 생성도 가능

	// hash된 메세지와 private key로 signature 생성
	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hashed)
	if err != nil {
		return "", err
	}

	// message.signature 형태로 반환
	// // signature는 lowercase hex string으로 변환
	return message + "." + fmt.Sprintf("%x", sig), nil
}
