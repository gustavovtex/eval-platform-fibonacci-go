# 004 — fibonacci sequence

## Goal

Provide `FibonacciSequence(count)`, the first `count` Fibonacci numbers, for callers that need the
whole sequence instead of a single term.

## Requirements

1. `FibonacciSequence(count int) ([]*big.Int, error)` is exported from `numkit/fibonacci.go` and
   returns `count` values: `[F(0), F(1), ..., F(count - 1)]`. `FibonacciSequence(0)` is empty.
2. `count` is validated like `FibonacciBig(n)`: an error that wraps `ErrOutOfRange` when it is
   below 0. It is also capped: the same error when it is above 10000.
3. The sequence is built in a single pass, O(count). Calling `FibonacciBig` once per term
   (O(count²)) does not meet this requirement.
4. Every call returns a new slice of new `*big.Int` values. A caller may sort, modify or truncate
   what it received, and no later call returns a different result because of it. In particular, a
   package-level cache, if one is used, is never handed to the caller.
5. `FibonacciSequence(10000)` completes in under 100 ms on a current laptop.
6. `README.md` lists this spec and shows `FibonacciSequence` in the usage example.
