package main

import (
	"fmt"
	"strings"
	"unicode"
)

var UserInput string

type TokenKind int

const (
	TKReserved TokenKind = iota
	TKNum
	TKEof
)

type Token struct {
	kind   TokenKind
	next   *Token
	val    int
	offset int
	rune   []rune
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
		ErrorAt(t.offset, "Rune is not %#U", op)
	}
	t.nextToken()
}

func (t *Token) expectNumber() int {
	if t.kind != TKNum {
		ErrorAt(t.offset, "Not a number.")
	}
	val := t.val
	t.nextToken()
	return val
}

func (t Token) atEof() bool {
	return t.kind == TKEof
}

func (t *Token) newToken(kind TokenKind, r []rune, i int) *Token {
	tok := &Token{
		kind:   kind,
		rune:   r,
		offset: i,
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
		} else if strings.ContainsRune("+-*/()", rs[i]) {
			cur = cur.newToken(TKReserved, []rune{rs[i]}, i)
			i++
			continue
		} else if IsDigit(rs[i]) {
			v, offset := RuneToInt(rs[i:])
			cur = cur.newToken(TKNum, rs[i:i+offset], i+offset)
			cur.val = v
			i += offset
			continue
		}
		ErrorAt(i, "Failed to tokenize")
	}
	cur.newToken(TKEof, nil, -1)
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

func (t *Token) expr() *Node {
	node := t.mul()
	for {
		if t.consume("+") {
			node = newNode(NdAdd, node, t.mul())
		} else if t.consume("-") {
			node = newNode(NdSub, node, t.mul())
		} else {
			return node
		}
	}
}

func (t *Token) mul() *Node {
	node := t.unary()
	for {
		if t.consume("*") {
			node = newNode(NdMul, node, t.unary())
		} else if t.consume("/") {
			node = newNode(NdDiv, node, t.unary())
		} else {
			return node
		}
	}
}

func (t *Token) primary() *Node {
	if t.consume("(") {
		node := t.expr()
		t.expect(")")
		return node
	}
	return newNodeNum(t.expectNumber())
}

func (t *Token) unary() *Node {
	if t.consume("+") {
		return t.primary()
	}
	if t.consume("-") {
		return newNode(NdSub, newNodeNum(0), t.primary())
	}
	return t.primary()
}
