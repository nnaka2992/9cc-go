package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid number of arguments.\n")
		os.Exit(1)
	}
	expr := os.Args[1]
	// Tokenize argument
	token := tokenize(expr)

	// Print Asembly header
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")

	// Check first argument of expr as number and
	// print the number as mov
	fmt.Printf("  mov rax, %d\n", token.expectNumber())
	for {
		if token.atEof() {
			break
		}
		if token.consume("+") {
			fmt.Printf("  add rax, %d\n", token.expectNumber())
			continue
		}
		if token.consume("-") {
			fmt.Printf("  sub rax, %d\n", token.expectNumber())
			continue
		}
	}
	fmt.Printf("  ret\n")
}
