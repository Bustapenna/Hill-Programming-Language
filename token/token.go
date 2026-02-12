package token

// -- TOKENS TYPES--

// Defines the token type and the token
type TokenType string

type Token struct {
	Type    TokenType
		Literal string
}

// Defines the types of tokens
const (
	// General
	ILLEGAL string = "ILLEGAL"
	EOF     string = "EOF"

	// Identifiers + literals
	IDENT string = "IDENT"
	INT   string = "INT"

	// Opreators
	ASSIGN_OPERATOR string = "="
	PLUS_OPERATOR   string = "+"
	MINUS_OPERATOR  string = "-"
	DIVIDE_OPERATOR string = "/"
	TIMES_OPERATROR string = "*"
	BANG            string = "!"
	LT              string = "<"
	GT              string = ">"
	EQ              string = "=="
	NOT_EQ          string = "!="
	LTEQ              string = "<="
	GTEQ              string = ">="

	// Delimiters
	COMMA     string = ","
	SEMICOLON string = ";"
	COLON     string = ":"

	LPAREN string = "("
	RPAREN string = ")"
	LBRACE string = "{"
	RBRACE string = "}"

	// Types
	INT_TYPE  string = "INT_TYPE"
	BOOL_TYPE string = "BOOL_TYPE"
	VOID_TYPE string = "VOID_TYPE"

	// Keywords
	FUNCTION        string = "FUNCTION"
	VAR             string = "VAR"
	IF              string = "IF"
	ELSE            string = "ELSE"
	RETURN          string = "RETURN"
	TRUE            string = "TRUE"
	FALSE           string = "FALSE"
)

// --TOKENS--

// Defines the keywords and the types
var keywords = map[string]TokenType{
	// Keywords
	"fnc":   TokenType(FUNCTION),
	"var":   TokenType(VAR),
	"ret":   TokenType(RETURN),
	"if":    TokenType(IF),
	"else":  TokenType(ELSE),
	"true":  TokenType(TRUE),
	"false": TokenType(FALSE),

	// Types
	"int":  TokenType(INT_TYPE),
	"bool": TokenType(BOOL_TYPE),
	"void": TokenType(VOID_TYPE),
}

// Tells if an ident is a token by returning its token type, otherwise, it returns the ident token type
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return TokenType(IDENT)
}
