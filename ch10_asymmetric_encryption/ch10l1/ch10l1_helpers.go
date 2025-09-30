package ch10l1

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

// keysArePaired verifies if the public and private keys are paired using ECDSA.
func keysArePaired(pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey) bool {
	msg := "a test message"
	// SHA-256 해시를 계산하여 hash 변수에 저장.
	hash := sha256.Sum256([]byte(msg))

	// private key와 ecdsa.SignASN1로 해시에 대한 서명(sig)을 생성
	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		return false
	}

	// public key로 ecdsa.VerifyASN1를 사용해, 서명(sig)이 위 해시와 공개키로 유효한지 검증
	return ecdsa.VerifyASN1(pubKey, hash[:], sig)
}
