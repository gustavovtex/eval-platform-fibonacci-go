package numkit_test

import (
	"testing"

	"github.com/gustavovtex/eval-platform-fibonacci-go/numkit"
)

func TestFibonacciOfZeroAndOne(t *testing.T) {
	if got := numkit.Fibonacci(0); got != 0 {
		t.Errorf("Fibonacci(0) = %d; want 0", got)
	}
	if got := numkit.Fibonacci(1); got != 1 {
		t.Errorf("Fibonacci(1) = %d; want 1", got)
	}
}

func TestTheFirstFibonacciNumbers(t *testing.T) {
	want := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for n, w := range want {
		if got := numkit.Fibonacci(n); got != w {
			t.Errorf("Fibonacci(%d) = %d; want %d", n, got, w)
		}
	}
}

func TestFibonacciOfTheLargestAcceptedInput(t *testing.T) {
	if numkit.MaxFibonacciInput != 92 {
		t.Fatalf("MaxFibonacciInput = %d; want 92", numkit.MaxFibonacciInput)
	}
	if got := numkit.Fibonacci(92); got != 7540113804746346429 {
		t.Errorf("Fibonacci(92) = %d; want 7540113804746346429", got)
	}
}
