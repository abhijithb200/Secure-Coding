package parser

import (
	"fmt"
)

type Values interface {
}

type Node interface {
	Out(from string) Values
}

type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

type Root struct {
	Position *Position

	Stmts []Node
}

type Expression struct {
	Position *Position

	Expr Node
}

func (e *Expression) Out(from string) Values {
	e.Expr.Out("")
	return nil
}

type Echo struct {
	Position *Position

	Exprs []Node
}

func (ec *Echo) Out(from string) Values {
	ec.Exprs[0].Out("echo")
	return nil
}

func vuln_reporter(a *VulnReport) {
	fmt.Print("[!]Vulnerability Found on line ", a.position.StartLine, "\n")
	fmt.Println("Type :", a.name)
	fmt.Println("Description :", a.message)
	fmt.Println("----------------------------------------------------")
	fmt.Println()
}

/*
Array - Dimention and Value
*/

type ArrayDimFetch struct {
	Position *Position

	Variable Node
	Dim      Node
}

type ArrayDimFetchNew struct {
	Variable Values
	Value    Values
}

func (a *ArrayDimFetch) Out(from string) Values {
	return ArrayDimFetchNew{Variable: a.Variable.Out("").(IdentifierNew).Value, Value: a.Dim.Out("")}
}

type Variable struct {
	Position *Position

	VarName Node
}

func (v *Variable) Out(from string) Values {

	return v.VarName.Out("")
}

type Name struct {
	Parts []Node
}

func (n *Name) Out(from string) Values {
	for _, i := range n.Parts {
		return i.Out("")
	}

	return nil
}

type ConstFetch struct {
	Constant Node
}

func (c *ConstFetch) Out(from string) Values {
	return c.Constant.Out("")
}
