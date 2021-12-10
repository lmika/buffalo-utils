package cli

import (
	"fmt"
	"os"
	"strings"
)

// Fatal prints v to stderr and terminates the app with exit code 1
func Fatal(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

// Fatalf prints v to stderr and terminates the app with exit code 1
func Fatalf(s string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, s, a...)
	if !strings.HasSuffix(s, "\n") {
		fmt.Fprintln(os.Stderr)
	}
	os.Exit(1)
}
