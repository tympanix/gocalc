package ast

import "github.com/tympanix/gocalc/debug"

type unaryExp struct {
	name  string
	param Node
	fn    func(float64) float64
	NopAnalyzer
}

func (u *unaryExp) Print() {
	debug.Println(u.name)
	debug.Indent()
	debug.Println(u.param)
	debug.Outdent()
}

func (u *unaryExp) Calc() float64 {
	return u.fn(u.param.Calc())
}

func (u *unaryExp) Type() Type {
	return u.param.Type()
}

// NewNegOp returns the AST node for unary negation operator
func NewNegOp(param Node) Node {
	return &unaryExp{
		name:  "-",
		param: param,
		fn: func(a float64) float64 {
			return -a
		},
	}
}
