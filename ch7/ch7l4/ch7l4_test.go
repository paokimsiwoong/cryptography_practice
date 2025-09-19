package ch7l4

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPadWithZeros(t *testing.T) {
	type testCase struct {
		input       []byte
		desiredSize int
		expected    []byte
	}

	runCases := []testCase{
		{[]byte{0xFF}, 4, []byte{0xFF, 0x00, 0x00, 0x00}},
		{[]byte{0xFA, 0xBC}, 8, []byte{0xFA, 0xBC, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{0x12, 0x34, 0x56}, 12, []byte{0x12, 0x34, 0x56, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}

	submitCases := append(runCases, []testCase{
		{[]byte{0xFA}, 16, []byte{0xFA, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{[]byte{}, 10, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passed, failed := 0, 0

	for _, test := range testCases {
		result := padWithZeros(test.input, test.desiredSize)
		if !bytes.Equal(result, test.expected) {
			t.Errorf(`---------------------------------
Input:     %v
Expecting: %v
Actual:    %v
Fail
`, test.input, test.expected, result)
			failed++
		} else {
			fmt.Printf(`---------------------------------
Input:     %v
Expecting: %v
Actual:    %v
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
