package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid number of arguments.\n")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(2)
	}

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")
	fmt.Printf("  mov rax, %d\n", i)
	fmt.Printf("  ret\n")
}
