package parser

import (
	"github.com/taylantutar/tt/cmd/lexer"
	"github.com/taylantutar/tt/pkg/token"
	"strconv"
)

type Statement interface {
	statementNode()
}

type Expression interface {
	expressionNode()
}

type SetStatement struct {
	Name  string
	Value Expression
}

func (s *SetStatement) statementNode() {}

type PrintStatement struct {
	Expr Expression
}

func (s *PrintStatement) statementNode() {}

type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}

type IntegerLiteral struct {
	Value int
}

func (i *IntegerLiteral) expressionNode() {}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() []Statement {
	var statements []Statement

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.nextToken()
	}

	return statements
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case token.SET:
		return p.parseSetStatement()
	case token.PRINT:
		return p.parsePrintStatement()
	default:
		return nil
	}
}

func (p *Parser) parseSetStatement() *SetStatement {
	p.nextToken() // IDENT
	name := p.curToken.Literal

	p.nextToken() // "="
	p.nextToken() // value

	value := p.parseExpression()

	return &SetStatement{Name: name, Value: value}
}

func (p *Parser) parsePrintStatement() *PrintStatement {
	p.nextToken()
	expr := p.parseExpression()
	return &PrintStatement{Expr: expr}
}

func (p *Parser) parseExpression() Expression {
	left := p.parsePrimary()

	if p.peekToken.Type == token.PLUS || p.peekToken.Type == token.MINUS || p.peekToken.Type == token.ASTERISK || p.peekToken.Type == token.SLASH {
		op := p.peekToken.Literal
		p.nextToken() // operatöre geç
		p.nextToken() // sağ operand
		right := p.parsePrimary()
		return &InfixExpression{Left: left, Operator: op, Right: right}
	}

	return left
}

func (p *Parser) parsePrimary() Expression {
	switch p.curToken.Type {
	case token.IDENT:
		return &Identifier{Value: p.curToken.Literal}
	case token.INT:
		val, _ := strconv.Atoi(p.curToken.Literal)
		return &IntegerLiteral{Value: val}
	default:
		return nil
	}
}
