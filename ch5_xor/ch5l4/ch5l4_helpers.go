package ch5l4

// encrypt function uses the crypt function for XOR encryption
func encrypt(plaintext, key []byte) []byte {
	return crypt(plaintext, key)
}

// decrypt function uses the crypt function for XOR decryption
func decrypt(ciphertext, key []byte) []byte {
	return crypt(ciphertext, key)
}

// @@@ XOR 연산은 자기 자신이 inverse: a = (a ⊕ b) ⊕ b
// // @@@ a: plaintext
// // @@@ b: key
// // @@@ a ⊕ b: ciphertext
