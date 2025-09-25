package ch9l1

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) { // @@@ named return value ==> 함수 시작할 떄 정의되므로 따로 정의할 필요 없이 사용 가능
	// ?

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if mode.NonceSize() != len(nonce) {
		return nil, errors.New("invalid nounce size")
	}

	plaintext, err = mode.Open(nil, nonce, ciphertext, nil)
	// @@@ ch9l1_helpers.go의 encrypt 함수에서 .Seal로 암호화 할 때 dst, additionalData에 nil 입력하므로 동일하게 여기서도 nil 입력
	// @@@ @@@ _, err = mode.Open(plaintext, nonce, ciphertext, nil)와 같이
	// @@@ @@@ dst에 plaintext를 넣으면 plaintext의 cap이 0이기 때문에 plaintext에는 어떠한 결과도 들어가지 않는다
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
