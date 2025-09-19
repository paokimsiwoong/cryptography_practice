package ch7l1

import (
	"fmt"
)

const (
	typeAES = iota
	typeDES
)

func getCipherTypeName(cipherType int) string {
	switch cipherType {
	case typeAES:
		return "AES"
	case typeDES:
		return "DES"
	}
	return "unknown"
}

func test(keyLen, cipherType int) {
	fmt.Printf(
		"Getting block size of %v cipher with key length %v...\n",
		getCipherTypeName(cipherType),
		keyLen,
	)
	blockSize, err := getBlockSize(keyLen, cipherType)
	if err != nil {
		fmt.Println(err)
		fmt.Println("========")
		return
	}
	fmt.Println("Block size:", blockSize)
	fmt.Println("========")
}
