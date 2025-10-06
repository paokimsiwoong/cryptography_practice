package ch14l6

import (
	"fmt"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	type testCase struct {
		password1 string
		password2 string
		saltLen   int
		expect    bool
	}

	runCases := []testCase{
		{"samepass", "samepass", 16, true},
		{"passone", "passtwo", 24, false},
		{"correct horse battery staple", "correct horse battery staple", 32, true},
	}

	submitCases := append(runCases, []testCase{
		{"bigtimepass", "notthesame", 16, false},
		{"kaladin", "kaladin", 24, true},
		{"stormlight archive", "stormlight archive", 32, true},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		salt, err := generateSalt(test.saltLen)
		if err != nil {
			t.Errorf(`---------------------------------
Error generating salt: %v
`, err)
			continue
		}

		hashed1 := hashPassword([]byte(test.password1), salt)
		hashed2 := hashPassword([]byte(test.password2), salt)

		match := string(hashed1) == string(hashed2)
		if match != test.expect {
			failed++
			t.Errorf(`---------------------------------
Password 1:  %s
Password 2:  %s
Salt length: %d
Expecting:   %v
Actual:      %v
Fail
`, test.password1, test.password2, test.saltLen, test.expect, match)
		} else {
			passed++
			fmt.Printf(`---------------------------------
Password 1:  %s
Password 2:  %s
Salt length: %d
Expecting:   %v
Actual:      %v
Pass
`, test.password1, test.password2, test.saltLen, test.expect, match)
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
