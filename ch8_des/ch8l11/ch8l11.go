package ch8l11

import "errors"

func sBox(b byte) (byte, error) {
	// ?
	// 0 2 1 3
	// 2 0 3 1
	// 1 3 0 2
	// 3 1 2 0
	// lookup table 생성
	m := [][]byte{{0b00, 0b10, 0b01, 0b11}, {0b10, 0b00, 0b11, 0b01}, {0b01, 0b11, 0b00, 0b10}, {0b11, 0b01, 0b10, 0b00}}

	// 3번쨰, 4번째 자리만 남기기 위해 0b1100 mask와 and 연산
	mask := byte(0b1100)
	// mask한 후 bit shift로 1, 2 번째 자리로 이동
	idx := int((b & mask) >> 2)
	if idx < 0 || idx > 3 {
		return 0, errors.New("invalid input")
	}

	// 1번째, 2번째 자리만 남기기 위해 0b0011 mask와 and 연산
	mask = byte(0b0011)
	jdx := int((b & mask))
	if jdx < 0 || jdx > 3 {
		return 0, errors.New("invalid input")
	}

	return m[idx][jdx], nil
}
