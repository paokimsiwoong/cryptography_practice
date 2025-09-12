package ch3l1

import (
	"fmt"
	"testing"
)

func TestAlphabetSize(t *testing.T) {
	type testCase struct {
		numBits  int
		expected float64
	}

	runCases := []testCase{
		{1, 2},      // 2^1 = 2
		{2, 4},      // 2^2 = 4
		{3, 8},      // 2^3 = 8
		{4, 16},     // 2^4 = 16
		{5, 32},     // 2^5 = 32
		{8, 256},    // 2^8 = 256
		{16, 65536}, // 2^16 = 65536
	}

	submitCases := append(runCases, []testCase{
		{6, 64},  // 2^6 = 64
		{7, 128}, // 2^7 = 128
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result := alphabetSize(test.numBits)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      numBits: %v
Expecting:   %v
Actual:      %v
Fail
`, test.numBits, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      numBits: %v
Expecting:   %v
Actual:      %v
Pass
`, test.numBits, test.expected, result)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
