package lexer

import (
	"github.com/impzero/ameba/token"
)

type Lexer struct {
	input           string
	currentPosition int  // current position in input (points to current char)
	readPosition    int  // current reading position in input (after current char)
	char            byte // current char under examination
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.readChar()

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok = newToken(token.EOF, token.ZeroLiteral)

	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier(l.currentPosition)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}
	return tok
}

func (l *Lexer) readIdentifier(startingPosition int) string {
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[startingPosition:l.currentPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.New(tokenType, string(ch))
}
