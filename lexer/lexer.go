package lexer

import (
	"fmt"
	"monkey-lang/token"
)

type Lexer struct {
	input    string
	inputLen int
	// Current position in input (points to current char)
	currPosition int
	// Current reading position in input (after current char)
	readPosition int
	// Current char under examination
	ch byte
}

func New(input string) *Lexer {
	l := Lexer{input: input, inputLen: len(input)}
	l.readChar()
	return &l
}

func (l *Lexer) NextToken() token.Token {
	shouldConsumeChar := true
	defer func() {
		if shouldConsumeChar {
			l.readChar()
		}
	}()

	l.consumeWhitespace()

	ch := l.ch
	switch ch {
	case '=':
		return newToken(token.ASSIGN, ch)
	case '+':
		return newToken(token.PLUS, ch)
	case ',':
		return newToken(token.COMMA, ch)
	case ';':
		return newToken(token.SEMICOLON, ch)
	case '(':
		return newToken(token.LPAREN, ch)
	case ')':
		return newToken(token.RPAREN, ch)
	case '{':
		return newToken(token.LBRACE, ch)
	case '}':
		return newToken(token.RBRACE, ch)
	case 0:
		return newToken(token.EOF, "")
	default:
		if isLetter(ch) {
			shouldConsumeChar = false
			ident := l.readIdent()
			return newToken(token.LookupIdent(ident), ident)
		} else if isDigit(ch) {
			shouldConsumeChar = false
			num := l.readNumber()
			fmt.Println(num)
			return newToken(token.INT, num)
		}
	}

	return newToken(token.ILLEGAL, ch)
}

func newToken[L byte | string](tokType token.TokenType, lit L) token.Token {
	return token.Token{Type: tokType, Literal: string(lit)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= l.inputLen {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.currPosition = l.readPosition
	l.readPosition += 1
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) consumeWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdent() string {
	start := l.currPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.currPosition]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	start := l.currPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.currPosition]
}
