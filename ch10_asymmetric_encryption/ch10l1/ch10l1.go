package ch10l1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	// ?

	// private key 생성
	privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// private key와 한 쌍인 public key는 .PublicKey 필드에 들어 있음
	pubKey = &privKey.PublicKey

	return pubKey, privKey, nil
}
