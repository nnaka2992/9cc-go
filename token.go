package main

import (
	"fmt"
	"unicode"
)

type TokenKind int

const (
	TKReserved TokenKind = iota
	TKNum
	TKEof
)

type Token struct {
	kind TokenKind
	next *Token
	val  int
	rune []rune
}

func (t *Token) nextToken() {
	*t = *t.next
}

func (t *Token) consume(op string) bool {
	if t.kind != TKReserved || string(t.rune) != op {
		return false
	}
	t.nextToken()
	return true
}

func (t *Token) expect(op string) {
	if t.kind != TKReserved || string(t.rune) != op {
		Errorf("Rune is not %#U", op)
	}
	t.nextToken()
}

func (t *Token) expectNumber() int {
	if t.kind != TKNum {
		Errorf("Not a number.")
	}
	val := t.val
	t.nextToken()
	return val
}

func (t Token) atEof() bool {
	return t.kind == TKEof
}

func (t *Token) newToken(kind TokenKind, r []rune) *Token {
	tok := &Token{
		kind: kind,
		rune: r,
	}
	t.next = tok
	return tok
}

func tokenize(s string) *Token {
	var head Token
	head.next = nil
	var cur *Token
	cur = &head

	rs := []rune(s)
	for i := 0; i < len(rs); {
		if unicode.IsSpace(rs[i]) {
			i++
			continue
		}
		if rs[i] == '+' || rs[i] == '-' {
			cur = cur.newToken(TKReserved, []rune{rs[i]})
			i++
			continue
		}
		if IsDigit(rs[i]) {
			v, offset := RuneToInt(rs[i:])
			cur = cur.newToken(TKNum, rs[i:i+offset])
			cur.val = v
			i += offset
			continue
		}
		Errorf("Failed to tokenize")
	}
	cur.newToken(TKEof, nil)
	return head.next
}

func (t Token) PrintTokens() {
	var cur *Token = &t
	pos := 1
	for !cur.atEof() {
		fmt.Printf("%d: %#U\n", pos, cur.rune)
		cur.nextToken()
		pos++
	}
}
