package numkit_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/gustavovtex/eval-platform-fibonacci-go/numkit"
)

func TestFibonacciSequenceOfZeroIsEmpty(t *testing.T) {
	got, err := numkit.FibonacciSequence(0)
	if err != nil || len(got) != 0 {
		t.Errorf("FibonacciSequence(0) = %v, %v; want empty, nil", got, err)
	}
}

func TestFibonacciSequenceStartsWithTheFirstFibonacciNumbers(t *testing.T) {
	want := []int64{0, 1, 1, 2, 3, 5, 8, 13}
	got, err := numkit.FibonacciSequence(len(want))
	if err != nil || len(got) != len(want) {
		t.Fatalf("FibonacciSequence(%d) = %v, %v", len(want), got, err)
	}
	for i, w := range want {
		if got[i].Cmp(big.NewInt(w)) != 0 {
			t.Errorf("term %d = %v; want %d", i, got[i], w)
		}
	}
}

func TestEveryTermMatchesFibonacciBig(t *testing.T) {
	sequence, err := numkit.FibonacciSequence(200)
	if err != nil || len(sequence) != 200 {
		t.Fatalf("FibonacciSequence(200) = %d terms, %v", len(sequence), err)
	}
	for n, term := range sequence {
		want, _ := numkit.FibonacciBig(n)
		if term.Cmp(want) != 0 {
			t.Errorf("term %d = %v; want %v", n, term, want)
		}
	}
}

func TestFibonacciSequenceAcceptsTheCapAndRejectsBeyondIt(t *testing.T) {
	if numkit.MaxSequenceLength != 10000 {
		t.Fatalf("MaxSequenceLength = %d; want 10000", numkit.MaxSequenceLength)
	}
	if got, err := numkit.FibonacciSequence(10000); err != nil || len(got) != 10000 {
		t.Errorf("FibonacciSequence(10000) = %d terms, %v; want 10000, nil", len(got), err)
	}
	if _, err := numkit.FibonacciSequence(10001); !errors.Is(err, numkit.ErrOutOfRange) {
		t.Errorf("FibonacciSequence(10001) error = %v; want ErrOutOfRange", err)
	}
}

func TestFibonacciSequenceValidatesCount(t *testing.T) {
	if _, err := numkit.FibonacciSequence(-1); !errors.Is(err, numkit.ErrOutOfRange) {
		t.Errorf("FibonacciSequence(-1) error = %v; want ErrOutOfRange", err)
	}
}

func TestEveryCallReturnsNewValues(t *testing.T) {
	first, _ := numkit.FibonacciSequence(5)
	first[4].SetInt64(-1)
	first[0], first[1] = first[1], first[0]
	second, _ := numkit.FibonacciSequence(5)
	for i, w := range []int64{0, 1, 1, 2, 3} {
		if second[i].Cmp(big.NewInt(w)) != 0 {
			t.Errorf("term %d = %v after the first result was modified; want %d", i, second[i], w)
		}
	}
}
