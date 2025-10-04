package ch13l10

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

// signature 검증 함수
func verifyECDSAMessage(token string, publicKey *ecdsa.PublicKey) error {
	// message.signature 형태이므로 .으로 split
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return errors.New("invalid token sections")
	}

	// lowercase hexstring으로 되어있는 signature를 []byte형태로 복원
	sig, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}

	// signature를 생성할 때, 원 message가 아니라 hashed message를 사용했으므로
	// hashed message 다시 생성
	hash := sha256.Sum256([]byte(parts[0]))

	// hashed message, sig, public key를 가지고 검증
	valid := ecdsa.VerifyASN1(publicKey, hash[:], sig)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}
