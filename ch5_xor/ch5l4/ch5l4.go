package ch5l4

func crypt(plaintext, key []byte) []byte {
	// ?

	result := []byte{}

	for i, b := range plaintext {
		result = append(result, b^key[i])
		// XOR 연산
	}

	return result
}
