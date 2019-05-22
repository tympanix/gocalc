package scanner

// Token reprensents a token from the parser
type Token int

const (
	EOF Token = iota
	IDENT
	PLUS
	MUL
	LPAR
	RPAR
)
