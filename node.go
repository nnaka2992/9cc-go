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
	NdEQ                  // ==
	NdNE                  // !=
	NdLT                  // <
	NdLE                  // <=
)

func (e NodeKind) String() string {
	switch e {
	case 0:
		return "NdAdd"
	case 1:
		return "NdSub"
	case 2:
		return "NdMul"
	case 3:
		return "NdDiv"
	case 4:
		return "NdNum"
	case 5:
		return "NdEQ"
	case 6:
		return "NdNE"
	case 7:
		return "NdLT"
	case 8:
		return "NdLE"
	default:
		return "Not a valid type"
	}
}

// BNF
// expr       = equality
// equality   = relational ("==" relational | "!=" relational)*
// relational = add ("<" add | "<=" add | ">" add | ">=" add)*
// add        = mul ("+" mul | "-" mul)*
// mul        = unary ("*" unary | "/" unary)*
// unary      = ("+" | "-")? primary
// primary    = num | "(" expr ")"

// Node replisent AST of compiler
type Node struct {
	kind NodeKind // Node type
	lhs  *Node    // left side leef
	rhs  *Node    // right side leef
	val  int      // Use only if kind is NdNum
}

func (n *Node) String() string {
	return fmt.Sprintf(
		"kind=%s\tlhs=%p\trhs=%p\tval=%d\n",
		n.kind, n.lhs, n.rhs, n.val,
	)
}

func newNode(kind NodeKind) *Node {
	return &Node{
		kind: kind,
	}
}

func newBinary(kind NodeKind, lhs *Node, rhs *Node) *Node {
	node := newNode(kind)
	node.lhs = lhs
	node.rhs = rhs
	return node
}

func newNodeNum(val int) *Node {
	node := newNode(NdNum)
	node.val = val
	return node
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
	case NdEQ:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  sete al\n")
		fmt.Printf("  movzb rax, al\n")
	case NdNE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setne al\n")
		fmt.Printf("  movzb rax, al\n")
	case NdLT:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax, al\n")
	case NdLE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setle al\n")
		fmt.Printf("  movzb rax, al\n")
	}

	fmt.Printf("  push rax\n")
}

func (n *Node) Print() {
	fmt.Println("==Print AST=====================================================================")
	var Print func(n *Node)
	Print = func(n *Node) {
		if n == nil {
			return
		}
		Print(n.lhs)
		Print(n.rhs)
		fmt.Printf(n.String())
	}
	Print(n)
}
