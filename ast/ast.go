import "Splice-Core/token"
package ast

type Node interface{ TokenLiteral() string}

type Program struct {
	Statements []Statement
}

func (x *Program) TokenLiteral() string{
	if len(x.Statements) > 0{
		return x.Statements[0].TokenLiteral()
	}
	else{
		return ""
	}
}



type Statement interface{ 
	Node
	statementNode()
}

type Expression interface{
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func(ls *LetStatement) statementNode() {}
func(ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct{
	Token token.Token
	Value string
}

func(i *Identifier) expressionNode() {}
func(i *Identifier) TokenLiteral() string {return i.Token.Literal}

