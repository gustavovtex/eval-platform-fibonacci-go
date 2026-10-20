package numkit_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/gustavovtex/eval-platform-fibonacci-go/numkit"
)

func TestFibonacciOfZeroAndOne(t *testing.T) {
	for n, want := range map[int]int64{0: 0, 1: 1} {
		if got, err := numkit.Fibonacci(n); err != nil || got != want {
			t.Errorf("Fibonacci(%d) = %d, %v; want %d, nil", n, got, err, want)
		}
	}
}

func TestTheFirstFibonacciNumbers(t *testing.T) {
	want := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for n, w := range want {
		if got, err := numkit.Fibonacci(n); err != nil || got != w {
			t.Errorf("Fibonacci(%d) = %d, %v; want %d, nil", n, got, err, w)
		}
	}
}

func TestFibonacciOfTheLargestAcceptedInput(t *testing.T) {
	if numkit.MaxFibonacciInput != 92 {
		t.Fatalf("MaxFibonacciInput = %d; want 92", numkit.MaxFibonacciInput)
	}
	if got, err := numkit.Fibonacci(92); err != nil || got != 7540113804746346429 {
		t.Errorf("Fibonacci(92) = %d, %v; want 7540113804746346429, nil", got, err)
	}
}

func TestFibonacciRejectsOutOfRangeIntegers(t *testing.T) {
	for _, n := range []int{-1, 93} {
		if _, err := numkit.Fibonacci(n); !errors.Is(err, numkit.ErrOutOfRange) {
			t.Errorf("Fibonacci(%d) error = %v; want ErrOutOfRange", n, err)
		}
	}
}

func TestFibonacciBigAgreesWithFibonacciUpToTheLimit(t *testing.T) {
	for n := 0; n <= numkit.MaxFibonacciInput; n++ {
		small, err := numkit.Fibonacci(n)
		if err != nil {
			t.Fatalf("Fibonacci(%d): %v", n, err)
		}
		big_, err := numkit.FibonacciBig(n)
		if err != nil || big_.Cmp(big.NewInt(small)) != 0 {
			t.Errorf("FibonacciBig(%d) = %v, %v; want %d", n, big_, err, small)
		}
	}
}

func TestFibonacciBigGoesBeyondTheLimit(t *testing.T) {
	for n, want := range map[int]string{79: "14472334024676221", 100: "354224848179261915075"} {
		if got, err := numkit.FibonacciBig(n); err != nil || got.String() != want {
			t.Errorf("FibonacciBig(%d) = %v, %v; want %s", n, got, err, want)
		}
	}
}

func TestFibonacciBigValidatesInputWithNoUpperBound(t *testing.T) {
	if _, err := numkit.FibonacciBig(-1); !errors.Is(err, numkit.ErrOutOfRange) {
		t.Errorf("FibonacciBig(-1) error = %v; want ErrOutOfRange", err)
	}
	if _, err := numkit.FibonacciBig(1000); err != nil {
		t.Errorf("FibonacciBig(1000) error = %v; want nil", err)
	}
}
