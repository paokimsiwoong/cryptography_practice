package ch1l11

import (
	"fmt"
	"testing"
)

func TestGenerateRandomKey(t *testing.T) {
	type testCase struct {
		length     int
		shouldFail bool
		expected   string
	}

	runCases := []testCase{
		{16, false, "0194fdc2fa2ffcc041d3ff12045b73c8"},                                 // Expected output for 16 bytes
		{32, false, "0194fdc2fa2ffcc041d3ff12045b73c86e4ff95ff662a5eee82abdf44a2d0b75"}, // Expected output for 32 bytes
	}

	submitCases := append(runCases, []testCase{
		{8, false, "0194fdc2fa2ffcc0"}, // Expected output for 8 bytes
		{64, false, "0194fdc2fa2ffcc041d3ff12045b73c86e4ff95ff662a5eee82abdf44a2d0b75fb180daf48a79ee0b10d394651850fd4a178892ee285ece1511455780875d64e"}, // Expected output for 64 bytes
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		key, err := generateRandomKey(test.length)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      length: %v
Expecting:   Error: %v
Actual:      Error: %v
Fail
`, test.length, test.shouldFail, err != nil)
		} else {
			if !test.shouldFail && key != test.expected {
				failCount++
				t.Errorf(`---------------------------------
Inputs:      length: %v
Expecting:   Key: %v
Actual:      Key: %v
Fail
`, test.length, test.expected, key)
			} else if test.shouldFail {
				passCount++
				fmt.Printf(`---------------------------------
Inputs:      length: %v
Expecting:   Error: %v
Actual:      Error: %v
Pass
`, test.length, test.shouldFail, err != nil)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Inputs:      length: %v
Expecting:   Key: %v
Actual:      Key: %v
Pass
`, test.length, test.expected, key)
			}
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
