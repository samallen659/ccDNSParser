package lexer

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
}
