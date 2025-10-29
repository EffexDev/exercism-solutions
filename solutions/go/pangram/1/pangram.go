package pangram

import "strings"

// IsPangram determines if a string contains every letter of the alphabet.
func IsPangram(input string) bool {
	seen := make(map[rune]bool)

	lowerInput := strings.ToLower(input)

	for _, char := range lowerInput {
		if char >= 'a' && char <= 'z' {
			seen[char] = true
		}
	}
	return len(seen) == 26
}
