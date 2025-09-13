package main

import (
	"fmt"
	"log"

	"github.com/paokimsiwoong/cryptography_practice/ch3/ch3l5"
)

func main() {
	key, err := ch3l5.FindKey([]byte{0x1b, 0x2c, 0x3d}, "yes")
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("Found key: %v\n", key)
}
