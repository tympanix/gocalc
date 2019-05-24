package parser

import (
	"log"
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
	}

	constants = map[string]constFactory{
		"pi": ast.NewPiOp,
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
		log.Fatalf("Expected token: %s\n", t.String())
	}
	return p.last()
}

func (p *Parser) last() *token.Token {
	return p.prev
}

// Parse parses the program
func (p *Parser) Parse() ast.Node {
	return p.parseExpression()
}

func (p *Parser) parseExpression() ast.Node {
	return p.parsePlus()
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
	if p.see(token.NUMBER) {
		return p.parseNumber()
	} else if p.have(token.MINUS) {
		return ast.NewNegOp(p.parseNumber())
	} else if p.have(token.LPAR) {
		exp := p.parseExpression()
		p.expect(token.RPAR)
		return exp
	} else if p.have(token.IDENTIFIER) {
		if p.see(token.LPAR) {
			return p.parseFunc()
		}
		return p.parseConstant()
	}
	log.Panicf("unexpected token: %s\n", p.current().String())
	return nil
}

func (p *Parser) parseNumber() ast.Node {
	p.expect(token.NUMBER)
	t := p.last()
	i, err := strconv.ParseFloat(t.String(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return ast.NewNumberLiteral(i)
}

func (p *Parser) parseConstant() ast.Node {
	t := p.last()
	if c, ok := constants[t.String()]; ok {
		return c()
	}
	log.Fatalf("unknown constant: %s\n", t.String())
	return nil
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
	log.Fatalf("Unknown function: %s\n", fn.String())
	return nil
}
