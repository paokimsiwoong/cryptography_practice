package ch12l1

import (
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	type testCase struct {
		passwords []string
		expected  string
	}

	runCases := []testCase{
		{[]string{"password1", "password2", "password3"}, "2ccb27b6da"},
		{[]string{"abercromni3", "f1tch", "123456", "abcdefg1234"}, "a03ea2f828"},
		{[]string{"IHeartNanciedrake", "m7B1rthd@y"}, "7bf580ff47"},
	}

	submitCases := append(runCases, []testCase{
		{[]string{"morepassw0rds", "evenmorepassw0rds"}, "2f595e539a"},
		{[]string{"s3cur3passw0rd"}, "db25c1918f"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed, failed := 0, 0

	for _, test := range testCases {
		h := newHasher()

		for _, password := range test.passwords {
			h.Write(password)
		}

		actual := h.GetHex()

		if actual[:10] != test.expected {
			failed++
			t.Errorf(`---------------------------------
Hashing vault with passwords: %v
Expecting:   hash starts with: %s
Actual:      hash starts with: %s
Fail
`, test.passwords, test.expected, actual[:10])
		} else {
			passed++
			fmt.Printf(`---------------------------------
Hashing vault with passwords: %v
Expecting:   hash starts with: %s
Actual:      hash starts with: %s
Pass
`, test.passwords, test.expected, actual[:10])
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
