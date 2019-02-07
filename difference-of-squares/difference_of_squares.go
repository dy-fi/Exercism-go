package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	formula := (n * (n + 1)) / 2
	return formula * formula
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i < n+1; i++ {
		sum += (i * i)
	}
	return sum
}

// Difference returns the difference between the square of sums and the sum of squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
