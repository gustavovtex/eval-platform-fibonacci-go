# eval-platform-fibonacci-go

Small number utilities in plain Go, with no dependencies.

This repository is a fixture for the [eval platform](https://github.com/vtex/eval-platform): its
pull requests are small, each one carries its own spec under `specs/`, and the whole test suite
runs in under a second. That makes it a cheap way to check that an evaluation works end to end,
locally and in production. It is not meant to tell models apart: the tasks are too easy for that.

Twin repositories have the same history in other languages: `eval-platform-fibonacci-node`, `eval-platform-fibonacci-python`, `eval-platform-fibonacci-rust`, `eval-platform-fibonacci-csharp`.

## Usage

```go
import (
	"fmt"

	"github.com/gustavovtex/eval-platform-fibonacci-go/numkit"
)

f, _ := numkit.Factorial(5)
fmt.Println(f) // 120
```

## Tests

```sh
go test ./...
```

Go 1.22 or newer. Nothing to install: the tests use the standard library's `testing` package.

## Specs

Every change starts with a spec in `specs/NNN-name/spec.md`.

| Spec | Feature |
|---|---|
| [001-factorial](specs/001-factorial/spec.md) | `Factorial(n)` |
