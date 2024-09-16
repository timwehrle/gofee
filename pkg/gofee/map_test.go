package gofee

import (
	"crypto/rand"
	"errors"
	"testing"
)

// TestMapToCharset runs a series of tests for the MapToCharset function.
// It covers various scenarios such as invalid input, empty charset, and valid configurations.
func TestMapToCharset(t *testing.T) {
	// Define the test arguments and expected results for each test case.
	type args struct {
		length int
		config PasswordConfig
	}

	// List of test cases to run, covering different configurations and expectations.
	tests := []struct {
		name        string // Test name for identification.
		args        args   // Test inputs including password length and configuration.
		wantErr     bool   // Whether an error is expected.
		expectedLen int    // The expected length of the generated password.
		expectedSet string // The expected set of characters used in the password.
	}{
		// Test case where the length is invalid (0).
		{
			name: "Invalid length",
			args: args{
				length: 0,
				config: PasswordConfig{
					IncludeLowers:  true,
					IncludeUppers:  true,
					IncludeDigits:  true,
					IncludeSymbols: true,
				},
			},
			wantErr:     true,
			expectedLen: 0,
			expectedSet: "",
		},
		// Test case where the charset is empty due to all false flags in the config.
		{
			name: "Empty charset",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  false,
					IncludeUppers:  false,
					IncludeDigits:  false,
					IncludeSymbols: false,
				},
			},
			wantErr:     true,
			expectedLen: 0,
			expectedSet: "",
		},
		// Test case with all character types included and a valid length.
		{
			name: "Valid charset",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  true,
					IncludeUppers:  true,
					IncludeDigits:  true,
					IncludeSymbols: true,
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: All, // Expected set is the entire charset.
		},
		// Test case without symbols, only lowercase, uppercase, and digits.
		{
			name: "Valid charset with no symbols",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  true,
					IncludeUppers:  true,
					IncludeDigits:  true,
					IncludeSymbols: false,
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: Lowers + Uppers + Digits, // Expected charset excludes symbols.
		},
		// Test case without lowercase characters, includes uppercase, digits, and symbols.
		{
			name: "Valid charset with no lowercase",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  false,
					IncludeUppers:  true,
					IncludeDigits:  true,
					IncludeSymbols: true,
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: Uppers + Digits + Symbols, // Expected charset excludes lowercase.
		},
		{
			name: "Valid charset with type pin",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  false,
					IncludeUppers:  false,
					IncludeDigits: false,
					IncludeSymbols: false,
					Type: "pin",
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: Digits, // Expected charset is only digits.
		},
		{
			name: "Valid charset with type memorable",
			args: args{
				length: 16,
				config: PasswordConfig{
					IncludeLowers:  false,
					IncludeUppers:  false,
					IncludeDigits: false,
					IncludeSymbols: false,
					Type: "memorable",
				},
			},
			wantErr:     false,
			expectedLen: 16,
			expectedSet: Lowers + Uppers, // Expected charset includes lowercase and uppercase.
		},
	}

	// Iterate through each test case and run the test.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the MapToCharset function with the test arguments.
			got, err := MapToCharset(tt.args.length, tt.args.config)

			// Check if the error status matches the expected error status.
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToCharset() error = %v, wanted error = %v", err, tt.wantErr)
			}

			// Check if the length of the generated password matches the expected length.
			if len(got) != tt.expectedLen {
				t.Errorf("MapToCharset() length = %v, wanted length = %v", len(got), tt.expectedLen)
			}

			// Verify that all characters in the generated password are from the expected set.
			for _, c := range got {
				if !Contains(tt.expectedSet, c) {
					t.Errorf("MapToCharset() character = %v, not found in expected set", c)
				}
			}
		})
	}
}

// errReader is a mock reader that always returns an error, used to simulate failure in rand.Reader.
type errReader struct{}

// Read method of errReader always returns an error, simulating failure during random number generation.
func (e *errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mocked error from rand.Reader")
}

// TestMapToCharset_RandIntError tests the MapToCharset function when rand.Reader fails to generate random numbers.
// It replaces rand.Reader with errReader to simulate the error and checks if the error is handled correctly.
func TestMapToCharset_RandIntError(t *testing.T) {
	// Save the original rand.Reader to restore later.
	originalReader := rand.Reader

	// Replace rand.Reader with the errReader to simulate a random number generation failure.
	rand.Reader = &errReader{}

	// Ensure that rand.Reader is restored to its original value after the test.
	defer func() {
		rand.Reader = originalReader
	}()

	// Test case configuration to generate a password.
	config := PasswordConfig{
		IncludeLowers:  true,
		IncludeUppers:  true,
		IncludeDigits:  true,
		IncludeSymbols: true,
	}

	// Call MapToCharset to test error handling when rand.Reader fails.
	_, err := MapToCharset(10, config)
	if err == nil || err.Error() != "error generating random number: mocked error from rand.Reader" {
		t.Errorf("MapToCharset() error = %v, wanted error: %v", err, "mocked error from rand.Reader")
	}
}
