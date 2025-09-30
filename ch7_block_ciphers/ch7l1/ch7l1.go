package ch7l1

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	// ?

	var block cipher.Block
	var err error

	switch cipherType {
	case typeAES:
		block, err = aes.NewCipher(make([]byte, keyLen))
		// aes block의 길이는 16바이트 고정
		if err != nil {
			return 0, err
		}
	case typeDES:
		block, err = des.NewCipher(make([]byte, keyLen))
		// des block의 길이는 8바이트 고정
		if err != nil {
			return 0, err
		}
	default:
		return 0, errors.New("invalid cipher type")
	}

	return block.BlockSize(), nil
}
