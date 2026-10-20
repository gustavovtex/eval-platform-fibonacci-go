package numkit

import "errors"

// ErrOutOfRange is wrapped by every error a numkit function returns for an input it refuses.
var ErrOutOfRange = errors.New("input out of range")
