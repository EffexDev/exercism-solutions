package logs

import "strings"
import "unicode/utf8"

// Application identifies the application emitting the log line.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	oldChar := string(oldRune)
	newChar := string(newRune)

	return strings.ReplaceAll(log, oldChar, newChar)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) > limit {
        return false
    }
    return true
}