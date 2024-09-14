package gofe

import "math"

func CalculateEntropy(charsetSize, passwordLength int) float64 {
	return float64(passwordLength) * math.Log2(float64(charsetSize))
}
