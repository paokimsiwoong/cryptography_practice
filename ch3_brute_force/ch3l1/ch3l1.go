package ch3l1

func alphabetSize(numBits int) float64 {
	// ?
	result := 1 << numBits
	return float64(result)
	// return math.Exp2(float64(numBits))
	// return math.Pow(2, float64(numBits)) 도 가능
}
