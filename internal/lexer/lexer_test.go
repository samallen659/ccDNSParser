package lexer

import (
	"testing"

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

	l := New(input)

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

func TestIsLetter(t *testing.T) {
	t.Run("lower case letter", func(t *testing.T) {
		if !isLetter('b') {
			t.Fatal("isLetter returned false for b. Expected=true")
		}
	})
	t.Run("upper case letter", func(t *testing.T) {
		if !isLetter('N') {
			t.Fatal("isLetter returned false for N. Expected=true")
		}
	})
	t.Run("int", func(t *testing.T) {
		if isLetter(10) {
			t.Fatal("isLetter return true for 10. Expected=false")
		}
	})
}
