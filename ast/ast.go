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

// PlusOp represents an addition of integers
type PlusOp struct {
	Lhs AST
	Rhs AST
}

// Calc returns the addition of the two operands
func (p PlusOp) Calc() int {
	return p.Lhs.Calc() + p.Rhs.Calc()
}

// Print prints the addition opeator to the screen
func (p PlusOp) Print() {
	fmt.Println("PlusOp")
	p.Lhs.Print()
	p.Rhs.Print()
}

// MulOp represents an multiplication of integers
type MulOp struct {
	Lhs AST
	Rhs AST
}

// Calc returns the multiplication of the two operands
func (p MulOp) Calc() int {
	return p.Lhs.Calc() * p.Rhs.Calc()
}

// Print prints the multiplcation to the screen
func (p MulOp) Print() {
	fmt.Println("MulOp")
	p.Lhs.Print()
	p.Rhs.Print()
}
