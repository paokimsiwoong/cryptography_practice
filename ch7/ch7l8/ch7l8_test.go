package ch7l8

import (
	"fmt"
	"testing"
)

func TestDeriveRoundKey(t *testing.T) {
	type testCase struct {
		masterKey   [4]byte
		roundNumber int
		expected    [4]byte
	}

	runCases := []testCase{
		{[4]byte{0xAA, 0xFF, 0x11, 0xBC}, 1, [4]byte{0xAB, 0xFE, 0x10, 0xBD}},
		{[4]byte{0xEB, 0xCD, 0x13, 0xFC}, 2, [4]byte{0xE9, 0xCF, 0x11, 0xFE}},
	}

	submitCases := append(runCases, []testCase{
		{[4]byte{0xAA, 0xFF, 0x11, 0xBC}, 5, [4]byte{0xAF, 0xFA, 0x14, 0xB9}},
		{[4]byte{0xEB, 0xCD, 0x13, 0xFC}, 7, [4]byte{0xEC, 0xCA, 0x14, 0xFB}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		result := deriveRoundKey(test.masterKey, test.roundNumber)
		if result != test.expected {
			failed++
			t.Errorf(`---------------------------------
Inputs:    masterKey: %X, roundNumber: %d
Expecting: roundKey: %X
Actual:    roundKey: %X
Fail
`, test.masterKey, test.roundNumber, test.expected, result)
		} else {
			passed++
			fmt.Printf(`---------------------------------
Inputs:    masterKey: %X, roundNumber: %d
Expecting: roundKey: %X
Actual:    roundKey: %X
Pass
`, test.masterKey, test.roundNumber, test.expected, result)
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
