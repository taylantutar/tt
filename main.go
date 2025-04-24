package main

import (
	"fmt"

	"github.com/taylantutar/tt/cmd/interpreter"
	"github.com/taylantutar/tt/cmd/lexer"
	"github.com/taylantutar/tt/cmd/parser"
)

func main() {
	input := `
set a = 10
set b = 5
print a + b
`

	lexer := lexer.NewLexer(input)
	fmt.Println("lexer: ",lexer)

	parser := parser.NewParser(lexer)
	fmt.Println("parser: ",parser)

	program := parser.ParseProgram()
	fmt.Println("program: ",program)

	env := interpreter.NewEnvironment()
	fmt.Println("env: ",env)

	interpreter.Eval(program, env)
}
