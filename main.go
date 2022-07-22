package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid number of arguments.\n")
		os.Exit(1)
	}
	expr := os.Args[1]
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".globl main\n")
	fmt.Printf("main:\n")
	v, i := StrToInt(expr)
	fmt.Printf("  mov rax, %d\n", v)
	for i < len(expr) {
		switch expr[i] {
		case '+':
			i++
			v, offset := StrToInt(expr[i:])
			i += offset
			fmt.Printf("  add rax, %d\n", v)
			continue
		case '-':
			i++
			v, offset := StrToInt(expr[i:])
			i += offset
			fmt.Printf("  sub rax, %d\n", v)
			continue
		default:
			fmt.Fprintf(os.Stderr, "Unexpected character %s.\n", string(expr[i]))
			os.Exit(1)
		}
	}
	fmt.Printf("  ret\n")
}

func StrToInt(s string) (v int, offset int) {
	offset = strings.IndexFunc(s, func(r rune) bool { return r < '0' || r > '9' })
	if offset == -1 {
		offset = len(s)
	}
	if offset == 0 {
		return
	} // Avoid Atoi on empty string
	v, _ = strconv.Atoi(s[:offset])
	return
}
