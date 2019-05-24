package ast

import "github.com/tympanix/gocalc/debug"

// Node represents a node in the abstract syntax tree
type Node interface {
	Calc() int
	Print()
}

// NumberLiteral is an Node node for integer literals
type NumberLiteral int

// Calc simply returns the integer itself
func (i NumberLiteral) Calc() int {
	return int(i)
}

// Print displays the integer literal on the screen
func (i NumberLiteral) Print() {
	debug.Println(i)
}
