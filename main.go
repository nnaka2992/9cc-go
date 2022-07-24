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
	// token.Print()
	nodes := token.program()
	// for _, node := range nodes {
	// 	node.Print()
	// }

	// Print Asembly header
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")

	// Prologue
	// allocate memory space for 26 variables
	fmt.Printf("  push rbp\n")
	fmt.Printf("  mov rbp, rsp\n")
	fmt.Printf("  sub rsp, 208\n")

	// traverse ASTs and generate Assembly code
	for _, node := range nodes {
		node.gen()
		// Must be left the result of statement
		// on the top of stack, so pop
		fmt.Printf("  pop rax\n")
	}

	// Epilogue
	// THe result of last statement must be left on the
	// top of stack, so retun the value as return code.
	fmt.Printf("  mov rsp, rbp\n")
	fmt.Printf("  pop rbp\n")
	fmt.Printf("  ret\n")
}
