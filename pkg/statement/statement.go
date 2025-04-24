package statement

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