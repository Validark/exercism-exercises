package hamming

import "errors"

const testVersion = 4

// Distance finds number of differences between two dna sequences
func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return -1, errors.New("Not Same Length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}