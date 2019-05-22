package ast

import "github.com/tympanix/tymlang/debug"

// Node represents a node in the abstract syntax tree
type Node interface {
	Calc() int
	Print()
}

// IntLiteral is an Node node for integer literals
type IntLiteral int

// Calc simply returns the integer itself
func (i IntLiteral) Calc() int {
	return int(i)
}

// Print displays the integer literal on the screen
func (i IntLiteral) Print() {
	debug.Println(i)
}
