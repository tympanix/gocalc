package ast

import "github.com/tympanix/gocalc/debug"

// Node represents a node in the abstract syntax tree
type Node interface {
	Calc() float64
	Print()
}

// NumberLiteral is an Node node for integer literals
type NumberLiteral float64

// Calc simply returns the integer itself
func (i NumberLiteral) Calc() float64 {
	return float64(i)
}

// Print displays the integer literal on the screen
func (i NumberLiteral) Print() {
	debug.Println(i)
}
