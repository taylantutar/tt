package parser

import (
	"github.com/taylantutar/tt/cmd/lexer"
	"github.com/taylantutar/tt/pkg/token"
	"github.com/taylantutar/tt/pkg/statement"
	"strconv"
)



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

func (p *Parser) ParseProgram() []statement.Statement {
	var statements []statement.Statement

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.nextToken()
	}

	return statements
}

func (p *Parser) parseStatement() statement.Statement {
	switch p.curToken.Type {
	case token.SET:
		return p.parseSetStatement()
	case token.PRINT:
		return p.parsePrintStatement()
	default:
		return nil
	}
}

func (p *Parser) parseSetStatement() *statement.SetStatement {
	p.nextToken() // IDENT
	name := p.curToken.Literal

	p.nextToken() // "="
	p.nextToken() // value

	value := p.parseExpression()

	return &statement.SetStatement{Name: name, Value: value}
}

func (p *Parser) parsePrintStatement() *statement.PrintStatement {
	p.nextToken()
	expr := p.parseExpression()
	return &statement.PrintStatement{Expr: expr}
}

func (p *Parser) parseExpression() statement.Expression {
	left := p.parsePrimary()

	if p.peekToken.Type == token.PLUS || p.peekToken.Type == token.MINUS || p.peekToken.Type == token.ASTERISK || p.peekToken.Type == token.SLASH {
		op := p.peekToken.Literal
		p.nextToken() // operatöre geç
		p.nextToken() // sağ operand
		right := p.parsePrimary()
		return &statement.InfixExpression{Left: left, Operator: op, Right: right}
	}

	return left
}

func (p *Parser) parsePrimary() statement.Expression {
	switch p.curToken.Type {
	case token.IDENT:
		return &statement.Identifier{Value: p.curToken.Literal}
	case token.INT:
		val, _ := strconv.Atoi(p.curToken.Literal)
		return &statement.IntegerLiteral{Value: val}
	default:
		return nil
	}
}
