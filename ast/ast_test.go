package ast

import (
	"monkey-lang/token"
	"testing"
)

func TestString(t *testing.T) {
	prog := &Program{[]Statement{
		&LetStatement{
			Token: token.Token{Type: token.LET, Literal: "let"},
			Name: &Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "variable"},
				Value: "variable",
			},
			Value: &Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "another_variable"},
				Value: "another_variable",
			},
		},
	}}

	expectedStr := `let variable = another_variable;
`

	if str := prog.String(); str != expectedStr {
		t.Errorf("program.String(): expected=%q, got=%q", expectedStr, str)
	}
}
