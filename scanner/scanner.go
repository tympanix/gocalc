package scanner

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"unicode"

	"github.com/tympanix/tymlang/scanner/token"
)

// Scanner is able to scan input files
type Scanner struct {
	r   *bufio.Reader
	buf bytes.Buffer
	i   int
}

// New creates a new scanner
func New(path string) (*Scanner, error) {

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return &Scanner{
		r: bufio.NewReader(f),
		i: 0,
	}, nil
}

func (s *Scanner) next() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return 0
	}
	s.buf.WriteRune(r)
	return r
}

func (s *Scanner) discard() {
	s.r.Discard(1)
}

func (s *Scanner) peek() rune {
	b, err := s.r.Peek(1)
	if err != nil {
		return 0
	}
	return rune(b[0])
}

func (s *Scanner) get() string {
	str := s.buf.String()
	s.buf.Reset()
	return str
}

func (s *Scanner) newToken(kind token.Kind) *token.Token {
	return token.New(kind, s.get())
}

// NextToken retrieves the next token from the scanner
func (s *Scanner) NextToken() *token.Token {

	for {
		for unicode.IsSpace(s.peek()) {
			s.discard()
		}

		r := s.next()

		if unicode.IsNumber(r) {
			for unicode.IsNumber(s.peek()) {
				r = s.next()
			}
			return s.newToken(token.INT)
		} else if r == '+' {
			return s.newToken(token.PLUS)
		} else if r == '*' {
			return s.newToken(token.MUL)
		} else if r == '/' {
			return s.newToken(token.DIV)
		} else if r == '(' {
			return s.newToken(token.LPAR)
		} else if r == ')' {
			return s.newToken(token.RPAR)
		} else if r == '^' {
			return s.newToken(token.POW)
		} else if r == 0 {
			return s.newToken(token.EOF)
		} else {
			log.Fatalf("Unknown token %c\n", r)
		}

	}

}
