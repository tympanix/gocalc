package scanner

import (
	"bufio"
	"bytes"
	"os"
	"unicode"
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

// NextToken retrieves the next token from the scanner
func (s *Scanner) NextToken() (Token, string) {

	for {
		for unicode.IsSpace(s.peek()) {
			s.discard()
		}

		r := s.next()

		if unicode.IsSpace(r) {
			continue
		} else if unicode.IsNumber(r) {
			for unicode.IsNumber(s.peek()) {
				r = s.next()
			}
			return IDENT, s.get()
		} else if r == '+' {
			return PLUS, s.get()
		} else if r == '*' {
			return MUL, s.get()
		} else if r == 0 {
			return EOF, ""
		}

	}

}
