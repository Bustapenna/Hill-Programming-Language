package ast

import (
	"Hill/token"
	"testing"
)

func TestString(t *testing.T) {
	var program *Program = &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.TokenType(token.VAR), Literal: "var"},
				Type: token.Token{Type: token.TokenType(token.INT_TYPE), Literal: "int "},
				Name: &Identifier{
					Token: token.Token{Type: token.TokenType(token.IDENT), Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.TokenType(token.IDENT), Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "var int myVar = anotherVar;" {
		t.Errorf("program.String() is wrong. got %q instead", program.String())
	} 
}
