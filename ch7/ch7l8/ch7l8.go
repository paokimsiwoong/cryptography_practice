package ch7l8

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	// ?

	roundKey := [4]byte{}

	for i, m := range masterKey {
		roundKey[i] = m ^ byte(roundNumber)
	}

	return roundKey
}
