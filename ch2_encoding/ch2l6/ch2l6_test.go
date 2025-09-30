package ch2l6

import (
	"fmt"
	"testing"
)

func TestGetHexBytes(t *testing.T) {
	type testCase struct {
		input      string
		expected   []byte
		shouldFail bool
	}

	runCases := []testCase{
		{"48:65:6c:6c:6f", []byte("Hello"), false},             // Hex for "Hello"
		{"57:6f:72:6c:64", []byte("World"), false},             // Hex for "World"
		{"50:61:73:73:77:6f:72:64", []byte("Password"), false}, // Hex for "Password"
		{"ZZ:YY:XX", nil, true},                                // Invalid hex
	}

	submitCases := append(runCases, []testCase{
		{"48:65:6c:6c:6f", []byte("Hello"), false},             // Hex for "Hello"
		{"57:6f:72:6c:64", []byte("World"), false},             // Hex for "World"
		{"50:61:73:73:77:6f:72:64", []byte("Password"), false}, // Hex for "Password"
		{"ZZ:YY:XX", nil, true},                                // Invalid hex
		{"4c:65:61:72:6e:69:6e:67", []byte("Learning"), false}, // Hex for "Learning"
		{"54:65:73:74", []byte("Test"), false},                 // Hex for "Test"
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result, err := getHexBytes(test.input)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      %v
Expecting:   Error: %v
Actual:      Error: %v
Fail
`, test.input, test.shouldFail, err != nil)
		} else if !test.shouldFail && string(result) != string(test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      %v
Expecting:   %v
Actual:      %v
Fail
`, test.input, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      %v
Expecting:   %v
Actual:      %v
Pass
`, test.input, test.expected, result)
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
