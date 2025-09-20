package ch8l1

import (
	"crypto/sha256"
)

// ch8l1_test의 generateRoundKeys함수로 생성된 라운드키의 순서를 뒤집는 함수 (복호화 과정은 암호화 과정과 라운드 키 사용 순서가 반대)
func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// 슬라이스 s의 길이가 n일 때,
		// i=0, j=n-1 부터 시작해서 s[0] = s[n-1], s[n-1] = s[0] 과 같이 순서를 뒤집어준다
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// xor 연산 함수
func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

// outputLength should be equal to or less than the length
// of the left half when used in feistel so that the XOR
// has sufficient bytes to operate on
// 두 개의 바이트 배열을 합친 뒤 SHA-256 해시를 계산하고,
// 출력 길이를 outputLength로 줄여서 반환하는 함수
func hash(first, second []byte, outputLength int) []byte {
	// SHA-256 해싱을 수행할 해시 객체(hash.Hash) 생성
	h := sha256.New()
	// 두 byte 슬라이스를 합치고 합친 데이터를 해시 객체에 입력
	h.Write(append(first, second...))
	// h.Sum(nil)은 현재까지 해싱된 데이터를 바이트 배열로 출력
	// sha256 해시는 32바이트(256 bits = 32 bytes) ==> outputlength 길이까지만 반환
	return h.Sum(nil)[:outputLength]
}
