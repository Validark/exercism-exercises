package pangram

import (
	"unicode"
)

func IsPangram(sentence string) bool {
	// Determines if a sentence is a panagram.
	// A panagram is a sentence with only one of each letter
	if len(sentence) == 0 {
		return false
	}
	var letters = make(map[rune]bool)
	for _, letter := range sentence {
		letters[unicode.ToLower(letter)] = true
	}
	for _,letter := range "abcdefghijklmnopqrstuvwxyz" {
		if !letters[unicode.ToLower(letter)] {
			return false
		}
	}
	return true
}
