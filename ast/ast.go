package ast

// Type denotes the result type of an AST node
type Type int

const (
	UNKNOWN Type = iota
	INTEGER
	FLOAT
)

// IntType is a embeddable helper struct for integer types
type IntType struct{}

// Type return the integer type
func (i IntType) Type() Type {
	return INTEGER
}

// FloatType is a embeddable helper struct for float types
type FloatType struct{}

// Type returns the float type
func (f FloatType) Type() Type {
	return FLOAT
}

// Node represents a node in the abstract syntax tree
type Node interface {
	Analyze()
	Type() Type
	Calc() float64
	Print()
}

// NopAnalyzer is an analyzer for nodes which do not require analysis
type NopAnalyzer struct{}

// Analyze performs no operation at all
func (n NopAnalyzer) Analyze() {}
