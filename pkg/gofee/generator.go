package gofee

import (
	"fmt"
)

// Generate creates a random password of the specified length using the given PasswordConfig.
// It returns the generated password or an error if the length is invalid or password generation fails.
func Generate(length int, config PasswordConfig) (string, error) {
	// Check if the provided length is valid (i.e., greater than 0).
	if length <= 0 {
		// Return an error if the length is not valid.
		return "", fmt.Errorf("length must be greater than 0")
	}

	// Call MapToCharset to generate a password based on the length and configuration.
	// This function generates a password by mapping random numbers to characters from the charset.
	password, err := MapToCharset(length, config)
	if err != nil {
		// Return an error if MapToCharset fails, including the specific error message.
		return "", fmt.Errorf("error mapping number to charset: %v", err)
	}

	// Return the successfully generated password.
	return password, nil
}
