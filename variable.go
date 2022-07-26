package main

var locals *LVar

type LVar struct {
	next   *LVar
	name   []rune
	len    int
	offset int
}
