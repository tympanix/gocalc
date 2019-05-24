package scanner

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/tympanix/gocalc/scanner/token"
)

var (
	symbols = map[rune]token.Kind{
		'+': token.PLUS,
		'-': token.MINUS,
		'*': token.MUL,
		'/': token.DIV,
		'^': token.POW,
		'(': token.LPAR,
		')': token.RPAR,
		',': token.COMMA,
	}
)

// Scanner is able to scan input files
type Scanner struct {
	r   *bufio.Reader
	buf bytes.Buffer
	i   int
}

// NewFromFile creates a new scanner from a file path
func NewFromFile(path string) (*Scanner, error) {

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return &Scanner{
		r: bufio.NewReader(f),
	}, nil
}

// NewFromReader returns a new scanner from a io.Reader object
func NewFromReader(r io.Reader) (*Scanner, error) {
	return &Scanner{
		r: bufio.NewReader(r),
	}, nil
}

// NewFromString returns a new scanner from a string
func NewFromString(str string) (*Scanner, error) {
	return &Scanner{
		r: bufio.NewReader(strings.NewReader(str)),
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

func (s *Scanner) peekRune() rune {
	b, err := s.r.Peek(1)
	if err != nil {
		return 0
	}
	return rune(b[0])
}

func (s *Scanner) peek(n int) string {
	str, err := s.r.Peek(n)
	if err != nil {
		return ""
	}
	return string(str)
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
		for unicode.IsSpace(s.peekRune()) {
			s.discard()
		}

		r := s.next()

		if unicode.IsNumber(r) {
			for unicode.IsNumber(s.peekRune()) {
				r = s.next()
			}
			return s.newToken(token.NUMBER)
		} else if unicode.IsLetter(r) {
			for unicode.IsNumber(s.peekRune()) || unicode.IsLetter(s.peekRune()) {
				r = s.next()
			}
			return s.newToken(token.IDENTIFIER)
		} else if r == '/' && s.peekRune() == '/' {
			for s.peekRune() != '\n' || s.peekRune() == 0 {
				s.next()
			}
		} else if t, ok := symbols[r]; ok {
			return s.newToken(t)
		} else if r == 0 {
			return s.newToken(token.EOF)
		} else {
			log.Fatalf("Unknown token %c\n", r)
		}
	}
}
