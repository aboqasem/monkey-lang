package parser

import (
	"fmt"
	"monkey-lang/ast"
	"monkey-lang/lexer"
	"monkey-lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	nextToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) (p *Parser) {
	p = &Parser{l: l, errors: []string{}}
	p.next()
	p.next()

	return
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token %q, got %q", t, p.nextToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) next() {
	p.currToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() (stmt *ast.LetStatement) {
	stmt = &ast.LetStatement{Token: p.currToken}

	if !p.expectNext(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectNext(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currTokenIs(token.SEMICOLON) {
		p.next()
	}

	return
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) nextTokenIs(t token.TokenType) bool {
	return p.nextToken.Type == t
}

func (p *Parser) expectNext(t token.TokenType) bool {
	if !p.nextTokenIs(t) {
		p.peekError(t)
		return false
	}
	p.next()
	return true
}

func (p *Parser) Parse() (prog *ast.Program) {
	prog = &ast.Program{}
	prog.Statements = []ast.Statement{}

	for !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.next()
	}
	return
}
