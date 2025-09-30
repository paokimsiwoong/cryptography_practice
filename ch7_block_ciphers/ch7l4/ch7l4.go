package ch7l4

func padWithZeros(block []byte, desiredSize int) []byte {
	// ?

	sizeDif := desiredSize - len(block)

	block = append(block, make([]byte, sizeDif)...)

	return block[:desiredSize]
}
