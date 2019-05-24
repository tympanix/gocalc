package ast

import (
	"math"

	"github.com/tympanix/gocalc/debug"
)

func newBinaryExp(op string, lhs Node, rhs Node) *binaryExpression {
	return &binaryExpression{
		op:  op,
		lhs: lhs,
		rhs: rhs,
	}
}

type binaryExpression struct {
	op  string
	lhs Node
	rhs Node
}

// Print prints the binary expression
func (p *binaryExpression) Print() {
	debug.Println(p.op)
	debug.Indent()
	p.lhs.Print()
	p.rhs.Print()
	debug.Outdent()
}

func (p *binaryExpression) LHS() Node {
	return p.lhs
}

func (p *binaryExpression) RHS() Node {
	return p.rhs
}

// NewPlusOp return a new AST node for the plus operator
func NewPlusOp(lhs Node, rhs Node) *PlusOp {
	return &PlusOp{newBinaryExp("+", lhs, rhs)}
}

// PlusOp represents an addition of integers
type PlusOp struct {
	*binaryExpression
}

// Calc returns the addition of the two operands
func (p *PlusOp) Calc() float64 {
	return p.LHS().Calc() + p.RHS().Calc()
}

// NewMinusOp returns a new AST node for the minus operator
func NewMinusOp(lhs Node, rhs Node) *MinusOp {
	return &MinusOp{newBinaryExp("-", lhs, rhs)}
}

// MinusOp represents an addition of integers
type MinusOp struct {
	*binaryExpression
}

// Calc returns the addition of the two operands
func (p *MinusOp) Calc() float64 {
	return p.LHS().Calc() - p.RHS().Calc()
}

// NewMulOp returns a new AST node for the mul operator
func NewMulOp(lhs Node, rhs Node) *MulOp {
	return &MulOp{newBinaryExp("*", lhs, rhs)}
}

// MulOp represents an multiplication of integers
type MulOp struct {
	*binaryExpression
}

// Calc returns the multiplication of the two operands
func (p *MulOp) Calc() float64 {
	return p.LHS().Calc() * p.RHS().Calc()
}

// NewDivOp returns a new AST node for the div operator
func NewDivOp(lhs Node, rhs Node) *DivOp {
	return &DivOp{newBinaryExp("/", lhs, rhs)}
}

// DivOp represents an multiplication of integers
type DivOp struct {
	*binaryExpression
}

// Calc returns the multiplication of the two operands
func (p *DivOp) Calc() float64 {
	return p.LHS().Calc() / p.RHS().Calc()
}

// NewPowOp returns a new AST node for the pow operator
func NewPowOp(lhs Node, rhs Node) *PowOp {
	return &PowOp{newBinaryExp("^", lhs, rhs)}
}

// PowOp represents an multiplication of integers
type PowOp struct {
	*binaryExpression
}

// Calc returns the multiplication of the two operands
func (p *PowOp) Calc() float64 {
	return math.Pow(p.LHS().Calc(), p.RHS().Calc())
}
