package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT = "INT"

	LPAREN = "("
	RPAREN = ")"

	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
)

type Token struct {
	Type    TokenType
	Literal string
}
