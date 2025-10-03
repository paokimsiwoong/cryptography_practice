package ch12l10

import (
	"fmt"
	"testing"
)

func TestHashFunction(t *testing.T) {
	type testCase struct {
		input    string
		expected [4]byte
	}

	runCases := []testCase{
		{"Example message", [4]byte{0x24, 0xC0, 0x40, 0xC4}},
		{"This is a slightly longer example to hash", [4]byte{0x28, 0x88, 0xC8, 0x48}},
		{"This is a much longer example of some text to hash, maybe it's the opening paragraph of a blog post", [4]byte{0xA8, 0x04, 0x44, 0xE8}},
	}

	submitCases := append(runCases, []testCase{
		{"A very secret password", [4]byte{0x64, 0xA0, 0x64, 0xAC}},
		{"Another very secret password", [4]byte{0x40, 0xCC, 0xEC, 0x2C}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		actual := hash([]byte(test.input))

		if actual != test.expected {
			failed++
			fmt.Printf(`---------------------------------
Hashing:	'%s'
Expecting:	%X
Actual:		%X
Fail
`, test.input, test.expected, actual)
		} else {
			passed++
			fmt.Printf(`---------------------------------
Hashing:	'%s'
Expecting:	%X
Actual:		%X
Pass
`, test.input, test.expected, actual)
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
