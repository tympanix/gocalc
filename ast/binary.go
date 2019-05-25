package ast

import (
	"fmt"
	"math"

	"github.com/tympanix/gocalc/debug"
)

type binaryAnalyzer func(*binaryExp, Node, Node) error

var defaultBinaryAnalyzer = func(b *binaryExp, lhs Node, rhs Node) error {
	if err := lhs.Analyze(); err != nil {
		return err
	}
	if err := rhs.Analyze(); err != nil {
		return err
	}
	return nil
}

var integerBinaryAnalyzer = func(b *binaryExp, lhs Node, rhs Node) error {
	if err := defaultBinaryAnalyzer(b, lhs, rhs); err != nil {
		return err
	}
	if lhs.Type() != INTEGER || rhs.Type() != INTEGER {
		return fmt.Errorf("illegal operands: %s", b.name)
	}
	return nil

}

type binaryTyper func(Node, Node) Type

var defaultBinaryTyper = func(lhs Node, rhs Node) Type {
	if lhs.Type() == INTEGER && rhs.Type() == INTEGER {
		return INTEGER
	}
	return FLOAT
}

var floatBinaryTyper = func(lhs Node, rhs Node) Type {
	return FLOAT
}

type binaryExp struct {
	name string
	lhs  Node
	rhs  Node
	a    binaryAnalyzer
	t    binaryTyper
	fn   func(float64, float64) float64
}

// Analyse performs analysis on the right- and lef-hand side
func (b *binaryExp) Analyze() error {
	return b.a(b, b.LHS(), b.RHS())
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
	return b.t(b.LHS(), b.RHS())
}

func (b *binaryExp) Calc() float64 {
	return b.fn(b.LHS().Calc(), b.RHS().Calc())
}

func (b *binaryExp) LHS() Node {
	return b.lhs
}

func (b *binaryExp) RHS() Node {
	return b.rhs
}

// NewPlusOp return a new AST node for the plus operator
func NewPlusOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "+",
		lhs:  lhs,
		rhs:  rhs,
		a:    defaultBinaryAnalyzer,
		t:    defaultBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return a + b
		},
	}
}

// NewMinusOp returns a new AST node for the minus operator
func NewMinusOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "-",
		lhs:  lhs,
		rhs:  rhs,
		a:    defaultBinaryAnalyzer,
		t:    defaultBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return a - b
		},
	}
}

// NewMulOp returns a new AST node for the mul operator
func NewMulOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "*",
		lhs:  lhs,
		rhs:  rhs,
		a:    defaultBinaryAnalyzer,
		t:    defaultBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return a * b
		},
	}
}

// NewDivOp returns a new AST node for the div operator
func NewDivOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "/",
		lhs:  lhs,
		rhs:  rhs,
		a:    defaultBinaryAnalyzer,
		t:    floatBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return a / b
		},
	}
}

// NewPowOp returns a new AST node for the pow operator
func NewPowOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "^",
		lhs:  lhs,
		rhs:  rhs,
		a:    defaultBinaryAnalyzer,
		t:    defaultBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return math.Pow(a, b)
		},
	}
}

// NewLogicalAndOp returns the AST node for logical and (&) operator
func NewLogicalAndOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "&",
		lhs:  lhs,
		rhs:  rhs,
		a:    integerBinaryAnalyzer,
		t:    defaultBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return float64(int64(a) & int64(b))
		},
	}
}
