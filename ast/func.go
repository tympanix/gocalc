package ast

import (
	"log"
	"math"

	"github.com/tympanix/gocalc/debug"
)

type funcExp struct {
	name    string
	nparams int
	params  []Node
	fn      func(params []Node) float64
}

func (f *funcExp) Print() {
	debug.Println(f.name)
	debug.Indent()
	for v := range f.params {
		debug.Println(v)
	}
	debug.Outdent()
}

func (f *funcExp) Analyze() {
	if len(f.params) != f.nparams {
		log.Fatalf("Expected %d parameters in %s, got %d\n", f.nparams, f.name, len(f.params))
	}
}

// Calc returns the result of the function
func (f *funcExp) Calc() float64 {
	return f.fn(f.params)
}

// NewSqrtOp returns a new square root operator
func NewSqrtOp(params []Node) Node {
	return &funcExp{
		name:    "sqrt",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Sqrt(params[0].Calc())
		},
	}
}

// NewLog10Op returns a new AST node for log operations (base 10)
func NewLog10Op(params []Node) Node {
	return &funcExp{
		name:    "log",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Log10(params[0].Calc())
		},
	}
}

// NewLog2Op returns the AST node for log2 operations
func NewLog2Op(params []Node) Node {
	return &funcExp{
		name:    "log2",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Log2(params[0].Calc())
		},
	}
}

// NewPowFnOp returns the AST node for the pow function
func NewPowFnOp(params []Node) Node {
	return &funcExp{
		name:    "pow",
		nparams: 2,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Pow(params[0].Calc(), params[1].Calc())
		},
	}
}

// NewSinOp returns the AST node for the sin function
func NewSinOp(params []Node) Node {
	return &funcExp{
		name:    "sin",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Sin(params[0].Calc())
		},
	}
}

// NewCosOp returns the AST node for the cos function
func NewCosOp(params []Node) Node {
	return &funcExp{
		name:    "cos",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Cos(params[0].Calc())
		},
	}
}

// NewTanOp returns the AST node for the tan function
func NewTanOp(params []Node) Node {
	return &funcExp{
		name:    "tan",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Tan(params[0].Calc())
		},
	}
}

// NewAsinOp returns the AST node for the asin function
func NewAsinOp(params []Node) Node {
	return &funcExp{
		name:    "asin",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Asin(params[0].Calc())
		},
	}
}

// NewAcosOp returns the AST node for the acos function
func NewAcosOp(params []Node) Node {
	return &funcExp{
		name:    "acos",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Acos(params[0].Calc())
		},
	}
}

// NewAtanOp returns the AST node for the acos function
func NewAtanOp(params []Node) Node {
	return &funcExp{
		name:    "atan",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Atan(params[0].Calc())
		},
	}
}

// NewAbsOp returns the AST node for the abs function
func NewAbsOp(params []Node) Node {
	return &funcExp{
		name:    "abs",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Abs(params[0].Calc())
		},
	}
}

// NewLnOp returns the AST node for the abs function
func NewLnOp(params []Node) Node {
	return &funcExp{
		name:    "ln",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Log(params[0].Calc())
		},
	}
}
