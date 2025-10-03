package ch12l10

import "math/bits"

func hash(input []byte) [4]byte {
	// ?

	// 길이 4 어레이 생성
	final := new([4]byte)
	// @@@ 배열(array)은 값 타입이며 make로는 생성불가
	// @@@ @@@ ==> make 사용 시 컴파일 에러

	for i, b := range input {
		// rotate left by 3
		b = bits.RotateLeft8(b, 3)
		// shift left by 2
		b = b << 2

		// final의 i%4 자리 값과 xor 연산 후 그 자리에 다시 저장
		final[i%4] = b ^ final[i%4]
		// @@@ final[i%4] ^= b 도 동일 결과
	}

	return *final
}
