package bob

import (
    "strings"
    "unicode"
)

// A switch case of responses that will be returned if something is said to Bob
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(trimmedRemark, "?")
	hasLetters := strings.ContainsFunc(trimmedRemark, unicode.IsLetter)
	isShouting := hasLetters && (strings.ToUpper(trimmedRemark) == trimmedRemark)

	if isShouting && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isShouting {
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}
	return "Whatever."
}