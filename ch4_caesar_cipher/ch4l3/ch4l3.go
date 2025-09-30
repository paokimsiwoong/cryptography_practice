package ch4l3

import (
	"strings"
)

func encrypt(plaintext string, key int) string {
	// ?
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	// ?
	return crypt(ciphertext, -key)
	// @@@ 복호화는 key값에 * -1
}

func crypt(text string, key int) string {
	// ?

	crypted := ""

	for _, c := range text {
		crypted += getOffsetChar(c, key)
	}

	return crypted
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	// ?
	// strings.IndexRune은 문자열에서 특정 rune이 처음 등장하는 인덱스를 반환한다
	idx := strings.IndexRune(alphabet, c)
	if idx == -1 {
		return ""
	}

	shifted := (idx + offset) % len(alphabet)

	// fmt.Println(shifted)

	if shifted < 0 {
		shifted += len(alphabet)
	}

	result := alphabet[shifted]

	return string(result)
}
