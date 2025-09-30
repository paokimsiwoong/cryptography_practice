package ch11l14

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

// RSA 암호화 (m^e (mod n) 계산) 함수
func encrypt(m, e, n *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, e, n)
	return c
}

// 두 소수의 곱 n = p * q 의 ϕ(n) 계산 함수
func gettot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

// (1, ϕ(n)) 사이에 ϕ(n)과 서로소인 숫자(e) 하나를 랜덤 생성 후 반환하는 함수
func getE(tot *big.Int) *big.Int {
	// [2, tot) 범위 숫자 생성을 위해
	// [0, tot - 2) 사이 숫자를 생성한 후
	// + 2를 한다

	// tot - 2 생성
	totMinusTwo := new(big.Int)
	totMinusTwo.Sub(tot, big.NewInt(2))

	// [0, tot - 2) 사이 숫자를 생성
	e, _ := crand.Int(randReader, totMinusTwo)
	// + 2를 해서 e의 범위를 [2, tot) 로 변경
	e.Add(e, big.NewInt(2))

	// gcd(e, tot)이 1이 될 때까지 숫자 생성 반복
	for gcd(e, tot).Cmp(big.NewInt(1)) != 0 {
		e, _ = crand.Int(randReader, totMinusTwo)
		e.Add(e, big.NewInt(2))
	}

	// gcd(e, tot) = 1인 e 반환
	return e
}

// 주어진 keysize를 가지는 두개의 소수 p, q 생성 및 반환하는 함수
func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, _ := getBigPrime(keysize)
	q, _ := getBigPrime(keysize)
	return p, q
}

// n = p * q 계산 함수
func getN(p, q *big.Int) *big.Int {
	n := new(big.Int)
	n.Mul(p, q)
	return n
}

// 최대공약수 계산 함수
func gcd(x, y *big.Int) *big.Int {
	xCopy := new(big.Int).Set(x)
	yCopy := new(big.Int).Set(y)
	for yCopy.Cmp(big.NewInt(0)) != 0 {
		xCopy, yCopy = yCopy, xCopy.Mod(xCopy, yCopy)
		// gcd(x, y) = gcd(r, y)를 이용 (x = yq + r)
		// // (.Mod(x, y) 는 r 반환)
	}

	// for 루프 종료 후에는 xCopy = gcd(x,y), yCopy = 0이 된다
	return xCopy
}

// firstNDigits returns the first 'numDigits' digits of the big integer n.
func firstNDigits(n big.Int, numDigits int) string {
	// big.Int 타입 숫자를 문자열로 변환한 뒤, 앞에서부터 numDigits자리까지(자리수가 부족하면 전부) 보여주고, 자리수가 많으면 그 뒤에 "..."를 붙여 간략히 표현
	if len(n.String()) < numDigits {
		return fmt.Sprintf("%v", n.String())
	}
	return fmt.Sprintf("%v...", n.String()[:numDigits])
}

// 랜덤 생성에 사용할 시드 고정 randReader
var randReader = mrand.New(mrand.NewSource(0))

