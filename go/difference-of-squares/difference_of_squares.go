package diffsquares

// Difference between SquareofSums and SumofSquares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

// SquareOfSums sums all numbers and returns the square.
func SquareOfSums(n int) (s int) {
	for i := 0; i <= n; i++ {
		s = s + i
	}
	return s * s
}

// SumOfSquares squares each number and then square it.
func SumOfSquares(n int) (s int) {
	// return sum([(i+1)**2 for i in range(n)])
	for i := 0; i <= n; i++ {
		s = s + i*i
	}
	return
}
