package raindrops

import "strconv"

const testVersion = 2

func factorize(i int) []int {
	var factors []int

	for f := 1; i >= f; f++ {
		if i%f == 0 {
			factors = append(factors, f)
		}
	}
	return factors
}

func Convert(str int) string {
	statement := ""
	for _, i := range factorize(str) {
		if i == 3 {
			statement = statement + "Pling"
		} else if i == 5 {
			statement = statement + "Plang"
		} else if i == 7 {
			statement = statement + "Plong"
		}
	}

	if len(statement) == 0 {
		statement = statement + strconv.Itoa(str)
	}
	return statement
}

// The test program has a benchmark too.  How fast does your Convert convert?
