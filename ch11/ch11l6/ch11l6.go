package ch11l6

import (
	"crypto/rand"
	"math/big"
)

// 두 소수의 곱 n = p * q 의 ϕ(n) 계산 함수
func getTot(p, q *big.Int) *big.Int {
	// ?

	tot := new(big.Int)

	one := big.NewInt(1)

	return tot.Mul(tot.Sub(p, one), tot.Sub(q, one))
}

// (1, ϕ(n)) 사이에 ϕ(n)과 서로소인 숫자(e) 하나를 랜덤 생성 후 반환하는 함수
func getE(tot *big.Int) *big.Int {
	// ?

	for {
		// [0, tot) 범위 숫자 랜덤 생성
		e, err := rand.Int(randReader, tot)
		if err != nil {
			return nil
		}

		one := big.NewInt(1)

		// e가 1보다 크지 않으면 다시 숫자 생성으로 복귀
		if e.Cmp(one) != 1 {
			continue
		}

		// e와 tot의 gcd가 1(서로수)가 아니면 다시 숫자 생성으로 복귀
		if gcd(tot, e).Cmp(one) != 0 {
			continue
		}

		return e
	}
}
