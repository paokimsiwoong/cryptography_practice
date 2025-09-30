package ch7l1

import (
	"fmt"
	"testing"
)

func TestGetBlockSize(t *testing.T) {
	type testCase struct {
		keyLen     int
		cipherType int
		expected   int
		shouldFail bool
	}

	runCases := []testCase{
		{64, typeAES, 0, true}, // Invalid AES key length
		{8, typeDES, 8, false}, // Valid DES key length
		{16, typeDES, 0, true}, // Invalid DES key length
		{1, -1, 0, true},       // Invalid cipher type
	}

	submitCases := append(runCases, []testCase{
		{16, typeAES, 16, false}, // Valid AES key length
		{24, typeAES, 16, false}, // Valid AES key length
		{32, typeAES, 16, false}, // Valid AES key length
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		blockSize, err := getBlockSize(test.keyLen, test.cipherType)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      keyLen: %v, cipherType: %v
Expecting:   Error: %v
Actual:      Error: %v
Fail
`, test.keyLen, test.cipherType, test.shouldFail, err != nil)
		} else if blockSize != test.expected && !test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      keyLen: %v, cipherType: %v
Expecting:   Block Size: %v
Actual:      Block Size: %v
Fail
`, test.keyLen, test.cipherType, test.expected, blockSize)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      keyLen: %v, cipherType: %v
Expecting:   Block Size: %v
Actual:      Block Size: %v
Pass
`, test.keyLen, test.cipherType, test.expected, blockSize)
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
