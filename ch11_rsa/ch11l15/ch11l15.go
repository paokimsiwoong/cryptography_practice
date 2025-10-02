package ch11l15

import (
	"math/big"
)

// RSA 복호화 (c^d (mod n) 계산)함수
func decrypt(c, d, n *big.Int) *big.Int {
	// ?

	decrypted := new(big.Int)
	decrypted.Exp(c, d, n)

	return decrypted
}
