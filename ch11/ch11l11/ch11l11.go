package ch11l11

import (
	"math/big"
)

// RSA 암호화 (m^e (mod n) 계산) 함수
func encrypt(m, e, n *big.Int) *big.Int {
	// @@@ m은 plaintext를 []byte로 변환한 후
	// @@@ big.NewInt(0).SetBytes(byteText)로 생성하고 있음
	// @@@ @@@ func (z *big.Int) SetBytes(buf []byte) *big.Int : SetBytes interprets buf as the bytes of a big-endian unsigned integer, sets z to that value, and returns z.

	// ?

	return new(big.Int).Exp(m, e, n)
}
