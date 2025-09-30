package ch8l8

import (
	"fmt"
	"testing"
)

func TestGenerateIV(t *testing.T) {
	type testCase struct {
		length   int
		expected int
	}

	runCases := []testCase{
		{8, 8},
		{10, 10},
		{16, 16},
	}

	submitCases := append(runCases, []testCase{
		{12, 12},
		{14, 14},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		iv, err := generateIV(test.length)
		if err != nil {
			t.Errorf("Failed to generate IV for length %d: %v", test.length, err)
			failed++
			continue
		}

		if len(iv) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      length: %d
Expecting:   IV length: %d
Actual:      IV length: %d
Fail
`, test.length, test.expected, len(iv))
			failed++
		} else {
			fmt.Printf(`---------------------------------
Inputs:      length: %d
Expecting:   IV length: %d
Actual:      IV length: %d
Pass
`, test.length, test.expected, len(iv))
			passed++
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
