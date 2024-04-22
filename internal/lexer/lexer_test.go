package lexer_test

import (
	"testing"

	"github.com/samallen659/ccJSONParser/internal/lexer"
	"github.com/samallen659/ccJSONParser/internal/token"
)

func TestLexer(t *testing.T) {
	input := `{"one":1}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "one"},
		{token.COLON, ":"},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.EOF, "EOF"},
	}

	l := lexer.New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"Unexpected tokenType returned. expected=%s got=%s",
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"Unexpected Literal returned. expected=%s got=%s",
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
