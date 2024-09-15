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
