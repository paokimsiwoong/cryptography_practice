package ch2l1

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	// ?
	ints := int(bits)
	// if ints > 7 {
	if ints > 7 || ints < 0 { // @@@ ints가 음수인 경우도 예외 처리 필요
		return ""
	}
	return string(base8Alphabet[ints])
}
