package gofe

import (
	"fmt"
)

func Generate(length int, config PasswordConfig) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	password, err := MapToCharset(length, config)
	if err != nil {
		return "", fmt.Errorf("error mapping number to charset: %v", err)
	}

	return password, nil
}
