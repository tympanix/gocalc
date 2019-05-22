package ast

import "fmt"

// AST represents a node in the abstract syntax tree
type AST interface {
	Calc() int
	Print()
}

// IntLiteral is an AST node for integer literals
type IntLiteral int

// Calc simply returns the integer itself
func (i IntLiteral) Calc() int {
	return int(i)
}

// Print displays the integer literal on the screen
func (i IntLiteral) Print() {
	fmt.Println(i)
}
