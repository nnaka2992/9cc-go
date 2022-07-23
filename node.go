package main

import (
	"fmt"
)

type NodeKind int

const (
	NdAdd NodeKind = iota // +
	NdSub                 // -
	NdMul                 // *
	NdDiv                 // /
	NdNum                 // Integer
)

type Node struct {
	kind NodeKind // Node type
	lhs  *Node    // left side leef
	rhs  *Node    // right side leef
	val  int      // Use only if kind is NdNum
}

func newNode(kind NodeKind, lhs *Node, rhs *Node) *Node {
	return &Node{
		kind: kind,
		lhs:  lhs,
		rhs:  rhs,
	}
}

func newNodeNum(val int) *Node {
	return &Node{
		kind: NdNum,
		val:  val,
	}
}

func (n *Node) gen() {
	if n.kind == NdNum {
		fmt.Printf("  push %d\n", n.val)
		return
	}
	n.lhs.gen()
	n.rhs.gen()

	fmt.Printf("  pop rdi\n")
	fmt.Printf("  pop rax\n")
	switch n.kind {
	case NdAdd:
		fmt.Printf("  add rax, rdi\n")
		break
	case NdSub:
		fmt.Printf("  sub rax, rdi\n")
		break
	case NdMul:
		fmt.Printf("  imul rax, rdi\n")
		break
	case NdDiv:
		fmt.Printf("  cqo\n")
		fmt.Printf("  idiv rdi\n")
		break
	}

	fmt.Printf("  push rax\n")
}
