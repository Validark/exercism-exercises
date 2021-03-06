package raindrops

import "strconv"

const testVersion = 2

func Factorize(i int) []int {
	// Takes a number and breaks it into all factors of that number as an Array
	var factors []int

	for f := 1; i >= f; f++ {
		if i%f == 0 {
			factors = append(factors, f)
		}
	}
	return factors
}

func Convert(str int) (statement string) {
	// Converts a statement into a combination of PlingPlangPlong and num depending on factors
	for _, i := range Factorize(str) {
		switch i {
		case 3:
			statement = statement + "Pling"
		case 5:
			statement = statement + "Plang"
		case 7:
			statement = statement + "Plong"
		}
	}

	if len(statement) == 0 {
		statement = strconv.Itoa(str)
	}
	return
}

// The test program has a benchmark too.  How fast does your Convert convert?