// getBigPrime generates a random prime number of the given size.
func getBigPrime(bits int) (*big.Int, error) {
	if bits < 2 {
		return nil, errors.New("prime size must be at least 2-bit")
	}
	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}
	// bits 수를 최소한으로 커버하는 바이트 개수만큼 생성
	bytes := make([]byte, (bits+7)/8)
	// @@@ bits/8을 하면 /는 소수점을 버리므로 올림 방식으로 바꾸기 위해 (bits+7)/8로 변경
	// @@@ @@@ ex: 15/8 = 1 ==> 실제로 필요한 바이트수 2가 안나옴
	// @@@ @@@ ex: (15+7)/8 = 2 ==> 실제로 필요한 바이트수 2가 나옴

	// 생성한 bytes를 담을 big.Int 컨테이너 생성
	p := new(big.Int)

	// 원하는 크기의 소수가 생성될 때까지 반복하는 for 루프 시작
	for {
		// 랜덤 바이트 배열 생성: 지정된 길이(bytes)에 맞춰 난수로 채움
		if _, err := io.ReadFull(randReader, bytes); err != nil {
			return nil, err
		}

		// @@@ &= 와 |=는 비트 단위 복합 대입 연산자
		// @@@ x &= y 는 x = x & y, x |= y 는 x = x | y

		// 상위 비트 마스킹: bytes의 비트 수가 특정 비트수(bits 변수)보다 클 경우, 초과 비트를 0으로 만듦
		bytes[0] &= uint8(int(1<<b) - 1)
		// @@@ b := uint(bits % 8)
		// @@@ bits = 8k + b 이고 bytes의 크기는 k+1, 즉 8(k+1) 비트이므로
		// @@@ io.ReadFull 로 8(k+1) 비트가 전부 가득차 있는 상태에는 원하는 bits 수 보다 8 - b bit 더 있음
		// @@@ bytes[0]은 k+1개의 바이트 중 최상위 바이트
		// @@@ 최상위 바이트 8 비트 중 제일 앞 8 - b bit 초과 비트 부분을 0으로 마스킹 해주면 bytes에 정확히 bits 수 만큼 데이터가 들어가 있게 된다.
		// @@@ ==> 8 비트 중 앞 8 - b 비트는 0 이고 나머지 b 비트는 1인 마스크를 uint8(int(1<<b) - 1)로 생성하고
		// @@@ ==> 최상위 바이트 bytes[0]에 &(and) 연산으로 마스킹 진행
		// @@@ @@@ b가 0인 경우 (bits = 8k) bytes의 크기는 k 이므로 bytes에 bits 수보다 큰 초과 비트가 없다
		// @@@ @@@ ==> 최상위 바이트의 전부분을 살려야 하므로 마스크는 11111111이어야 한다 ==> b가 0 일 때 b = 8로 변경해야 마스크가 11111111이 된다

		// 소수 후보의 상위 비트 조작: 상위 2개의 비트를 1로 만듦 (충분히 큰 수가 되도록 제일 앞 2자리가 무조건 1이 되도록 강제)
		if b >= 2 {
			// if: 최상위 바이트(bytes[0])에 유효 비트 수(b)가 2 이상인 경우
			bytes[0] |= 3 << (b - 2)
			// @@@ 유효 비트 제일 앞 두자리에 11로 |(or) 마스킹 연산이 되도록 3 << (b - 2)로 마스크 생성
			// @@@ @@@ 3은 2진법 표현 시 11이고 마스크는 11 뒤에 0이 (b-2)개 붙어 있어야 함
			// @@@ @@@ ===> 유효 비트 제일 앞 2자리는 11과 or 연산이 되어 무조건 11 보장
		} else {
			// else: 최상위 바이트(bytes[0])에 유효 비트 수(b)가 1인 경우
			bytes[0] |= 1
			// @@@ bytes[0]의 2^0 자리만 유효 비트 ==> 이 2^0자리가 무조건 1이 되도록 0b1과 or 연산
			if len(bytes) > 1 {
				// if: 차상위 바이트 (bytes[1])이 있는지 확인
				bytes[1] |= 0x80
				// @@@ bytes[1]가 있으면 bytes[1]의 최상위 비트(2^7 자리)가 무조건 1이 되도록 0b10000000 = 0x80과 or 연산
			}
		}

		// 홀수 보장: 최하위 비트(맨 마지막 바이트)를 1로 만들어 반드시 홀수(짝수면 2를 제외하고 소수 불가)
		bytes[len(bytes)-1] |= 1
		// @@@ 최하위 바이트(bytes[len(bytes)-1])의 최하위 비트(2^0자리)에 0b1 마스크를 or 연산하면 최하위 비트가 무조건 1

		// big.Int로 변환: bytes 슬라이스를 big.Int로 변환
		p.SetBytes(bytes)

		// 수 판별: 해당 수가 소수인지 밀러-라빈(Miller-Rabin) 테스트 20회 실시
		if p.ProbablyPrime(20) {
			// if: 20회 확률적 소수 판정(밀러-라빈) 테스트 통과 시(true) 소수로 인정, 반환
			return p, nil
		}
		// else: 테스트 실패 시(false) 루프 처음으로 돌아가 다시 숫자 생성 반복
	}
}
