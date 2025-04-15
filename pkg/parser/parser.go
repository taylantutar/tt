package parser

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	ASTERISK = "*"
	SLASH  = "/"

	SET   = "SET"
	PRINT = "PRINT"
)

type Token struct {
	Type    TokenType
	Literal string
}