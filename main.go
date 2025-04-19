package main

import (
	"github.com/taylantutar/tt/cmd/interpreter"
	"github.com/taylantutar/tt/cmd/lexer"
	"github.com/taylantutar/tt/cmd/parser"
)

func main() {
	input := `
set a = 10
set b = 5
print a + b
print a * b
`

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)
	program := parser.ParseProgram()

	env := interpreter.NewEnvironment()
	interpreter.Eval(program, env)
}
