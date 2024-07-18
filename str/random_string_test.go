package str

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	// Test case 1: Check if the length of the generated string is correct
	result := RandomString(10)
	if len(result) != 10 {
		t.Errorf("RandomString failed. Expected string length 10, got %d", len(result))
	}

	// Test case 2: Check if the generated string contains only valid characters
	for _, char := range result {
		if !isValidCharacter(char) {
			t.Errorf("RandomString failed. Invalid character found: %c", char)
		}
	}

	// Add more test cases here if needed
}

func isValidCharacter(char rune) bool {
	for _, validChar := range letters {
		if char == validChar {
			return true
		}
	}
	return false
}
