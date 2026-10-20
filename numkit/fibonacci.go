package numkit

import (
	"fmt"
	"math/big"
)

// MaxFibonacciInput is the largest n whose Fibonacci number fits in an int64.
const MaxFibonacciInput = 92

func checkNonNegative(name string, n int) error {
	if n < 0 {
		return fmt.Errorf("%s expects n >= 0, got %d: %w", name, n, ErrOutOfRange)
	}
	return nil
}

// Fibonacci returns F(n) for an integer n from 0 to 92 (specs/002-fibonacci, 003), computed
// iteratively in O(n).
func Fibonacci(n int) (int64, error) {
	if err := checkNonNegative("Fibonacci", n); err != nil {
		return 0, err
	}
	if n > MaxFibonacciInput {
		return 0, fmt.Errorf("Fibonacci expects n <= %d, got %d; use FibonacciBig: %w", MaxFibonacciInput, n, ErrOutOfRange)
	}
	if n == 0 {
		return 0, nil
	}
	previous, current := int64(0), int64(1)
	for i := 1; i < n; i++ {
		previous, current = current, previous+current
	}
	return current, nil
}

// FibonacciBig returns F(n) for any integer n >= 0 (specs/003-fibonacci-validation-bigint),
// iteratively, in O(n).
func FibonacciBig(n int) (*big.Int, error) {
	if err := checkNonNegative("FibonacciBig", n); err != nil {
		return nil, err
	}
	if n == 0 {
		return big.NewInt(0), nil
	}
	previous, current := big.NewInt(0), big.NewInt(1)
	for i := 1; i < n; i++ {
		previous.Add(previous, current)
		previous, current = current, previous
	}
	return current, nil
}
