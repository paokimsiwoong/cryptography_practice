package ch10l1

import (
	"fmt"
	"testing"
)

func TestGenKeys(t *testing.T) {
	type testCase struct {
		expected bool
	}

	runCases := []testCase{
		{true},
	}

	submitCases := append(runCases, []testCase{
		{true},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		pubKey, privKey, err := genKeys()
		if err != nil {
			t.Errorf("Failed to generate key pair, error: %v", err)
			failed++
			continue
		}

		arePaired := keysArePaired(pubKey, privKey)
		if arePaired != test.expected {
			t.Errorf(`---------------------------------
Inputs:      public key: %v, private key: %v
Expecting:   paired: %v
Actual:      paired: %v
Fail
`, pubKey, privKey, test.expected, arePaired)
			failed++
		} else {
			fmt.Printf(`---------------------------------
Inputs:      public key: %v, private key: %v
Expecting:   paired: %v
Actual:      paired: %v
Pass
`, pubKey, privKey, test.expected, arePaired)
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
