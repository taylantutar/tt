package main

import (
	"fmt"
	"github.com/taylantutar/tt/cmd/lexer"
	"github.com/taylantutar/tt/cmd/parser"

)

func main() {
	input := `
set a = 3
set b = 7
print a + b
`

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	program := parser.ParseProgram()

	for _, stmt := range program {
		fmt.Printf("%#v\n", stmt)
	}
}
