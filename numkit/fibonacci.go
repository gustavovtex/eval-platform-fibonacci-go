package numkit

// MaxFibonacciInput is the largest n whose Fibonacci number fits in an int64.
const MaxFibonacciInput = 92

// Fibonacci returns F(n) for an integer n from 0 to 92 (specs/002-fibonacci), computed
// iteratively in O(n).
func Fibonacci(n int) int64 {
	if n == 0 {
		return 0
	}
	previous, current := int64(0), int64(1)
	for i := 1; i < n; i++ {
		previous, current = current, previous+current
	}
	return current
}
