package ch6l1

import (
	"errors"
)

// encrypt function reads from channels, performs XOR encryption using the crypt function
func encrypt(plaintext, key []byte) ([]byte, error) {
	if len(plaintext) != len(key) {
		return nil, errors.New("plaintext and key must be the same length")
	}

	plaintextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	// 3 채널 모두 길이가 1
	// @@@ 1번, 2번, crypt go 루틴 각각 실행 후, encrypt 메인 루틴은 result for 루프 순회에 돌입 후 block
	// @@@ ==> 1번, 2번 go 루틴은 for 루프 한번 실행 후 plaintextCh, keyCh 안의 데이터가 나갈때까지 block
	// @@@ crypt go 루틴 crypt 함수에서 plaintextCh, keyCh 데이터 들어올 때까지 block,
	// @@@ 들어오면 두 채널을 비워주고 따라서 1번, 2번 go 루틴은 다음 루프 실행 후 다시 plaintextCh, keyCh 안의 데이터가 나갈때까지 block
	// @@@ crypt 함수에서 result 채널에 데이터 입력하면 encrypt 메인 루틴은 result for 루프 1회 실행 후 다시 block
	// @@@ encrypt 메인 루틴의 result for 루프가 result 채널을 비워주기 전에 crypt 함수에서 result 채널에 데이터 입력을 또하려 하면 block

	// 1번 go 루틴
	go func() {
		defer close(plaintextCh)
		for _, v := range plaintext {
			plaintextCh <- v
		}
	}()

	// 2번 go 루틴
	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	// crypt go 루틴
	go crypt(plaintextCh, keyCh, result)

	// result for 루프 순회(encrypt 함수 메인 루틴)
	res := []byte{}
	for v := range result {
		res = append(res, v)
	}

	return res, nil
}

// decrypt function performs XOR decryption using the crypt function
func decrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) != len(key) {
		return nil, errors.New("ciphertext and key must be the same length")
	}

	ciphertextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(ciphertextCh)
		for _, v := range ciphertext {
			ciphertextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(ciphertextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}
