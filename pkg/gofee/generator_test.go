package gofee

import "testing"

// Benchmark for password generation, tests perfomance in parallel.
// The function runs password generation in multiple goroutines, simulating real-world load.
func BenchmarkGenerate(b *testing.B) {
	length := 16
	config := PasswordConfig{
		IncludeLowers:  true,
		IncludeUppers:  true,
		IncludeDigits:  true,
		IncludeSymbols: true,
	}

	// Run the benchmark in parallel to test performance under concurrent conditions.
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Generate a password and check for errors.
			_, err := Generate(length, config)
			if err != nil {
				b.Fatalf("Error generating password: %v", err)
			}
		}
	})
}

// TestGenerateRandomness checks for randomness in password generation.
// It ensures that no duplicate passwords are generated and that character distribution is reasonably even.
func TestGenerateRandomness(t *testing.T) {
	length := 16
	numPws := 100000
	config := PasswordConfig{
		IncludeLowers:  true,
		IncludeUppers:  true,
		IncludeDigits:  true,
		IncludeSymbols: true,
	}

	// Maps to track generated passwords and character counts.
	pws := make(map[string]struct{}, numPws)
	charCount := make(map[rune]int)

	// Generate passwords and track occurrences.
	for i := 0; i < numPws; i++ {
		pw, err := Generate(length, config)
		if err != nil {
			t.Fatalf("Error generating password: %v", err)
		}

		// Check if password is unique, otherwise fail.
		if _, exists := pws[pw]; exists {
			t.Fatalf("Duplicate password found: %s", pw)
		}
		pws[pw] = struct{}{}

		// Count the occurrences of each character in the password.
		for _, c := range pw {
			charCount[c]++
		}
	}

	// Estimate the expected count for each character, with a 10% tolerance.
	expectedCount := numPws * length / len(Charset)
	tolerance := expectedCount / 10

	// Verify that the character frequencies are within the expected range.
	for c, count := range charCount {
		if count < expectedCount-tolerance || count > expectedCount+tolerance {
			t.Errorf("Character %q has count %d, expected around %d", c, count, expectedCount)
		}
	}
}

// TestGenerate tests the Generate function for various password generation scenarios.
func TestGenerate(t *testing.T) {
    type args struct {
        length int
        config PasswordConfig
    }

	// Define a list of test cases with expected outcomes.
    tests := []struct {
        name         string
        args         args
        wantedLength int
        wantedSet    string
        wantErr      bool
    }{
        {
            name: "Valid length and charset",
            args: args{
                length: 16,
                config: PasswordConfig{
                    IncludeLowers:  true,
                    IncludeUppers:  true,
                    IncludeDigits:  true,
                    IncludeSymbols: true,
                },
            },
            wantedLength: 16,
            wantedSet:    All,
            wantErr:      false,
        },
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
            wantedLength: 0,
            wantedSet:    All,
            wantErr:      true,
        },
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
            wantedLength: 0,
            wantedSet:    "",
            wantErr:      true,
        },
        {
            name: "Only lowercase",
            args: args{
                length: 8,
                config: PasswordConfig{
                    IncludeLowers:  true,
                    IncludeUppers:  false,
                    IncludeDigits:  false,
                    IncludeSymbols: false,
                },
            },
            wantedLength: 8,
            wantedSet:    Lowers,
            wantErr:      false,
        },
        {
            name: "Only uppercase",
            args: args{
                length: 8,
                config: PasswordConfig{
                    IncludeLowers:  false,
                    IncludeUppers:  true,
                    IncludeDigits:  false,
                    IncludeSymbols: false,
                },
            },
            wantedLength: 8,
            wantedSet:    Uppers,
            wantErr:      false,
        },
        {
            name: "Only digits",
            args: args{
                length: 8,
                config: PasswordConfig{
                    IncludeLowers:  false,
                    IncludeUppers:  false,
                    IncludeDigits:  true,
                    IncludeSymbols: false,
                },
            },
            wantedLength: 8,
            wantedSet:    Digits,
            wantErr:      false,
        },
        {
            name: "Only symbols",
            args: args{
                length: 8,
                config: PasswordConfig{
                    IncludeLowers:  false,
                    IncludeUppers:  false,
                    IncludeDigits:  false,
                    IncludeSymbols: true,
                },
            },
            wantedLength: 8,
            wantedSet:    Symbols,
            wantErr:      false,
        },
    }

	// Iterate over the defined test cases.
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
			// Call the Generate function with the provided arguments.
            pw, err := Generate(tt.args.length, tt.args.config)

			// Check if an error was expected or not.
            if (err != nil) != tt.wantErr {
                t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

			// Check if the generated password length matches the expected length.
            if len(pw) != tt.wantedLength {
                t.Errorf("Generate() length = %v, wantedLength %v", len(pw), tt.wantedLength)
			}

			// Validate if the generated password contains only characters from the expected set.
            for _, char := range pw {
                if !Contains(tt.wantedSet, char) {
                    t.Errorf("Generate() contains invalid character = %v, wantedSet %v", string(char), tt.wantedSet)
                }
            }
        })
    }
}