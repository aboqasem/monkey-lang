package ast

import (
	"bytes"
	"monkey-lang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() (lit string) {
	if len(p.Statements) > 0 {
		lit = p.Statements[0].TokenLiteral()
	}

	return
}

func (p *Program) String() string {
	out := bytes.Buffer{}
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
		out.WriteRune('\n')
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode()       {}
func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LetStatement) String() string {
	out := bytes.Buffer{}
	out.WriteString(s.TokenLiteral())
	out.WriteRune(' ')
	out.WriteString(s.Name.String())
	out.WriteString(" = ")
    // TODO: should not be nil, remove when fully implemented
	if v := s.Value; v != nil {
		out.WriteString(v.String())
	}
	out.WriteRune(';')

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (s *ReturnStatement) statementNode()       {}
func (s *ReturnStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ReturnStatement) String() string {
	out := bytes.Buffer{}
	out.WriteString(s.TokenLiteral())
	out.WriteRune(' ')
    // TODO: should not be nil, remove when fully implemented
	if v := s.Value; v != nil {
		out.WriteString(v.String())
	}
	out.WriteRune(';')

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
