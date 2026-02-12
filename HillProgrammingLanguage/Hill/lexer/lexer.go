package lexer

import "Hill/token"

type Lexer struct {
	input        string
	position     int  // Current position in input (current character)
	readPosition int  // Current reading position in input (after current character)
	ch           byte // Current character being processed
}

// Creates a new lexer
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Readers

// Reads the identifier (it's also a helper function)
func (l *Lexer) ReadIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	var position int = l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Reads the next character and updates the "ch" of the current lexer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// Reads the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespaces()

	// Switch part
	switch l.ch {
	// One letters chars
	case '=':
		if l.peekChar() == '=' {
			var ch byte = l.ch
			l.readChar()
			var literal string = string(ch) + string(l.ch)
			tok = token.Token{Type: token.TokenType(token.EQ), Literal: literal}
		} else {
			tok = newToken(token.TokenType(token.ASSIGN_OPERATOR), l.ch)
		}
	case ';':
		tok = newToken(token.TokenType(token.SEMICOLON), l.ch)
	case '(':
		tok = newToken(token.TokenType(token.LPAREN), l.ch)
	case ')':
		tok = newToken(token.TokenType(token.RPAREN), l.ch)
	case ',':
		tok = newToken(token.TokenType(token.COMMA), l.ch)
	case '+':
		tok = newToken(token.TokenType(token.PLUS_OPERATOR), l.ch)
	case '{':
		tok = newToken(token.TokenType(token.LBRACE), l.ch)
	case '}':
		tok = newToken(token.TokenType(token.RBRACE), l.ch)
	case '-':
		tok = newToken(token.TokenType(token.MINUS_OPERATOR), l.ch)
	case '*':
		tok = newToken(token.TokenType(token.TIMES_OPERATROR), l.ch)
	case '/':
		tok = newToken(token.TokenType(token.DIVIDE_OPERATOR), l.ch)
	case ':':
		tok = newToken(token.TokenType(token.COLON), l.ch)
	case '!':
		if l.peekChar() == '=' {
			var ch byte = l.ch
			l.readChar()
			var literal string = string(ch) + string(l.ch)
			tok = token.Token{Type: token.TokenType(token.NOT_EQ), Literal: literal}
		} else {
			tok = newToken(token.TokenType(token.BANG), l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			var ch byte = l.ch
			l.readChar()
			var literal string = string(ch) + string(l.ch)
			tok = token.Token{Type: token.TokenType(token.LTEQ), Literal: literal}
		} else {
			tok = newToken(token.TokenType(token.LT), l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			var ch byte = l.ch
			l.readChar()
			var literal string = string(ch) + string(l.ch)
			tok = token.Token{Type: token.TokenType(token.GTEQ), Literal: literal}
		} else {
			tok = newToken(token.TokenType(token.GT), l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.TokenType(token.EOF)

	// Finds multiletters ones
	default:
		if isLetter(l.ch) {
			tok.Literal = l.ReadIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.TokenType(token.INT)
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.TokenType(token.ILLEGAL), l.ch)
		}
	}

	l.readChar()
	return tok
}

// Helper functions

// Helper function that creates a token following the given data
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Helper function that helps define if a charcter is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Helper function that simply help to define if the given "ch" is a number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Helper function that helps to find to characters tokens composed by existing tokens by looking at the next character
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Helper functions that simply skips whitespaces
func (l *Lexer) eatWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
