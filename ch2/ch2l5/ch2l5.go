package ch2l5

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	// ?
	var strs []string
	// @@@ strs := []string{}

	for _, c := range b {
		strs = append(strs, fmt.Sprintf("%02x", c))
	}

	return strings.Join(strs, ":")
}

func getBinaryString(b []byte) string {
	// ?
	var strs []string

	for _, c := range b {
		strs = append(strs, fmt.Sprintf("%08b", c))
	}

	return strings.Join(strs, ":")
}

// @@@ zero pad 동적 길이 설정 * 활용
// func getHexStringWithDynamicLenPad(b []byte) string {
// 	// ?
// 	var strs []string

// 	max := findMax(b)
// 	maxLen := len(fmt.Sprintf("%x", max))

// 	for _, c := range b {
// 		strs = append(strs, fmt.Sprintf("%0*x", maxLen, c))
// 	}

// 	return strings.Join(strs, ":")
// }

// func getBinaryStringWithDynamicLenPad(b []byte) string {
// 	// ?
// 	var strs []string

// 	max := findMax(b)
// 	maxLen := len(fmt.Sprintf("%b", max))

// 	for _, c := range b {
// 		strs = append(strs, fmt.Sprintf("%0*b", maxLen, c))
// 	}

// 	return strings.Join(strs, ":")
// }

// func findMax(bs []byte) byte {
// 	max := byte(0)

// 	for _, b := range bs {
// 		if b > max {
// 			max = b
// 		}
// 	}

// 	return max
// }
