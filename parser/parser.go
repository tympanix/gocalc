package parser

import (
	"log"
	"strconv"

	"github.com/tympanix/tymlang/ast"
	"github.com/tympanix/tymlang/scanner"
)

// Parser parses the input program from a scanner
type Parser struct {
	s      *scanner.Scanner
	tokens []tokenEntry
	last   tokenEntry
}

type tokenEntry struct {
	token scanner.Token
	str   string
}

// New return a new parser
func New(s *scanner.Scanner) *Parser {
	return &Parser{s: s}
}

func (p *Parser) have(t scanner.Token) bool {
	if len(p.tokens) == 0 {
		next, str := p.s.NextToken()
		e := tokenEntry{
			token: next,
			str:   str,
		}
		p.tokens = append(p.tokens, e)

		p.last = e
	}

	e := p.tokens[0]

	if e.token == t {
		p.tokens = p.tokens[1:]
	}

	return e.token == t

}

func (p *Parser) expect(t scanner.Token) (scanner.Token, string) {
	if !p.have(t) {
		log.Fatal("Unexpected token")
	}
	return p.get()
}

func (p *Parser) get() (scanner.Token, string) {
	e := p.last
	return e.token, e.str
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

	for p.have(scanner.PLUS) {
		lhs = ast.PlusOp{
			Lhs: lhs,
			Rhs: p.parseMul(),
		}
	}
	return lhs
}

func (p *Parser) parseMul() ast.Node {
	lhs := p.parsePow()

	for p.have(scanner.MUL) {
		lhs = ast.MulOp{
			Lhs: lhs,
			Rhs: p.parsePow(),
		}
	}
	return lhs
}

func (p *Parser) parsePow() ast.Node {
	lhs := p.parseInteger()

	for p.have(scanner.POW) {
		lhs = ast.PowOp{
			Lhs: lhs,
			Rhs: p.parseInteger(),
		}
	}
	return lhs
}

func (p *Parser) parseInteger() ast.Node {
	if p.have(scanner.INT) {
		_, str := p.get()
		i, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		return ast.IntLiteral(i)
	} else if p.have(scanner.LPAR) {
		exp := p.parseExpression()
		p.expect(scanner.RPAR)
		return exp
	}
	log.Fatal("Unexpected token")
	return nil
}
