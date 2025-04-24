package interpreter

import (
	"fmt"
	"github.com/taylantutar/tt/pkg/statement"
)

type Environment struct {
	store map[string]int
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]int)}
}

func (e *Environment) Set(name string, val int) {
	e.store[name] = val
}

func (e *Environment) Get(name string) (int, bool) {
	val, ok := e.store[name]
	return val, ok
}

func Eval(stmts []statement.Statement, env *Environment) {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *statement.SetStatement:
			val := evalExpression(s.Value, env)
			env.Set(s.Name, val)
		case *statement.PrintStatement:
			val := evalExpression(s.Expr, env)
			fmt.Println(val)
		}
	}
}

func evalExpression(expr statement.Expression, env *Environment) int {
	switch e := expr.(type) {
	case *statement.IntegerLiteral:
		return e.Value
	case *statement.Identifier:
		val, ok := env.Get(e.Value)
		if !ok {
			panic("Tanımsız değişken: " + e.Value)
		}
		return val
	case *statement.InfixExpression:
		left := evalExpression(e.Left, env)
		right := evalExpression(e.Right, env)
		switch e.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		default:
			panic("Bilinmeyen operator: " + e.Operator)
		}
	default:
		panic("Bilinmeyen ifade")
	}
}
