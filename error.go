package main

import (
	"fmt"
	"os"
)

func ErrorAt(loc int, format string, a ...interface{}) {
	pos := loc - len(UserInput)

	fmt.Fprintf(os.Stderr, "%s\n", UserInput)
	fmt.Fprintf(os.Stderr, "%*s", pos, " ")
	fmt.Fprintf(os.Stderr, "^ ")
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
