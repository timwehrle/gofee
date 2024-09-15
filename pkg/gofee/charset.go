package gofee

import "strings"

// Charset constants for lowercase letters, uppercase letters, digits, and symbols.
const (
	Lowers  = "abcdefghijklmnopqrstuvwxyz"
	Uppers  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits  = "0123456789"
	Symbols = "!@#$%^&*()-_=+[]{}|;:,.<>?/~"
	All     = Lowers + Uppers + Digits + Symbols
)

type PasswordConfig struct {
	IncludeLowers  bool
	IncludeUppers  bool
	IncludeDigits  bool
	IncludeSymbols bool
}

func BuildCharset(config PasswordConfig) string {
	if config.IncludeLowers && config.IncludeUppers && config.IncludeDigits && config.IncludeSymbols {
		return All
	}

	var builder strings.Builder

	if config.IncludeLowers {
		builder.WriteString(Lowers)
	}

	if config.IncludeUppers {
		builder.WriteString(Uppers)
	}

	if config.IncludeDigits {
		builder.WriteString(Digits)
	}

	if config.IncludeSymbols {
		builder.WriteString(Symbols)
	}

	return builder.String()
}
