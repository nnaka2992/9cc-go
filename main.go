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
	UserInput = os.Args[1]
	// Tokenize and parse expression
	token := tokenize(UserInput)
	node := token.expr()
	// Print Asembly header
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")

	// traverse AST and generate Assembly code
	node.gen()

	// Load rest of stuck and return
	fmt.Printf("  pop rax\n")
	fmt.Printf("  ret\n")
}
