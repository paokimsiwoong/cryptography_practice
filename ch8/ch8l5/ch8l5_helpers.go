package ch8l5

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
)

// CBC 모드로 DES 복호화를 진행하는 함수
func decrypt(key, ciphertext []byte) ([]byte, error) {
	// @@@ 복호화 과정
	// @@@ P_i = D(C_i) ⊕ C_i-1 for i = 1, 2, 3, ...
	// @@@ P_0 = D(C_0) ⊕ iv
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	// ciphertext 맨 앞 첫 블록은 iv
	iv := ciphertext[:des.BlockSize]

	// iv를 제외한 나머지가 실제 암호문
	ciphertext = ciphertext[des.BlockSize:]
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	// CBC 모드 생성 후 복호화 진행
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

// 주어진 block의 크기가 desiredSize가 될떄까지 0을 append하는 zero pad 함수
func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}
