package lexer

import (
	"testing"

	"github.com/mdaisuke/goscheme/token"
)

func TestNextToken(t *testing.T) {
	input := `
	(+ (* 1 2 3) (/ 10 10))
	(define e 23)
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.ASTERISK, "*"},
		{token.INT, "1"},
		{token.INT, "2"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.SLASH, "/"},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},

		{token.LPAREN, "("},
		{token.DEFINE, "define"},
		{token.IDENT, "e"},
		{token.INT, "23"},
		{token.RPAREN, ")"},

		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("[%d] wrong type. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("[%d] wrong literal. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
