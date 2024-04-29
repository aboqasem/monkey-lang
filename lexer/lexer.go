package lexer

import (
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
		if l.peekChar() == '=' {
			lit := string(ch) + string(l.readChar())
			return newToken(token.EQ, lit)
		}
		return newToken(token.ASSIGN, ch)
	case '+':
		return newToken(token.PLUS, ch)
	case '-':
		return newToken(token.MINUS, ch)
	case '*':
		return newToken(token.AESTRIK, ch)
	case '/':
		return newToken(token.SLASH, ch)
	case '!':
		if l.peekChar() == '=' {
			lit := string(ch) + string(l.readChar())
			return newToken(token.NOT_EQ, lit)
		}
		return newToken(token.BANG, ch)
	case '<':
		return newToken(token.LT, ch)
	case '>':
		return newToken(token.GT, ch)
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
			return newToken(token.INT, num)
		}
	}

	return newToken(token.ILLEGAL, ch)
}

func newToken[L byte | string](tokType token.TokenType, lit L) token.Token {
	return token.Token{Type: tokType, Literal: string(lit)}
}

func (l *Lexer) readChar() (ch byte) {
	ch = l.peekChar()
	l.ch = ch
	l.currPosition = l.readPosition
	l.readPosition += 1
	return
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= l.inputLen {
		return 0
	} else {
		return l.input[l.readPosition]
	}
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
