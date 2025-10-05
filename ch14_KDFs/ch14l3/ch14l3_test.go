package ch14l3

import (
	"fmt"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	type testCase struct {
		password1 string
		password2 string
		expected  bool
	}

	runCases := []testCase{
		{"thisIsAPassword", "thisIsAPassword", true},
		{"thisIsAPassword", "thisIsAnotherPassword", false},
		{"corr3ct h0rse", "corr3ct h0rse", true},
	}

	submitCases := append(runCases, []testCase{
		{"thisIsAPassword", "thisIsAPassword", true},
		{"thisIsAPassword", "thisIsAnotherPassword", false},
		{"corr3ct h0rse", "corr3ct h0rse", true},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed := 0
	failed := 0

	for _, test := range testCases {
		hashed, err := hashPassword(test.password1)
		if err != nil {
			t.Errorf(`---------------------------------
hashing password: %
`, err)
			failed++
			continue
		}

		match := checkPasswordHash(test.password2, hashed)

		if match != test.expected {
			t.Errorf(`---------------------------------
Password 1:  %s
Password 2:  %s
Expecting:   %v
Actual:      %v
Fail
`, test.password1, test.password2, test.expected, match)
			failed++
		} else {
			fmt.Printf(`---------------------------------
Password 1:  %s
Password 2:  %s
Expecting:   %v
Actual:      %v
Pass
`, test.password1, test.password2, test.expected, match)
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
