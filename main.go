package main

import (
	"fmt"
	"github.com/taylantutar/tt/cmd/lexer"
	"github.com/taylantutar/tt/pkg/parser"
)

func main() {
	input := `
set a = 3
set b = 7
print a + b
`

	lexer := lexer.NewLexer(input)

	for tok := lexer.NextToken(); tok.Type != parser.EOF; tok = lexer.NextToken() {
		fmt.Printf("Token{Type: %-6s, Literal: %s}\n", tok.Type, tok.Literal)
	}
}
