package secret

import "strconv"

const testVersion = 1

// Handshake passes back an array of strings dependenting on bit significance
func Handshake(i uint) (a []string) {

	// Convert Number to Base 2
	n := strconv.FormatInt(int64(i), 2)
	if n[len(n)-1] == '1' {
		a = append(a, "wink")
	}
	if len(n) > 1 && n[len(n)-2] == '1' {
		a = append(a, "double blink")
	}

	if len(n) > 2 && n[len(n)-3] == '1' {
		a = append(a, "close your eyes")
	}

	if len(n) > 3 && n[len(n)-4] == '1' {
		a = append(a, "jump")
	}

	// If greater than 16 we need to reverse, but only on odd sets
	if (i/16)%2 == 1 {
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	}

	return
}
