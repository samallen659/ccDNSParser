package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

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
