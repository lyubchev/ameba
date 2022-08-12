package token

type TokenType string

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
	RRAPEN TokenType = ")"
	LBRACE TokenType = "{"
	RBRAE  TokenType = "}"

	// keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}
