package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

func abbreviate(str string) (abbrv string) {
	// Split on all characters that are not alphannumeric
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(str, " ")
	for _, word := range strings.Split(safe, " ") {
		abbrv += strings.ToUpper(word[:1])
		// Check if not an acronym already
		if strings.ToUpper(word) != word {
			// Handle special case where abbreviation text is in middle of word
			for k, l := range word {
				if unicode.IsUpper(l) && k != 0 {
					abbrv += string(l)
				}
			}
		}
	}
	return
}
