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
