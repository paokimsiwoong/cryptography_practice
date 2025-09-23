package ch8l11

import (
	"fmt"
	"testing"
)

func TestSBox(t *testing.T) {
	type testCase struct {
		input    byte
		expected byte
	}

	runCases := []testCase{
		{0b0000, 0b00},
		{0b0001, 0b10},
		{0b0110, 0b11},
		{0b1111, 0b00},
	}

	submitCases := append(runCases, []testCase{
		{0b1001, 0b11},
		{0b0010, 0b01},
		{0b110010, 0b01},
		{0b01111001, 0b11},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passed := 0
	failed := 0

	for _, test := range testCases {
		result, err := sBox(test.input)
		if err != nil {
			t.Errorf("Failed for input: %04b, error: %v", test.input, err)
			failed++
			continue
		}

		if result != test.expected {
			t.Errorf(`---------------------------------
Inputs:      input: %04b
Expecting:   output: %02b
Actual:      output: %02b
Fail
`, test.input, test.expected, result)
			failed++
		} else {
			fmt.Printf(`---------------------------------
Inputs:      input: %04b
Expecting:   output: %02b
Actual:      output: %02b
Pass
`, test.input, test.expected, result)
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
