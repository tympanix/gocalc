package parser

import (
	"log"
	"strconv"

	"github.com/tympanix/gocalc/ast"
	"github.com/tympanix/gocalc/scanner"
	"github.com/tympanix/gocalc/scanner/token"
)

// Parser parses the input program from a scanner
type Parser struct {
	s      *scanner.Scanner
	tokens []*token.Token
	last   *token.Token
}

// New return a new parser
func New(s *scanner.Scanner) *Parser {
	return &Parser{s: s}
}

func (p *Parser) have(t token.Kind) bool {
	if len(p.tokens) == 0 {
		next := p.s.NextToken()
		p.tokens = append(p.tokens, next)
		p.last = next
	}

	e := p.tokens[0]

	if e.Kind() == t {
		p.tokens = p.tokens[1:]
	}

	return e.Kind() == t

}

func (p *Parser) expect(t token.Kind) *token.Token {
	if !p.have(t) {
		log.Fatal("Unexpected token")
	}
	return p.get()
}

func (p *Parser) get() *token.Token {
	return p.last
}

// ParseProgram parses the program
func (p *Parser) ParseProgram() ast.Node {
	return p.parseExpression()
}

func (p *Parser) parseExpression() ast.Node {
	return p.parsePlus()
}

func (p *Parser) parsePlus() ast.Node {
	lhs := p.parseMul()

	for p.have(token.PLUS) {
		lhs = ast.PlusOp{
			Lhs: lhs,
			Rhs: p.parseMul(),
		}
	}
	return lhs
}

func (p *Parser) parseMul() ast.Node {
	lhs := p.parsePow()

	for {
		if p.have(token.MUL) {
			lhs = ast.MulOp{
				Lhs: lhs,
				Rhs: p.parsePow(),
			}
		} else if p.have(token.DIV) {
			lhs = ast.DivOp{
				Lhs: lhs,
				Rhs: p.parsePow(),
			}
		} else {
			break
		}
	}
	return lhs
}

func (p *Parser) parsePow() ast.Node {
	lhs := p.parseInteger()

	for p.have(token.POW) {
		lhs = ast.PowOp{
			Lhs: lhs,
			Rhs: p.parseInteger(),
		}
	}
	return lhs
}

func (p *Parser) parseInteger() ast.Node {
	if p.have(token.INT) {
		t := p.get()
		i, err := strconv.Atoi(t.String())
		if err != nil {
			log.Fatal(err)
		}
		return ast.IntLiteral(i)
	} else if p.have(token.LPAR) {
		exp := p.parseExpression()
		p.expect(token.RPAR)
		return exp
	}
	log.Fatal("Unexpected token")
	return nil
}
