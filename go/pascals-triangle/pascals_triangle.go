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
			// top := factorial(n)
			// bot := factorial(c) * factorial(n-c)
			// r[c] = top / bot
			left := (c - 1) / n
			right := (c - 1) / (n - 1)
			r[c] = left + right
		}
		t = append(t, r)
	}
	return
}
