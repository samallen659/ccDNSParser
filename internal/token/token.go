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
	NULL  = "null"

	MINUS = "-"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func LookupKeywords(s string) TokenType {
	if tokenType, ok := keywords[s]; ok {
		return tokenType
	} else {
		return ILLEGAL
	}
}
