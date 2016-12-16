package bob // package name must match the package name in bob_test.go
import "strings"

const testVersion = 2 // same as targetTestVersion

func Hey(what string) string {

	what = strings.TrimSpace(what)
	switch {
	case len(what) == 0:
		return "Fine. Be that way!"
	case strings.ToUpper(what) == what && hasUpper(what):
		return "Whoa, chill out!"
	case what[len(what)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

// Checks if any characters are upper.
func hasUpper(w string) bool {
	for _, r := range w {
		if r > 'A' && r < 'Z' {
			return true
		}
	}
	return false
}
