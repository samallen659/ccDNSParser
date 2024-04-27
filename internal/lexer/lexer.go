package lexer

import (
	"bytes"

	"github.com/samallen659/ccJSONParser/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, []byte{l.ch})
	case '(':
		tok = newToken(token.LPAREN, []byte{l.ch})
	case ')':
		tok = newToken(token.RPAREN, []byte{l.ch})
	case '{':
		tok = newToken(token.LBRACE, []byte{l.ch})
	case '}':
		tok = newToken(token.RBRACE, []byte{l.ch})
	case '[':
		tok = newToken(token.LBRACKET, []byte{l.ch})
	case ']':
		tok = newToken(token.RBRACKET, []byte{l.ch})
	case '-':
		tok = newToken(token.MINUS, []byte{l.ch})
	case ',':
		tok = newToken(token.COMMA, []byte{l.ch})
	case '"':
		tok = newToken(token.STRING, l.readString())
	case 0:
		tok = token.Token{Type: token.EOF, Literal: "EOF"}
	default:
		if isDigit(l.ch) {
			tok = newToken(token.INT, l.readNumber())
		} else if isLetter(l.ch) {
			keyword := l.readKeyword()
			tokenType := token.LookupKeywords(keyword)
			switch tokenType {
			case token.TRUE:
				tok = token.Token{Type: tokenType, Literal: "true"}
			case token.FALSE:
				tok = token.Token{Type: tokenType, Literal: "false"}
			case token.NULL:
				tok = token.Token{Type: tokenType, Literal: "null"}
			case token.ILLEGAL:
				tok = token.Token{Type: token.ILLEGAL, Literal: "ILLEGAL"}
			}
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: "ILLEGAL"}
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() []byte {
	var out bytes.Buffer

	for l.peekChar() != '"' {
		l.readChar()
		out.WriteByte(l.ch)
	}
	l.readChar()

	return out.Bytes()
}

func (l *Lexer) readNumber() []byte {
	var out bytes.Buffer

	for isDigit(l.ch) {
		out.WriteByte(l.ch)
		if !isDigit(l.peekChar()) {
			break
		}
		l.readChar()
	}

	return out.Bytes()
}

func (l *Lexer) readKeyword() string {
	var out bytes.Buffer

	for isLetter(l.ch) {
		out.WriteByte(l.ch)
		if !isLetter(l.peekChar()) {
			break
		}
		l.readChar()
	}

	return out.String()
}

func newToken(tokenType token.TokenType, ch []byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
