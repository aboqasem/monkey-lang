package parser

import (
	"monkey-lang/ast"
	"monkey-lang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	prog := p.Parse()
	checkParserErrors(t, p)

	if prog == nil {
		t.Fatal("Parse() returned nil")
	}
	if len := len(prog.Statements); len != 3 {
		t.Fatalf("statements: expected=3, got=%d", len)
	}

	tests := []struct{ expectedName string }{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		testLetStatements(t, stmt, tt.expectedName)
	}
}

func testLetStatements(t *testing.T, stmt ast.Statement, expectedName string) bool {
	if lit := stmt.TokenLiteral(); lit != "let" {
		t.Errorf(`stmt=%+v TokenLiteral(): expected="let", got=%s`, stmt, lit)
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf(`stmt=%+v type: expected=%T, got=%T`, stmt, &ast.LetStatement{}, letStmt)
		return false
	}

	if name := letStmt.Name.Value; name != expectedName {
		t.Errorf(`letStmt=%+v Name.Value: expected=%s, got=%s`, letStmt, expectedName, name)
		return false
	}

	if name := letStmt.Name.TokenLiteral(); name != expectedName {
		t.Errorf(`letStmt=%+v Name.TokenLiteral(): expected=%s, got=%s`, letStmt, expectedName, name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}
