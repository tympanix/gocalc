package ast

import (
	"fmt"
	"math"

	"github.com/tympanix/gocalc/debug"
)

type funcExp struct {
	name    string
	nparams int
	params  []Node
	fn      func(params []Node) float64
	FloatType
}

func (f *funcExp) Print() {
	debug.Println(f.name)
	debug.Indent()
	for v := range f.params {
		debug.Println(v)
	}
	debug.Outdent()
}

func (f *funcExp) Analyze() error {
	if len(f.params) != f.nparams {
		return fmt.Errorf("expected %d parameters in %s, got %d", f.nparams, f.name, len(f.params))
	}
	return nil
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

// NewDegOp returns the AST node for the deg function
func NewDegOp(params []Node) Node {
	return &funcExp{
		name:    "deg",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return params[0].Calc() * 180 / math.Pi
		},
	}
}

// NewRadOp returns the AST node for the rad function
func NewRadOp(params []Node) Node {
	return &funcExp{
		name:    "rad",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return params[0].Calc() * math.Pi / 180
		},
	}
}

// NewRoundOp returns the AST node for the round function
func NewRoundOp(params []Node) Node {
	return &funcExp{
		name:    "round",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Round(params[0].Calc())
		},
	}
}

// NewFloorOp returns the AST node for the round function
func NewFloorOp(params []Node) Node {
	return &funcExp{
		name:    "floor",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Floor(params[0].Calc())
		},
	}
}

// NewCeilOp returns the AST node for the round function
func NewCeilOp(params []Node) Node {
	return &funcExp{
		name:    "ceil",
		nparams: 1,
		params:  params,
		fn: func(params []Node) float64 {
			return math.Ceil(params[0].Calc())
		},
	}
}
