package parser

import (
	"Hill/ast"
	"Hill/lexer"
	"testing"
)

func TestVarStatements(t *testing.T) {
	var input string = `var int x = 5;
	var int y = 10;
	var int foobar = 3456345;`

	var lex *lexer.Lexer = lexer.NewLexer(input)
	var pars *Parser = NewParser(lex)

	var program *ast.Program = pars.ParseProgram()
	checkParserErrors(t, pars)

	// Checks basic errors
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain3 statements. Got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got = %q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got = %q", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not '%s'. Got '%s'", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("varStmt.Name.TokenLiteral() not '%s'. Got = '%s'", name, varStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}

	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	var input string = `ret 5;
	ret 10;
	ret 89898;`
	
	var l *lexer.Lexer = lexer.NewLexer(input)
	var p *Parser = NewParser(l)
	
	var program *ast.Program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements should be equal to 3 but isn't, got %d instead", len(program.Statements))
	}

	for _, statmts := range program.Statements {
		retStmt, ok := statmts.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Statement is not *ast.ReturnStatement. got = %T", retStmt)
			continue
		}

		if retStmt.TokenLiteral() != "ret" {
			t.Errorf("retStmt.TokenLiteral is not 'ret', got %q instead", retStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	var input string = "myVar;"

	var lex *lexer.Lexer = lexer.NewLexer(input)
	var p *Parser = NewParser(lex)
	var program *ast.Program = p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got %d instead.", len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ExpressionStatement. got %T instead", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got %T  instead.", stmt.Expression)
	}
	if ident.Value != "myVar" {
		t.Errorf("ident.Value is not %s. got %s instead.", "myVar", ident.Value)
	}
	if ident.TokenLiteral() != "myVar" {
		t.Errorf("ident.TokenLiteral is not %s. got %s instead", "myVar", ident.TokenLiteral())
	}
}
