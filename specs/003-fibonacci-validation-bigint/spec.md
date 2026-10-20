# 003 — fibonacci: input validation and big integers

## Goal

Make `Fibonacci(n)` refuse bad input the same way `Factorial(n)` does, and add `FibonacciBig(n)`
for Fibonacci numbers beyond F(92).

## Requirements

1. `Fibonacci(n int)` now returns `(int64, error)`. It keeps its behavior for integers from 0 to
   92, and returns `0` and an error that wraps `ErrOutOfRange` when `n` is below 0 or above 92.
2. `FibonacciBig(n int) (*big.Int, error)` is exported from `numkit/fibonacci.go` and returns
   `F(n)` for any integer `n >= 0`.
3. `FibonacciBig(n)` returns `nil` and an error that wraps `ErrOutOfRange` when `n` is below 0.
   There is no upper bound.
4. For every `n` from 0 to 92, `FibonacciBig(n)` equals `Fibonacci(n)`.
5. Both functions stay iterative and O(n).
6. `README.md` lists this spec and shows `FibonacciBig` in the usage example.
