package gofee

import (
	"fmt"
	"math"
)

// CalculateEntropy returns the entropy (in bits) of a password,
// given the character set size and password length.
func CalculateEntropy(charsetSize, passwordLength int) (float64, error) {
	// Return an error if the charset size or password length is invalid.
	if charsetSize <= 0 {
		return 0, fmt.Errorf("charset size must be greater than 0")
	}
	if passwordLength <= 0 {
		return 0, fmt.Errorf("password length must be greater than 0")
	}

	// Calculate entropy using the formula: entropy = passwordLength * log2(charsetSize)
	return float64(passwordLength) * math.Log2(float64(charsetSize)), nil
}
