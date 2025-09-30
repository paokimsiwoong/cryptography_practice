package ch11l5

import (
	"fmt"
	"testing"
)

func TestGeneratePrivateNumsAndGetN(t *testing.T) {
	type testCase struct {
		keySize  int
		expected int // Number of digits in n
	}

	runCases := []testCase{
		{512, 309},  // Expected number of digits for 512-bit primes
		{1024, 617}, // Expected number of digits for 1024-bit primes
	}

	submitCases := append(runCases, []testCase{
		{2048, 1233}, // Expected number of digits for 2048-bit primes
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	var passed, failed int

	for _, test := range testCases {
		p, q := generatePrivateNums(test.keySize)
		n := getN(p, q)

		firstP := firstNDigits(*p, 10)
		firstQ := firstNDigits(*q, 10)
		firstN := firstNDigits(*n, 10)

		if len(n.String()) != test.expected {
			failed++
			t.Errorf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
n: %s
Expecting:   n digits: %d
Actual:      n digits: %d
Fail
`, test.keySize, firstP, firstQ, firstN, test.expected, len(n.String()))
		} else {
			passed++
			fmt.Printf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
n: %s
Expecting:   n digits: %d
Actual:      n digits: %d
Pass
`, test.keySize, firstP, firstQ, firstN, test.expected, len(n.String()))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passed, failed)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
