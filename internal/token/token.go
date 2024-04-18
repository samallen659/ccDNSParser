package token

type TokenType string

const (
	ILLEGAL = "illegal"
	EOF     = "eof"

	INT    = "INT"
	STRING = "STRING"

	COLON = ":"
	COMMA = ","

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	TRUE  = "true"
	FALSE = "false"

	MINUS = "-"
)

type Token struct {
	Type    TokenType
	Literal string
}
