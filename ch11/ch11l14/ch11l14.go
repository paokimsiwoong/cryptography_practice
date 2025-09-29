package ch11l14

import (
	"math/big"
)

// Get the private exponent
// e (mod tot) 의 곱셈 역원을 반환하는 함수
func getD(e, tot *big.Int) *big.Int {
	// ?

	d := new(big.Int)

	return d.ModInverse(e, tot)
}
