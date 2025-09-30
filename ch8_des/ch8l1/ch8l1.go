package ch8l1

func feistel(msg []byte, roundKeys [][]byte) []byte {
	// ?

	// msg를 절반으로 나누어 좌우 두 부분(lhs, rhs)으로 구분
	half := len(msg) / 2
	lhs := msg[:half]
	rhs := msg[half:]
	var oldRhs []byte

	for _, roundKey := range roundKeys {
		// 이전 오른쪽 절반을 임시 저장
		oldRhs = rhs
		// 새로운 오른쪽 절반은 xor(lhs, hash(rhs, roundKey))로 계산
		rhs = xor(lhs, hash(rhs, roundKey, len(lhs)))
		// 왼쪽 절반은 이전의 오른쪽 절반으로 교체(스왑)
		lhs = oldRhs
	}

	// 최종 결과는 오른쪽+왼쪽 순서로 결합
	return append(rhs, lhs...)
}
