package numkit

import (
	"fmt"
	"math/big"
)

// MaxSequenceLength is the longest sequence FibonacciSequence builds (specs/004-fibonacci-sequence).
const MaxSequenceLength = 10000

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

// FibonacciSequence returns [F(0), ..., F(count - 1)], built in one pass (specs/004-fibonacci-sequence).
// Every call builds and returns a new slice of new values.
func FibonacciSequence(count int) ([]*big.Int, error) {
	if err := checkNonNegative("FibonacciSequence", count); err != nil {
		return nil, err
	}
	if count > MaxSequenceLength {
		return nil, fmt.Errorf("FibonacciSequence expects count <= %d, got %d: %w", MaxSequenceLength, count, ErrOutOfRange)
	}
	sequence := make([]*big.Int, 0, count)
	previous, current := big.NewInt(0), big.NewInt(1)
	for i := 0; i < count; i++ {
		sequence = append(sequence, new(big.Int).Set(previous))
		next := new(big.Int).Add(previous, current)
		previous, current = current, next
	}
	return sequence, nil
}
