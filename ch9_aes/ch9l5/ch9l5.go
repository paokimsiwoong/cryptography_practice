package ch9l5

// nonceStrength returns the number of bits of entropy in the nonce.
func nonceStrength(nonce []byte) int {
	// ?

	// nonce 총 비트수 계산
	bytes := len(nonce)
	bits := bytes * 8

	// 2 ^ bits 계산은 비트 시프트 연산자로 계산
	return 1 << bits
}
