package ch11l5

import (
	"math/big"
)

// 주어진 keysize를 가지는 두개의 소수 p, q 생성 및 반환하는 함수
func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	// ?

	p, err := getBigPrime(keysize)
	if err != nil {
		return nil, nil
	}

	q, err := getBigPrime(keysize)
	if err != nil {
		return nil, nil
	}

	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	// ?

	// p * q를 저장할 big.Int container 생성
	n := new(big.Int)

	// (z *big.Int).Mul(x, y)는 x * y의 결과를 z에 저장하고 z를 반환
	return n.Mul(p, q)
}
