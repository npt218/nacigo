package dbg

import (
	"fmt"
)

func E(format string, a ...any) {
	fmt.Printf("[Err] " + format + " -- ", a...)
}