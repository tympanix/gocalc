package ast

import (
	"math"

	"github.com/tympanix/tymlang/debug"
)

type binaryExpression struct {
	Lhs Node
	Rhs Node
}

// PlusOp represents an addition of integers
type PlusOp binaryExpression

// Calc returns the addition of the two operands
func (p PlusOp) Calc() int {
	return p.Lhs.Calc() + p.Rhs.Calc()
}

// Print prints the addition opeator to the screen
func (p PlusOp) Print() {
	debug.Println("PlusOp(+)")
	debug.Indent()
	p.Lhs.Print()
	p.Rhs.Print()
	debug.Outdent()
}

// MulOp represents an multiplication of integers
type MulOp binaryExpression

// Calc returns the multiplication of the two operands
func (p MulOp) Calc() int {
	return p.Lhs.Calc() * p.Rhs.Calc()
}

// Print prints the multiplcation to the screen
func (p MulOp) Print() {
	debug.Println("MulOp(*)")
	debug.Indent()
	p.Lhs.Print()
	p.Rhs.Print()
	debug.Outdent()
}

// PowOp represents an multiplication of integers
type PowOp binaryExpression

// Calc returns the multiplication of the two operands
func (p PowOp) Calc() int {
	return int(math.Pow(float64(p.Lhs.Calc()), float64(p.Rhs.Calc())))
}

// Print prints the multiplcation to the screen
func (p PowOp) Print() {
	debug.Println("PowOp(^)")
	debug.Indent()
	p.Lhs.Print()
	p.Rhs.Print()
	debug.Outdent()
}
