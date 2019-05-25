package ast

import (
	"fmt"
	"math"

	"github.com/tympanix/gocalc/debug"
)

func newBinaryExp(name string, lhs Node, rhs Node) *binaryExp {
	return &binaryExp{
		name: name,
		lhs:  lhs,
		rhs:  rhs,
	}
}

type binaryExp struct {
	name string
	lhs  Node
	rhs  Node
}

// Analyse performs analysis on the right- and lef-hand side
func (b *binaryExp) Analyze() error {
	if err := b.LHS().Analyze(); err != nil {
		return err
	}
	if err := b.RHS().Analyze(); err != nil {
		return err
	}
	return nil
}

// Print prints the binary expression
func (b *binaryExp) Print() {
	debug.Println(b.name)
	debug.Indent()
	b.LHS().Print()
	b.RHS().Print()
	debug.Outdent()
}

func (b *binaryExp) Type() Type {
	if b.LHS().Type() == INTEGER && b.RHS().Type() == INTEGER {
		return INTEGER
	}
	return FLOAT
}

func (b *binaryExp) LHS() Node {
	return b.lhs
}

func (b *binaryExp) RHS() Node {
	return b.rhs
}

// NewPlusOp return a new AST node for the plus operator
func NewPlusOp(lhs Node, rhs Node) *PlusOp {
	return &PlusOp{newBinaryExp("+", lhs, rhs)}
}

// PlusOp represents an addition of integers
type PlusOp struct {
	*binaryExp
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
	*binaryExp
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
	*binaryExp
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
	*binaryExp
}

// Calc returns the multiplication of the two operands
func (p *DivOp) Calc() float64 {
	return p.LHS().Calc() / p.RHS().Calc()
}

// Type returns float since result can be real number
func (p *DivOp) Type() Type {
	return FLOAT
}

// NewPowOp returns a new AST node for the pow operator
func NewPowOp(lhs Node, rhs Node) *PowOp {
	return &PowOp{newBinaryExp("^", lhs, rhs)}
}

// PowOp represents an multiplication of integers
type PowOp struct {
	*binaryExp
}

// Calc returns the multiplication of the two operands
func (p *PowOp) Calc() float64 {
	return math.Pow(p.LHS().Calc(), p.RHS().Calc())
}

// NewLogicalAndOp returns the AST node for logical and (&) operator
func NewLogicalAndOp(lhs Node, rhs Node) Node {
	return &LogicalAndOp{newBinaryExp("&", lhs, rhs)}
}

// LogicalAndOp represents an multiplication of integers
type LogicalAndOp struct {
	*binaryExp
}

// Calc returns the multiplication of the two operands
func (p *LogicalAndOp) Calc() float64 {
	return float64(int64(p.LHS().Calc()) & int64(p.RHS().Calc()))
}

// Analyze makes sure both lhs and rhs are integer values
func (p *LogicalAndOp) Analyze() error {
	if err := p.binaryExp.Analyze(); err != nil {
		return err
	}
	if p.LHS().Type() != INTEGER || p.RHS().Type() != INTEGER {
		return fmt.Errorf("illegal operands: %s", p.name)
	}
	return nil
}
