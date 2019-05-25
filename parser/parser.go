package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tympanix/gocalc/ast"
	"github.com/tympanix/gocalc/scanner"
	"github.com/tympanix/gocalc/scanner/token"
)

type funcExpFactory func(params []ast.Node) ast.Node
type constFactory func() ast.Node

var (
	functions = map[string]funcExpFactory{
		"sqrt":  ast.NewSqrtOp,
		"log":   ast.NewLog10Op,
		"log10": ast.NewLog10Op,
		"log2":  ast.NewLog2Op,
		"pow":   ast.NewPowFnOp,
		"sin":   ast.NewSinOp,
		"cos":   ast.NewCosOp,
		"tan":   ast.NewTanOp,
		"asin":  ast.NewAsinOp,
		"acos":  ast.NewAcosOp,
		"atan":  ast.NewAtanOp,
		"ln":    ast.NewLnOp,
		"abs":   ast.NewAbsOp,
		"rad":   ast.NewRadOp,
		"deg":   ast.NewDegOp,
		"round": ast.NewRoundOp,
		"floor": ast.NewFloorOp,
		"ceil":  ast.NewCeilOp,
	}

	constants = map[string]constFactory{
		"pi": ast.NewPiOp,
		"π":  ast.NewPiOp,
		"e":  ast.NewEulerOp,
	}
)

// Parser parses the input program from a scanner
type Parser struct {
	s      *scanner.Scanner
	tokens []*token.Token
	prev   *token.Token
	i      int
}

// New return a new parser
func New(s *scanner.Scanner) *Parser {
	return &Parser{s: s}
}

func (p *Parser) pump(n int) {
	for len(p.tokens) < n {
		next := p.s.NextToken()
		p.tokens = append(p.tokens, next)
	}
}

func (p *Parser) current() *token.Token {
	p.pump(1)
	return p.tokens[0]
}

func (p *Parser) pop() {
	if len(p.tokens) > 0 {
		p.prev = p.tokens[0]
		p.tokens = p.tokens[1:]
	}
	p.pump(1)
}

func (p *Parser) have(t token.Kind) bool {
	p.pump(1)
	e := p.current()

	if e.Kind() == t {
		p.pop()
	}

	return e.Kind() == t
}

func (p *Parser) see(t token.Kind) bool {
	return p.current().Kind() == t
}

func (p *Parser) expect(t token.Kind) *token.Token {
	if !p.have(t) {
		panic(fmt.Sprintf("expected token: %s\n", t.String()))
	}
	return p.last()
}

func (p *Parser) last() *token.Token {
	return p.prev
}

// Parse parses the program
func (p *Parser) Parse() (exp ast.Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()
	exp = p.parseExpression()
	p.expect(token.EOF)
	return exp, nil
}

func (p *Parser) parseExpression() ast.Node {
	return p.parseBitwiseOr()
}

func (p *Parser) parseBitwiseOr() ast.Node {
	lhs := p.parseBitwiseXor()

	for {
		if p.have(token.OR) {
			lhs = ast.NewBitwiseOrOp(lhs, p.parseBitwiseXor())
		} else {
			break
		}
	}
	return lhs
}

func (p *Parser) parseBitwiseXor() ast.Node {
	lhs := p.parseBitwiseAnd()

	for {
		if p.have(token.XOR) {
			lhs = ast.NewBitwiseXorOp(lhs, p.parseBitwiseAnd())
		} else {
			break
		}
	}
	return lhs
}

func (p *Parser) parseBitwiseAnd() ast.Node {
	lhs := p.parsePlus()

	for {
		if p.have(token.AND) {
			lhs = ast.NewBitwiseAndOp(lhs, p.parsePlus())
		} else {
			break
		}
	}
	return lhs
}

func (p *Parser) parsePlus() ast.Node {
	lhs := p.parseMul()

	for {
		if p.have(token.PLUS) {
			lhs = ast.NewPlusOp(lhs, p.parseMul())
		} else if p.have(token.MINUS) {
			lhs = ast.NewMinusOp(lhs, p.parseMul())
		} else {
			break
		}
	}
	return lhs
}

func (p *Parser) parseMul() ast.Node {
	lhs := p.parsePow()

	for {
		if p.have(token.MUL) {
			lhs = ast.NewMulOp(lhs, p.parsePow())
		} else if p.have(token.DIV) {
			lhs = ast.NewDivOp(lhs, p.parsePow())
		} else if p.have(token.MOD) {
			lhs = ast.NewModOp(lhs, p.parsePow())
		} else if p.see(token.IDENT) || p.see(token.LPAR) {
			// implicit multiplication
			lhs = ast.NewMulOp(lhs, p.parseAtomic())
		} else {
			break
		}
	}
	return lhs
}

func (p *Parser) parsePow() ast.Node {
	lhs := p.parseAtomic()

	for p.have(token.POW) {
		lhs = ast.NewPowOp(lhs, p.parseAtomic())
	}
	return lhs
}

func (p *Parser) parseAtomic() ast.Node {
	if p.have(token.MINUS) {
		return ast.NewNegOp(p.parseAtomic())
	} else if p.have(token.LPAR) {
		exp := p.parseExpression()
		p.expect(token.RPAR)
		return exp
	} else if p.have(token.IDENT) {
		if p.see(token.LPAR) {
			return p.parseFunc()
		}
		return p.parseConstant()
	} else {
		return p.parseNumber()
	}
}

func (p *Parser) parseNumber() ast.Node {
	if p.have(token.FLOAT_LITERAL) {
		t := p.last()
		i, err := strconv.ParseFloat(t.String(), 64)
		if err != nil {
			panic(err)
		}
		return ast.NewFloatLiteral(i)
	} else if p.have(token.INT_LITERAL) {
		t := p.last()
		i, err := strconv.ParseUint(t.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return ast.NewIntegerLiteral(float64(i))
	} else if p.have(token.HEX_LITERAL) {
		t := p.last()
		i, err := strconv.ParseUint(t.String()[2:], 16, 64)
		if err != nil {
			panic(err)
		}
		return ast.NewIntegerLiteral(float64(i))
	} else if p.have(token.BIN_LITERAL) {
		t := p.last()
		i, err := strconv.ParseUint(t.String()[2:], 2, 64)
		if err != nil {
			panic(err)
		}
		return ast.NewIntegerLiteral(float64(i))
	}
	panic(fmt.Sprintf("unexpected token: %s\n", p.current().Kind().String()))
}

func (p *Parser) parseConstant() ast.Node {
	t := p.last()
	if c, ok := constants[t.String()]; ok {
		return c()
	}
	panic(fmt.Sprintf("undefined constant: %s\n", t.String()))
}

func (p *Parser) parseFunc() ast.Node {
	fn := p.last()

	var params []ast.Node
	p.expect(token.LPAR)
	for {
		exp := p.parseExpression()
		if exp != nil {
			params = append(params, exp)
		}
		if !p.have(token.COMMA) {
			break
		}
	}
	p.expect(token.RPAR)
	if f, ok := functions[fn.String()]; ok {
		return f(params)
	}
	panic(fmt.Sprintf("undefined function: %s\n", fn.String()))
}
