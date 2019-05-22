package ast

import "fmt"

type binaryExpression struct {
	Lhs AST
	Rhs AST
}

// PlusOp represents an addition of integers
type PlusOp binaryExpression

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
type MulOp binaryExpression

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
