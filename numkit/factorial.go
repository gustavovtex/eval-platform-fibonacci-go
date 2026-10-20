package numkit

import "fmt"

// MaxFactorialInput is the largest n whose factorial fits in an int64.
const MaxFactorialInput = 20

// Factorial returns n! for an integer n from 0 to 20 (specs/001-factorial).
func Factorial(n int) (int64, error) {
	if n < 0 || n > MaxFactorialInput {
		return 0, fmt.Errorf("factorial expects 0 <= n <= %d, got %d: %w", MaxFactorialInput, n, ErrOutOfRange)
	}
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result, nil
}
