package pangram

func IsPangram(sentence string) bool {
	// Determines if a sentence is a panagram.
	// A panagram is a sentance with only one of each letter
	if len(sentence) == 0 {
		return false
	}
	var letters = make(map[rune]bool)
	for _, letter := range sentence {
		if letter != 32 && letter < 'A' || letter > 'z' {
			return false
		}
		if letters[letter] == true {
			return false
		}
		letters[letter] = true
	}
	return true
}
