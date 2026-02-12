package parser

import (
	"Hill/ast"
	"Hill/lexer"
	"Hill/token"

	"fmt"
)

type (
	prefixParseFnc func() ast.Expression
	infixParseFnc func(ast.Expression) ast.Expression
)

const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGRATER // < or >
	SUM // + 
	PRODUCT // *
	PREFIX // -X or !X
	CALL // myFunction(x)
)

// Define's what's a parser
type Parser struct {
	lex *lexer.Lexer

	errors []string

	currentToken token.Token
	peekToken token.Token

	prefixParseFncs map[token.TokenType]prefixParseFnc
	infixParseFncs map[token.TokenType]infixParseFnc
}

// Parser's helpers functions
func (p *Parser) registrerPrefix(tokenType token.TokenType, fnc prefixParseFnc) {
	p.prefixParseFncs[tokenType] = fnc
}

func (p *Parser) registrerInfix(tokenType token.TokenType, fnc infixParseFnc) {
	p.infixParseFncs[tokenType] = fnc
}

// Create's a new instance of the parser
func NewParser(l *lexer.Lexer) *Parser {
	var p *Parser = &Parser{lex: l, errors: []string{}}

	// Sets initial tokens
	p.nextToken()
	p.nextToken()

	p.prefixParseFncs = make(map[token.TokenType]prefixParseFnc)
	p.registrerPrefix(token.TokenType(token.IDENT), p.parseIdentifier)
	
	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

// It sets the tokens of the Parser to the next lexers token (helper function)
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

// Parse's the program
func (p *Parser) ParseProgram() *ast.Program {
	var program *ast.Program = &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.TokenType(token.EOF) {
		var stmt ast.Statement = p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

	p.nextToken()
	}

	return program
}

// Helper function that helps the "ParseProgram" function for parsing a statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.TokenType(token.VAR):
		return p.parseVarStatement()
	case token.TokenType(token.RETURN):
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Helper function that parses a var statement
func (p *Parser) parseVarStatement() *ast.VarStatement {
	var stmt *ast.VarStatement = &ast.VarStatement{Token: p.currentToken}

	if !p.peekTokenIsType() {
		p.typeError(p.peekToken)

		return nil
	}

	p.nextToken()
	stmt.Type = p.currentToken

	if !p.expectPeek(token.TokenType(token.IDENT)) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	
	if !p.expectPeek(token.TokenType(token.ASSIGN_OPERATOR)) {
		return nil
	}

	// TODO skipping expression until encounter a semicolon
	for !p.currentTokenIs(token.TokenType(token.SEMICOLON)) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	var statemnt *ast.ReturnStatement = &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	// TODO Skipping the expression until encountering a semicolon
	for !p.currentTokenIs(token.TokenType(token.SEMICOLON)) {
		p.nextToken()
	}

	return statemnt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	var stmt *ast.ExpressionStatement = &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if !p.expectPeek(token.TokenType(token.SEMICOLON)) {
		return nil
	}

	return stmt
}						

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFncs[p.currentToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

// Simple self-explenatory helper function
func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

// Siple self-explenatory helper function
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return  false
	}
}

// Helper function that helps to define if the current token is a type
func (p *Parser) currentTokenIsType() bool {
	t := p.currentToken.Type

	return t == token.TokenType(token.INT_TYPE) || t == token.TokenType(token.BOOL_TYPE)
}

// Helper function that helps to define if the peek token is a type
func (p *Parser) peekTokenIsType() bool {
	t := p.peekToken.Type

	return t == token.TokenType(token.INT_TYPE) || t == token.TokenType(token.BOOL_TYPE)
}

// Errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", t, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}

func (p *Parser) typeError(t token.Token) {
	msg := fmt.Sprintf("Expected a type, got %s ('%s') instead", t.Type, t.Literal)

	p.errors = append(p.errors, msg)
}

