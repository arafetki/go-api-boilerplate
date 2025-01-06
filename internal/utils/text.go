package utils

import (
	"strings"
	"unicode"
)

func Capitalize(text string) string {
	if len(text) == 0 {
		return text
	}
	words := strings.Fields(text)
	for i, word := range words {
		runes := []rune(word)
		if unicode.IsLetter(runes[0]) {
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}

	return strings.Join(words, " ")
}
