package ch5l1

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	type testCase struct {
		lhs      bool
		rhs      bool
		expected bool
	}

	runCases := []testCase{
		{false, true, true},   // false XOR true = true
		{false, false, false}, // false XOR false = false
	}

	submitCases := append(runCases, []testCase{
		{true, true, false}, // true XOR true = false
		{true, false, true}, // true XOR false = true
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result := xor(test.lhs, test.rhs)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      lhs: %v, rhs: %v
Expecting:   %v
Actual:      %v
Fail
`, test.lhs, test.rhs, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      lhs: %v, rhs: %v
Expecting:   %v
Actual:      %v
Pass
`, test.lhs, test.rhs, test.expected, result)
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
