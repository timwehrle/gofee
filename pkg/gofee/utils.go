package gofee

// Contains checks if a given rune (character) is part of the expected character set.
// It returns true if the character is in the set, false otherwise.
func Contains(set string, char rune) bool {
	for _, c := range set {
		if c == char {
			return true
		}
	}
	return false
}
