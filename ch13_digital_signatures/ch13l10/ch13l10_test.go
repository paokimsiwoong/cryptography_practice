package ch13l10

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestECDSAMessage(t *testing.T) {
	type testCase struct {
		message  string
		expected string
	}

	runCases := []testCase{
		{"userid:2f9c584e-5d25-4516-a0ed-ddfa6e152006", "valid"},
		{"userid:0e803af6-292f-4432-a285-84a7591e25a8", "valid"},
		{"userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f", "valid"},
	}

	submitCases := append(runCases, []testCase{
		{"userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f", "valid"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passed := 0
	failed := 0

	for _, test := range testCases {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf(`---------------------------------
Error generating key: %v
`, err)
			failed++
			continue
		}

		token, err := createECDSAMessage(test.message, privateKey)
		if err != nil {
			t.Errorf(`---------------------------------
Error creating ECDSA message: %v
`, err)
			failed++
			continue
		}

		err = verifyECDSAMessage(token, &privateKey.PublicKey)

		if err != nil {
			t.Errorf(`---------------------------------
Message:      %s
Expecting:    valid
Actual:       invalid
Fail
`, test.message)
			failed++
		} else {
			fmt.Printf(`---------------------------------
Message:      %s
Expecting:    valid
Actual:       valid
Pass
`, test.message)
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
