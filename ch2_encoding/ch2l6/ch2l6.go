package ch2l6

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	// ?
	split := strings.Split(s, ":")

	joined := strings.Join(split, "")

	return hex.DecodeString(joined)
}

// func getHexBytes(s string) ([]byte, error) {
// 	// ?
// 	result := []byte{}
// 	split := strings.Split(s, ":")

// 	for _, v := range split {
// 		decoded, err := hex.DecodeString(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, decoded...)
// 		// @@@ hex.DecodeString는 []byte를 반환
// 		// @@@ []byte에 append하려면 ...로 슬라이스를 unpack해야 한다
// 	}

// 	return result, nil
// }
