package scanner

// Token reprensents a token from the parser
type Token int

const (
	EOF Token = iota
	INT
	PLUS
	MUL
	LPAR
	RPAR
)
