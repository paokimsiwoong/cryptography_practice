package ch12l1

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

// ?

// 초기화된 hasher 구조체의 포인터를 반환하는 함수
func newHasher() *hasher {
	new := new(hasher)
	new.hash = sha256.New()
	return new
}

// 주어진 string을 []byte로 변환한 뒤, 해쉬 계산에 쓰일 input 데이터로 hasher 구조체에 저장하는 메소드
func (h *hasher) Write(input string) (int, error) {
	return h.hash.Write([]byte(input))
}

// h에 지금까지 입력된 input 데이터들의 해시 계산 결과를 hex(16진수) 문자열로 출력하는 메소드
func (h *hasher) GetHex() string {
	// Sum(b []byte)는 h.Write() 로 현재까지 입력된 input 데이터의 해시 계산 결과를 반환
	// // @@@ Sum(b []byte)는 인자로 받은 b 뒤에 해시 결과를 붙여 반환하며, nil을 주면 새 슬라이스 반환
	result := h.hash.Sum(nil)

	return fmt.Sprintf("%x", result)
}
