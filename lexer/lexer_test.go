package lexer

import (
	"testing"

	"Hill/token"
)

func TestNextToken(t *testing.T) {
	input := `var int five = 5;
	var int ten = 10;

	fnc int Add(x int, y int) 
	{
		ret x+y;

		if (x <= y) 
		{
			ret true;
		} else
		{
			ret false;
		}
	}

	var result: int = Add(5, 2);`

	

	tests := []struct {
		expectedType   token.TokenType
		expetedLiteral string
	}{
		{token.TokenType(token.VAR), "var"},
		{token.TokenType(token.INT_TYPE), "int"},
		{token.TokenType(token.IDENT), "five"},
		{token.TokenType(token.ASSIGN_OPERATOR), "="},
		{token.TokenType(token.INT), "5"},
		{token.TokenType(token.SEMICOLON), ";"},
		{token.TokenType(token.VAR), "var"},
		{token.TokenType(token.INT_TYPE), "int"},
		{token.TokenType(token.IDENT), "ten"},
		{token.TokenType(token.ASSIGN_OPERATOR), "="},
		{token.TokenType(token.INT), "10"},
		{token.TokenType(token.SEMICOLON), ";"},
		{token.TokenType(token.FUNCTION), "fnc"},
		{token.TokenType(token.INT_TYPE), "int"},
		{token.TokenType(token.IDENT), "Add"},
		{token.TokenType(token.LPAREN), "("},
		{token.TokenType(token.IDENT), "x"},
		{token.TokenType(token.INT_TYPE), "int"},
		{token.TokenType(token.COMMA), ","},
		{token.TokenType(token.IDENT), "y"},
		{token.TokenType(token.INT_TYPE), "int"},
		{token.TokenType(token.RPAREN), ")"},
		{token.TokenType(token.LBRACE), "{"},
		{token.TokenType(token.RETURN), "ret"},
		{token.TokenType(token.IDENT), "x"},
		{token.TokenType(token.PLUS_OPERATOR), "+"},
		{token.TokenType(token.IDENT), "y"},
		{token.TokenType(token.SEMICOLON), ";"},
		{token.TokenType(token.IF), "if"},
		{token.TokenType(token.LPAREN), "("},
		{token.TokenType(token.IDENT), "x"},
		{token.TokenType(token.LTEQ), "<="},
		{token.TokenType(token.IDENT), "y"},
		{token.TokenType(token.RPAREN), ")"},
		{token.TokenType(token.LBRACE), "{"},
		{token.TokenType(token.RETURN), "ret"},
		{token.TokenType(token.TRUE), "true"},
		{token.TokenType(token.SEMICOLON), ";"},
		{token.TokenType(token.RBRACE), "}"},
		{token.TokenType(token.ELSE), "else"},
		{token.TokenType(token.LBRACE), "{"},
		{token.TokenType(token.RETURN), "ret"},
		{token.TokenType(token.FALSE), "false"},
		{token.TokenType(token.SEMICOLON), ";"},
		{token.TokenType(token.RBRACE), "}"},
		{token.TokenType(token.RBRACE), "}"},
		{token.TokenType(token.VAR), "var"},
		{token.TokenType(token.IDENT), "result"},
		{token.TokenType(token.COLON), ":"},
		{token.TokenType(token.INT_TYPE), "int"},
		{token.TokenType(token.ASSIGN_OPERATOR), "="},
		{token.TokenType(token.IDENT), "Add"},
		{token.TokenType(token.LPAREN), "("},
		{token.TokenType(token.INT), "5"},
		{token.TokenType(token.COMMA), ","},
		{token.TokenType(token.INT), "2"},
		{token.TokenType(token.RPAREN), ")"},
		{token.TokenType(token.SEMICOLON), ";"},
		{token.TokenType(token.EOF), ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Exptected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expetedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected = %q, got = %q", i, tt.expetedLiteral, tok.Literal)
		}
	}
}

