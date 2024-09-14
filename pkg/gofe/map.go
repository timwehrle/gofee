package gofe

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Charset is a global variable that holds the set of characters used to generate passwords.
// It is built based on the provided PasswordConfig.
var Charset string

// MapToCharset generates a random password of length 'n' using the configured Charset.
// It returns the generated password or an error if Charset is empty or random number generation fails.
func MapToCharset(n int, config PasswordConfig) (string, error) {
	// Build the Charset based on the provided configuration.
	Charset = BuildCharset(config)
	charsetLen := int64(len(Charset))

	// Return an error if no characters are available in the Charset.
	if charsetLen == 0 {
		return "", fmt.Errorf("charset is empty")
	}

	// Allocate space for the generated password.
	ret := make([]byte, n)

	// Generate 'n' random characters from the Charset.
	for i := 0; i < n; i++ {
		// Generate a random number in the range [0, charsetLen).
		num, err := rand.Int(rand.Reader, big.NewInt(charsetLen))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}
		// Assign the corresponding character to the password.
		ret[i] = Charset[num.Int64()]
	}

	// Convert the byte slice to a string and return the generated password.
	return string(ret), nil
}