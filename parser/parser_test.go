package parser

import (
	"monkey-lang/ast"
	"monkey-lang/lexer"
	"testing"
)

func parseProgram(t *testing.T, input string) (prog *ast.Program) {
	l := lexer.New(input)
	p := New(l)

	prog = p.Parse()
	if prog == nil {
		t.Fatal("Parse() returned nil")
	}
	checkParserErrors(t, p)

	return
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	prog := parseProgram(t, input)

	if expectedLen, len := 3, len(prog.Statements); len != expectedLen {
		t.Fatalf("statements: expected=%d, got=%d", expectedLen, len)
	}

	tests := []struct{ expectedName string }{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		testLetStatement(t, stmt, tt.expectedName)
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, expectedName string) bool {
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

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 838383;
`

	prog := parseProgram(t, input)

	if expectedLen, len := 3, len(prog.Statements); len != expectedLen {
		t.Fatalf("statements: expected=%d, got=%d", expectedLen, len)
	}

	for _, stmt := range prog.Statements {
		testReturnStatement(t, stmt)
	}
}

func testReturnStatement(t *testing.T, stmt ast.Statement) bool {
	if expectedLit, lit := "return", stmt.TokenLiteral(); lit != expectedLit {
		t.Errorf(`stmt=%+v TokenLiteral(): expected=%s, got=%s`, stmt, expectedLit, lit)
		return false
	}

	returnStmt, ok := stmt.(*ast.ReturnStatement)
	if !ok {
		t.Errorf(`stmt=%+v type: expected=%T, got=%T`, stmt, &ast.ReturnStatement{}, returnStmt)
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
