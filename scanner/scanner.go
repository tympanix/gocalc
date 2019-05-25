package scanner

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
		'.': token.DOT,
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
func NewFromReader(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

// NewFromString returns a new scanner from a string
func NewFromString(str string) *Scanner {
	return &Scanner{
		r: bufio.NewReader(strings.NewReader(str)),
	}
}

func (s *Scanner) next() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return 0
	}
	s.buf.WriteRune(r)
	return r
}

func (s *Scanner) nextN(n int) {
	for i := 0; i < n; i++ {
		s.next()
	}
}

func (s *Scanner) has(r rune) bool {
	if s.peekRune() == r {
		s.next()
		return true
	}
	return false
}

func (s *Scanner) clear() {
	s.buf.Reset()
}

func (s *Scanner) hasString(str string) bool {
	if s.peek(len(str)) == str {
		s.nextN(len(str))
		return true
	}
	return false
}

func (s *Scanner) hasDigit() bool {
	if unicode.IsNumber(s.peekRune()) {
		s.next()
		return true
	}
	return false
}

func (s *Scanner) hasLetter() bool {
	if unicode.IsLetter(s.peekRune()) {
		s.next()
		return true
	}
	return false
}

func (s *Scanner) discard() {
	s.r.Discard(1)
}

func (s *Scanner) rune() rune {
	return rune(s.buf.Bytes()[0])
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

		if s.has('0') {
			if s.has('.') {
				return s.scanIntegerToken()
			} else if s.has('x') {
				return s.scanHexToken()
			} else if s.has('b') {
				return s.scanBinToken()
			}
		} else if s.hasDigit() {
			s.scanInteger()
			s.has('.')
			return s.scanIntegerToken()
		} else if s.hasLetter() {
			for s.hasLetter() || s.hasDigit() {
				// noop
			}
			return s.newToken(token.IDENTIFIER)
		} else if s.hasString("//") {
			s.clear()
			for s.peekRune() != '\n' && s.peekRune() != 0 {
				s.discard()
			}
		} else if t, ok := symbols[s.peekRune()]; ok {
			s.next()
			return s.newToken(t)
		} else if s.has(0) {
			return s.newToken(token.EOF)
		} else {
			panic(fmt.Sprintf("unknown token: %c\n", s.next()))
		}
	}
}

func (s *Scanner) scanInteger() {
	for {
		if !s.hasDigit() {
			break
		}
	}
}

func (s *Scanner) scanIntegerToken() *token.Token {
	s.scanInteger()
	return s.newToken(token.DEC_LITERAL)
}

func (s *Scanner) scanHexToken() *token.Token {
	for {
		if s.hasDigit() {
			continue
		}
		if s.peekRune() >= 'a' && s.peekRune() <= 'f' {
			s.next()
			continue
		}
		if s.peekRune() >= 'A' && s.peekRune() <= 'F' {
			s.next()
			continue
		}
		return s.newToken(token.HEX_LITERAL)
	}
}

func (s *Scanner) scanBinToken() *token.Token {
	for s.has('0') || s.has('1') {
		// noop
	}
	return s.newToken(token.BIN_LITERAL)
}
