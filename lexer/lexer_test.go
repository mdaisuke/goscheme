package lexer

import (
	"testing"

	"github.com/mdaisuke/goscheme/token"
)

func TestNextToken(t *testing.T) {
	input := `(* (- 2 1) (+ 3 4))`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.ASTERISK, "*"},
		{token.LPAREN, "("},
		{token.MINUS, "-"},
		{token.INT, "2"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
