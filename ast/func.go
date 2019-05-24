package ast

import (
	"math"

	"github.com/tympanix/gocalc/debug"
)

func newFuncExp(name string, params []Node) *funcExp {
	return &funcExp{
		name:   name,
		params: params,
	}
}

type funcExp struct {
	name   string
	params []Node
}

func (f *funcExp) Print() {
	debug.Println(f.name)
	debug.Indent()
	for v := range f.params {
		debug.Println(v)
	}
	debug.Outdent()
}

// NewSqrtOp returns a new square root operator
func NewSqrtOp(params []Node) Node {
	return &SqrtOp{newFuncExp("sqrt", params)}
}

// SqrtOp is the AST node for the square root operator
type SqrtOp struct {
	*funcExp
}

// Calc returns the square root of the argument
func (s *SqrtOp) Calc() float64 {
	return math.Sqrt(s.params[0].Calc())
}

// NewLog10Op returns a new AST node for log operations (base 10)
func NewLog10Op(params []Node) Node {
	return &Log10Op{newFuncExp("log", params)}
}

// Log10Op calculates log (base 10)
type Log10Op struct {
	*funcExp
}

// Calc performs log10 on the parameter
func (l *Log10Op) Calc() float64 {
	return math.Log10(l.params[0].Calc())
}

// NewLog2Op returns the AST node for log2 operations
func NewLog2Op(params []Node) Node {
	return &Log2Op{newFuncExp("log2", params)}
}

// Log2Op is the structure of a log2 AST node
type Log2Op struct {
	*funcExp
}

// Calc performs log2 on the paramter
func (l *Log2Op) Calc() float64 {
	return math.Log2(l.params[0].Calc())
}
