package ast

import "github.com/tympanix/gocalc/debug"

// Node represents a node in the abstract syntax tree
type Node interface {
	Analyze()
	Calc() float64
	Print()
}

// NopAnalyzer is an analyzer for nodes which do not require analysis
type NopAnalyzer struct{}

// Analyze performs no operation at all
func (n NopAnalyzer) Analyze() {}

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
