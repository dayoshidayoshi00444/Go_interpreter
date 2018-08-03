package ast

import "github.com/dayoshidayoshi00444/Go_interpreter/02/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	NodestatementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
