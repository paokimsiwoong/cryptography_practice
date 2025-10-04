package ch13l7

import (
	"fmt"
	"testing"
)

func TestHmac(t *testing.T) {
	type testCase struct {
		message  string
		key      string
		expected string
	}

	runCases := []testCase{
		{"I hope no one finds the Bitcoin keys I keep under my mailbox", "super_secret_password", "70a31170565d94a0616d4323928212be7068ff783bef9e627564d2f726c6b8d1"},
		{"No really, they're just written on a piece of paper", "correct horse battery staple", "75048eb8a6aa4d70fe9a9e1f720c635da325b69de9bdab7dd9f7b3120632473a"},
		{"It's like a gazillion satoshis worth of BTC", "aFiveDoll@rWr3nch", "e80961791badc022608464fb871a2b1ac68c6a1b97b5b1eb1cfa880d0f121912"},
	}

	submitCases := append(runCases, []testCase{
		{"This is a submit case message", "submitKeyTest", "a1058359982ecab5c35e822c74dff54391d01d15064a921b4be32bb4552513f2"},
		{"Another submit case", "submitKeyTestTwo", "6419bb7eaa75ab4a82aa8ce076a90dde0d3602ba08b893f2a4e5c2ed6db29e0b"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passed := 0
	failed := 0

	for _, test := range testCases {
		h := hmac(test.message, test.key)

		if h != test.expected {
			failed++
			t.Errorf(`---------------------------------
Message:     %s
Key:         %s
Expecting:   %s
Actual:      %s
Fail
`, test.message, test.key, test.expected, h)
		} else {
			passed++
			fmt.Printf(`---------------------------------
Message:     %s
Key:         %s
Expecting:   %s
Actual:      %s
Pass
`, test.message, test.key, test.expected, h)
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
