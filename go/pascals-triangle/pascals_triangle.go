package pascal

func factorial(r int) (sum int) {
	sum = 1
	for i := 1; i <= r; i++ {
		sum *= i
	}
	return
}

// Triangle returns a pascal triangle
func Triangle(row int) (t [][]int) {
	for n := 0; n < row; n++ {
		r := make([]int, n+1)
		for c := 0; c <= n; c++ {
			// TODO use more efficient simplified formula
			left := factorial(n)
			right := factorial(c) * factorial(n-c)
			r[c] = left + right
		}
		t = append(t, r)
	}
	return
}
