package parser

import (
	"monkey-lang/ast"
	"monkey-lang/lexer"
	"monkey-lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	nextToken token.Token
}

func New(l *lexer.Lexer) (p *Parser) {
	p = &Parser{l: l}
	p.next()
	p.next()

	return
}

func (p *Parser) next() {
	p.currToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) Parse() (prog *ast.Program) {
	return
}
