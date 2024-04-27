package lexer

import (
	"testing"

	"github.com/samallen659/ccJSONParser/internal/token"
)

func TestLexer(t *testing.T) {
	input := `
     {
        "one":1,
        "two": {
            "three": 3
        },
        "four": [
            1,
            2,
            3,
            4
        ],
        "five": true,
        "six": false,
        "seven": null
    }`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "one"},
		{token.COLON, ":"},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.STRING, "two"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "three"},
		{token.COLON, ":"},
		{token.INT, "3"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "four"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.COMMA, ","},
		{token.INT, "4"},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.STRING, "five"},
		{token.COLON, ":"},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.STRING, "six"},
		{token.COLON, ":"},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.STRING, "seven"},
		{token.COLON, ":"},
		{token.NULL, "null"},
		{token.RBRACE, "}"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"Test(%d): Unexpected tokenType returned. expected=%s got=%s",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"Test(%d): Unexpected Literal returned. expected=%s got=%s",
				i,
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
