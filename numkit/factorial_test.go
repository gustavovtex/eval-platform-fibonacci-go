package numkit_test

import (
	"errors"
	"testing"

	"github.com/gustavovtex/eval-platform-fibonacci-go/numkit"
)

func TestFactorialOfZeroAndOneIsOne(t *testing.T) {
	for _, n := range []int{0, 1} {
		got, err := numkit.Factorial(n)
		if err != nil || got != 1 {
			t.Errorf("Factorial(%d) = %d, %v; want 1, nil", n, got, err)
		}
	}
}

func TestFactorialOfSmallIntegers(t *testing.T) {
	for n, want := range map[int]int64{5: 120, 10: 3628800} {
		got, err := numkit.Factorial(n)
		if err != nil || got != want {
			t.Errorf("Factorial(%d) = %d, %v; want %d, nil", n, got, err, want)
		}
	}
}

func TestFactorialOfTheLargestAcceptedInput(t *testing.T) {
	if numkit.MaxFactorialInput != 20 {
		t.Fatalf("MaxFactorialInput = %d; want 20", numkit.MaxFactorialInput)
	}
	got, err := numkit.Factorial(20)
	if err != nil || got != 2432902008176640000 {
		t.Errorf("Factorial(20) = %d, %v; want 2432902008176640000, nil", got, err)
	}
}

func TestFactorialRejectsOutOfRangeIntegers(t *testing.T) {
	for _, n := range []int{-1, 21} {
		if _, err := numkit.Factorial(n); !errors.Is(err, numkit.ErrOutOfRange) {
			t.Errorf("Factorial(%d) error = %v; want ErrOutOfRange", n, err)
		}
	}
}
