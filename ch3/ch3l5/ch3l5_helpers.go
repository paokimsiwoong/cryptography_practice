package ch3l5

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Helper function: crypt performs XOR-based encryption/decryption
func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
		// go에서 ^는 2항일때는 XOR 연산자 1항일 때는 NOT 연산자
		// XOR : 두 값의 비트별 비교 뒤, 각 비트가 서로 다르면 1, 같으면 0 반환
		// // => 5 ^ 2 = 0101 ^ 0010 = 0111 = 7
		//  @@@ dat과 key의 각자리를 XOR 연산한 결과를 final로 반환
	}
	return final
}

// Helper function: intToBytes converts an integer to a 3-byte slice (little-endian)
func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	// binary.LittleEndian는 binary데이터를 8비트 단위로 쪼갠 뒤 가장 작은 조각부터 표시하는 방식
	// ex: 0x12345678 (16진법)의 경우 2자리 마다 8비트 ==> 78 56 34 12
	// cf: big endian 방식의 경우 큰 조각부터 표시 ==> 12 34 56 78
	if err != nil {
		return nil
	}
	bs := buf.Bytes()
	fmt.Printf("current bs: %x\n", bs)
	if len(bs) > 3 {
		// @@@ num 최대값이 2^24 - 1 이면 16진법으로 표현시 0xFFFFFF ==> len(bs)는 3이므로
		// @@@ num 최대값이 2^24 - 1 보다 커지지 않는 이상 len(bs) > 3인 경우는 일어나지 않는다
		return bs[:3]
	}
	return bs
}
