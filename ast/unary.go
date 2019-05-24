package ast

import "github.com/tympanix/gocalc/debug"

type unaryExp struct {
	name  string
	param Node
	fn    func(Node) float64
	NopAnalyzer
}

func (u *unaryExp) Print() {
	debug.Println(u.name)
	debug.Indent()
	debug.Println(u.param)
	debug.Outdent()
}

func (u *unaryExp) Calc() float64 {
	return u.fn(u.param)
}

// NewNegOp returns the AST node for unary negation operator
func NewNegOp(param Node) Node {
	return &unaryExp{
		name:  "-",
		param: param,
		fn: func(param Node) float64 {
			return -param.Calc()
		},
	}
}
