package logerror

import (
	"fmt"
	"os"
)

func F(format string, args ...any) {
	fmt.Fprintf(os.Stdout, "DEBUG: ")
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintln(os.Stdout, "")
}
