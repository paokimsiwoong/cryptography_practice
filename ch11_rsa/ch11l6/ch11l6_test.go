package ch11l6

import (
	"fmt"
	"testing"
)

func TestGetTotAndGetE(t *testing.T) {
	type testCase struct {
		keySize  int
		expected int
	}

	runCases := []testCase{
		{512, 309},  // Expected number of digits for tot with 512-bit primes
		{1024, 617}, // Expected number of digits for tot with 1024-bit primes
	}

	submitCases := append(runCases, []testCase{
		{2048, 1233}, // Expected number of digits for tot with 2048-bit primes
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		p, q := generatePrivateNums(test.keySize)
		tot := getTot(p, q)
		e := getE(tot)

		firstP := firstNDigits(*p, 10)
		firstQ := firstNDigits(*q, 10)
		firstTot := firstNDigits(*tot, 10)
		firstE := firstNDigits(*e, 10)

		if len(tot.String()) != test.expected {
			failed++
			t.Errorf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
ϕ(n): %s
Expecting:   tot digits: %d
Actual:      tot digits: %d
Fail
`, test.keySize, firstP, firstQ, firstTot, test.expected, len(tot.String()))
		} else {
			passed++
			fmt.Printf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
ϕ(n): %s
Expecting:   tot digits: %d
Actual:      tot digits: %d
Pass
`, test.keySize, firstP, firstQ, firstTot, test.expected, len(tot.String()))
		}

		fmt.Printf("Generated e: %s\n", firstE)
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
