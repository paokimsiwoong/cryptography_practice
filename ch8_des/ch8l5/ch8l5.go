package ch8l5

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

// CBC 모드로 DES 암호화를 진행하는 함수
func encrypt(key, plaintext []byte) ([]byte, error) {
	// ?

	block, err := des.NewCipher(key)
	// des block의 길이는 8바이트 고정
	if err != nil {
		return nil, err
	}

	padded := padMsg(plaintext, block.BlockSize())

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	// @@@ ciphertext의 길이는 원래 plaintext 길이 + iv(initialization vector) 길이로 생성
	ciphertext := make([]byte, block.BlockSize()+len(padded))
	// @@@ iv는 block.BlockSize()와 크기가 동일해야 한다
	iv := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 암호화를 진행하는 cipher.BlockMode 객체를 cipher.Block과 iv로 생성
	mode := cipher.NewCBCEncrypter(block, iv)
	// @@@ CBC(Cipher Block Chaining)모드는 각 라운드의 i번째 plain text block(전 라운드 결과를 block으로 나눈 뒤 i번째 블록)에 해당 라운드 block cipher를 진행하기 전에
	// @@@ 이전 i-1번째 단계의 결과(i-1번째 cipher text block)를 XOR 연산한다
	// @@@ (최초 plain text block의 경우 iv를 XOR 연산)
	// @@@ @@@ C_i = E(P_i ⊕ C_i-1) for i = 1, 2, 3, ...
	// @@@ @@@ C_0 = E(P_0 ⊕ iv)
	// 암호화 진행
	mode.CryptBlocks(ciphertext[block.BlockSize():], padded)
	// mode에 iv는 이미 저장되어 있으므로 dst에는 ciphertext에서 iv를 제외한 부분만 입력

	return ciphertext, nil
	// @@@ 반환할 떄는 암호화된 부분 앞에 iv를 반드시 같이 포함해서 반환해야한다
	// @@@ (복호화 과정에서 iv가 다시 사용됨)
	// @@@ @@@ P_i = D(C_i) ⊕ C_i-1 for i = 1, 2, 3, ...
	// @@@ @@@ P_0 = D(C_0) ⊕ iv
}

// blocksize 길이로 plaintext를 여러개의 block으로 나누었을 때 마지막 block의 길이가 blocksize와 같도록 zero padding을 하는 함수
func padMsg(plaintext []byte, blockSize int) []byte {
	// ?

	lastBlockLength := len(plaintext) % blockSize

	if lastBlockLength == 0 {
		return plaintext
	}

	result := plaintext[:len(plaintext)-lastBlockLength]
	result = append(result, padWithZeros(plaintext[len(plaintext)-lastBlockLength:], blockSize)...)

	return result
}
