package logwarn

import (
	"fmt"
	"os"
)

func F(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "WARNING: ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr, "")
}
