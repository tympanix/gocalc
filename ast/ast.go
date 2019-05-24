package ast

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
