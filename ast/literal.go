package ast

import (
	"math"

	"github.com/tympanix/gocalc/debug"
)

// NewNumberLiteral returns the AST node for number literals
func NewNumberLiteral(n float64) Node {
	return &NumberLiteral{n: n}
}

// NumberLiteral is an Node node for integer literals
type NumberLiteral struct {
	n float64
	NopAnalyzer
}

// Calc simply returns the integer itself
func (i *NumberLiteral) Calc() float64 {
	return i.n
}

// Print displays the integer literal on the screen
func (i *NumberLiteral) Print() {
	debug.Println(i)
}

type constantExp struct {
	name  string
	value float64
	NopAnalyzer
}

func (c *constantExp) Calc() float64 {
	return c.value
}

func (c *constantExp) Print() {
	debug.Println(c.name)
}

// NewPiOp return the AST node for PI
func NewPiOp() Node {
	return &constantExp{name: "pi", value: math.Pi}
}

// NewEulerOp returns the AST node for Eurler's number
func NewEulerOp() Node {
	return &constantExp{name: "e", value: math.E}
}
