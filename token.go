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

func (e TokenKind) String() string {
	switch e {
	case 0:
		return "TKReserved"
	case 1:
		return "TKNum"
	case 2:
		return "TKEof"
	default:
		return "Not a valid type"
	}
}

type Token struct {
	kind   TokenKind // Type of token
	next   *Token    // Next token
	val    int       // interger: set only if kind is TkNum
	start  int       // starting location on user input
	offset int       // length of token
	rune   []rune    // token string
}

func (t *Token) String() string {
	return fmt.Sprintf(
		"kind=%s\tnext=%p\tval=%d\tstart=%d\toffset=%d\trune=%#U",
		t.kind, t.next, t.val, t.start, t.offset, t.rune,
	)

}

func (t *Token) nextToken() {
	*t = *t.next
}

func (t *Token) consume(op string) bool {
	if t.kind != TKReserved || t.offset != len([]rune(op)) || string(t.rune) != op {
		return false
	}
	t.nextToken()
	return true
}

func (t *Token) expect(op string) {
	if t.kind != TKReserved || t.offset != len([]rune(op)) || string(t.rune) != op {
		ErrorAt(t.start, "Rune is not %#U", op)
	}
	t.nextToken()
}

func (t *Token) expectNumber() int {
	if t.kind != TKNum {
		ErrorAt(t.start, "Not a number.")
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
		start:  i,
		offset: len(r),
	}
	t.next = tok
	return tok
}

func tokenize(s string) *Token {
	var head Token
	head.next = nil
	var cur *Token
	cur = &head

	pos := 0
	rs := []rune(s)
	//for i := 0; i < len(rs); {
	for 0 < len(rs) {
		if unicode.IsSpace(rs[0]) {
			pos++
			rs = rs[1:]
			continue
		} else if 2 <= len(rs) && ("==" == string(rs[:2]) || "!=" == string(rs[:2])) {
			cur = cur.newToken(TKReserved, rs[:2], pos)
			rs = rs[2:]
			pos += 2
			continue
		} else if 2 <= len(rs) && ("<=" == string(rs[:2]) || ">=" == string(rs[:2])) {
			cur = cur.newToken(TKReserved, rs[:2], pos)
			rs = rs[2:]
			pos += 2
			continue
		} else if strings.ContainsRune("+-*/()<>", rs[0]) {
			cur = cur.newToken(TKReserved, []rune{rs[0]}, pos)
			rs = rs[1:]
			pos++
			continue
		} else if IsDigit(rs[0]) {
			v, offset := RuneToInt(rs)
			cur = cur.newToken(TKNum, rs[:offset], pos)
			rs = rs[offset:]
			cur.val = v
			pos += offset
			continue
		}
		ErrorAt(pos, "Failed to tokenize")
	}
	cur.newToken(TKEof, nil, -1)
	return head.next
}

func (t Token) Print() {
	fmt.Println("==Print Token===================================================================")
	var cur *Token = &t
	for !cur.atEof() {
		fmt.Println(cur.String())
		cur.nextToken()
	}
}

func (t *Token) expr() *Node {
	return t.equality()
}

func (t *Token) equality() *Node {
	node := t.relational()
	for {
		if t.consume("==") {
			node = newBinary(NdEQ, node, t.relational())
		} else if t.consume("!=") {
			node = newBinary(NdNE, node, t.relational())
		} else {
			return node
		}
	}
}

func (t *Token) relational() *Node {
	node := t.add()
	for {
		if t.consume("<") {
			node = newBinary(NdLT, node, t.add())
		} else if t.consume(">") {
			node = newBinary(NdLT, t.add(), node)
		} else if t.consume("<=") {
			node = newBinary(NdLE, node, t.add())
		} else if t.consume(">=") {
			node = newBinary(NdLE, t.add(), node)
		} else {
			return node
		}
	}
}

func (t *Token) add() *Node {
	node := t.mul()
	for {
		if t.consume("+") {
			node = newBinary(NdAdd, node, t.mul())
		}
		if t.consume("-") {
			node = newBinary(NdSub, node, t.mul())
		} else {
			return node
		}
	}
}

func (t *Token) mul() *Node {
	node := t.unary()
	for {
		if t.consume("*") {
			node = newBinary(NdMul, node, t.unary())
		} else if t.consume("/") {
			node = newBinary(NdDiv, node, t.unary())
		} else {
			return node
		}
	}
}
func (t *Token) unary() *Node {
	if t.consume("+") {
		return t.unary()
	}
	if t.consume("-") {
		return newBinary(NdSub, newNodeNum(0), t.primary())
	} else {
		return t.primary()
	}
}

func (t *Token) primary() *Node {
	if t.consume("(") {
		node := t.expr()
		t.expect(")")
		return node
	} else {
		return newNodeNum(t.expectNumber())
	}
}
