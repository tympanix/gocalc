package ast

import (
	"fmt"
	"math"

	"github.com/tympanix/gocalc/debug"
)

type binaryAnalyzer func(*binaryExp) error

var defaultBinaryAnalyzer = func(b *binaryExp) error {
	if err := b.LHS().Analyze(); err != nil {
		return err
	}
	if err := b.RHS().Analyze(); err != nil {
		return err
	}
	return nil
}

var integerBinaryAnalyzer = func(b *binaryExp) error {
	if err := defaultBinaryAnalyzer(b); err != nil {
		return err
	}
	if b.LHS().Type() != INTEGER || b.RHS().Type() != INTEGER {
		return fmt.Errorf("illegal operands for: %s", b.name)
	}
	return nil

}

type binaryTyper func(b *binaryExp) Type

var defaultBinaryTyper = func(b *binaryExp) Type {
	if b.LHS().Type() == INTEGER && b.RHS().Type() == INTEGER {
		return INTEGER
	}
	return FLOAT
}

var floatBinaryTyper = func(b *binaryExp) Type {
	return FLOAT
}

var integerBinaryTyper = func(b *binaryExp) Type {
	return INTEGER
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
	return b.a(b)
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
	return b.t(b)
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

// NewBitwiseAndOp returns the AST node for bitwise and (&) operator
func NewBitwiseAndOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "&",
		lhs:  lhs,
		rhs:  rhs,
		a:    integerBinaryAnalyzer,
		t:    integerBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return float64(int64(a) & int64(b))
		},
	}
}

// NewBitwiseOrOp returns the AST node for bitwise or (|) operator
func NewBitwiseOrOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "|",
		lhs:  lhs,
		rhs:  rhs,
		a:    integerBinaryAnalyzer,
		t:    integerBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return float64(int64(a) | int64(b))
		},
	}
}

// NewModOp returns the AST node for mod (%) operator
func NewModOp(lhs Node, rhs Node) Node {
	return &binaryExp{
		name: "%",
		lhs:  lhs,
		rhs:  rhs,
		a:    integerBinaryAnalyzer,
		t:    integerBinaryTyper,
		fn: func(a float64, b float64) float64 {
			return float64(int64(a) % int64(b))
		},
	}
}
