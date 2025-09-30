package ch9l1

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	type testCase struct {
		key       []byte
		plaintext []byte
		nonce     []byte
		expected  string
	}

	runCases := []testCase{
		{[]byte("d00c5215-60f6-4ac4-9648-532b5dad"), []byte("I wonder what he's thinking about me??"), generateNonce(12), "I wonder what he's thinking about me??"},
		{[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), []byte("I knew it, Becky has been cheating this whole time!"), generateNonce(12), "I knew it, Becky has been cheating this whole time!"},
	}

	submitCases := append(runCases, []testCase{
		{[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), []byte("AES in GCM mode is strong!"), generateNonce(12), "AES in GCM mode is strong!"},
		{[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), []byte("Testing AES-GCM encryption."), generateNonce(12), "Testing AES-GCM encryption."},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	var passed, failed int

	for _, test := range testCases {
		ciphertext, err := encrypt(test.key, test.plaintext, test.nonce)
		if err != nil {
			t.Errorf("Encryption failed for key: %v, plaintext: %v, error: %v", string(test.key), string(test.plaintext), err)
			failed++
			continue
		}

		decryptedText, err := decrypt(test.key, ciphertext, test.nonce)
		if err != nil {
			t.Errorf("Decryption failed for key: %v, ciphertext: %v, error: %v", string(test.key), ciphertext, err)
			failed++
			continue
		}

		if string(decryptedText) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      key: %v, plaintext: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Fail
`, string(test.key), string(test.plaintext), test.expected, string(decryptedText))
			failed++
		} else {
			fmt.Printf(`---------------------------------
Inputs:      key: %v, plaintext: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Pass
`, string(test.key), string(test.plaintext), test.expected, string(decryptedText))
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
