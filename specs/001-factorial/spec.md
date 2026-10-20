# 001 — factorial

## Goal

Provide `Factorial(n)`, the product of all positive integers up to `n`.

## Requirements

1. `Factorial(n int) (int64, error)` is exported from `numkit/factorial.go`.
2. `Factorial(0)` is `1`.
3. For an integer `n` from 1 to 20, `Factorial(n)` returns `n!` and a nil error. 20! is the
   largest factorial that fits in an `int64`.
4. Any other input returns `0` and an error that wraps `ErrOutOfRange` (`errors.Is` is true) when
   `n` is below 0 or above 20.

## Out of scope

Factorials above 20 (they need `math/big`).
