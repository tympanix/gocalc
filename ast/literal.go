package ast

import (
	"math"

	"github.com/tympanix/gocalc/debug"
)

type literal struct {
	n float64
	t Type
	NopAnalyzer
}

func (l *literal) Calc() float64 {
	return l.n
}

func (l *literal) Print() {
	debug.Println(l.n)
}

func (l *literal) Type() Type {
	return l.t
}

// NewFloatLiteral returns the AST node for float literals
func NewFloatLiteral(n float64) Node {
	return &literal{n: n, t: FLOAT}
}

// NewIntegerLiteral returns the AST node for integer literals
func NewIntegerLiteral(n float64) Node {
	return &literal{n: n, t: INTEGER}
}

type constantExp struct {
	name  string
	t     Type
	value float64
	NopAnalyzer
}

func (c *constantExp) Calc() float64 {
	return c.value
}

func (c *constantExp) Print() {
	debug.Println(c.name)
}

func (c *constantExp) Type() Type {
	return c.t
}

// NewPiOp return the AST node for PI
func NewPiOp() Node {
	return &constantExp{name: "pi", t: FLOAT, value: math.Pi}
}

// NewEulerOp returns the AST node for Eurler's number
func NewEulerOp() Node {
	return &constantExp{name: "e", t: FLOAT, value: math.E}
}
