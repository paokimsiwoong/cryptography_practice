package ch1l7

import (
	"fmt"
	"testing"
)

func TestKeyToCipher(t *testing.T) {
	type testCase struct {
		key        string
		shouldFail bool
	}

	runCases := []testCase{
		{"thisIsMySecretKeyIHopeNoOneFinds", false}, // Valid key
		{"short", true}, // Too short key
		{"an extremely long key that exceeds the block size", true}, // Too long key
	}

	submitCases := append(runCases, []testCase{
		{"thisIsA32ByteKeyForAES256Testing!", true}, // Valid 32-byte key for AES-256
		{"valid16ByteKeyHere", true},                // Valid 16-byte key for AES-128
		{"invalid-key", true},                       // Invalid key, not the correct length
		{"ThisIsA24ByteKeyForAES192Testing", false}, // Valid key for AES-192
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		_, err := keyToCipher(test.key)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      key: %v
Expecting:   Error: %v
Actual:      Error: %v
Fail
`, test.key, test.shouldFail, err != nil)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      key: %v
Expecting:   Error: %v
Actual:      Error: %v
Pass
`, test.key, test.shouldFail, err != nil)
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
