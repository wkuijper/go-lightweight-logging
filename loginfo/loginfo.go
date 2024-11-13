package loginfo

import (
	"fmt"
	"os"
)

func F(format string, args ...any) {
	fmt.Fprintf(os.Stdout, "INFO: ")
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintln(os.Stdout, "")
}
