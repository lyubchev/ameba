package token

type TokenType string

const (
	ZeroLiteral byte = 0x00
)

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// identifiers + literals
	IDENT   TokenType = "IDENT"  // add, foobar, x, y, ...
	INT     TokenType = "INT"    // 1343456
	STRING  TokenType = "STRING" // "foobar"
	BOOLEAN TokenType = "BOOL"   // true, false

	// operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

func New(t TokenType, l string) Token {
	return Token{Type: t, Literal: l}
}
