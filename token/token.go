package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT   = "INT"
	IDENT = "IDENT"

	LPAREN = "("
	RPAREN = ")"

	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"

	DEFINE = "DEFINE"
)

var keywords = map[string]TokenType{
	"define": DEFINE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupKeywords(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
