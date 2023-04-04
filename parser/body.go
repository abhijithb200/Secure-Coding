package parser

import (
	"fmt"
	"reflect"
)

type Values interface {
}

type Node interface {
	Out(argstore ArgStore) Values
}

// used to pass argument across functions
type ArgStore struct {
	from     string
	variable string
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

type Function struct {
	ReturnsRef    bool
	PhpDocComment string
	FunctionName  Node
	Stmts         []Node
}

func (f *Function) Out(argstore ArgStore) Values {
	for _, r := range f.Stmts {
		r.Out(ArgStore{})
	}

	return nil
}

type FunctionCall struct {
	Function     *Name
	ArgumentList Node
}

func (f *FunctionCall) Out(argstore ArgStore) Values {
	return nil
}

type ArgumentList struct {
}

func (f *ArgumentList) Out(argstore ArgStore) Values {
	return nil
}

type Expression struct {
	Position *Position

	Expr Node
}

func (e *Expression) Out(argstore ArgStore) Values {
	e.Expr.Out(ArgStore{})
	return nil
}

type Echo struct {
	Position *Position

	Exprs []Node
}

func (ec *Echo) Out(argstore ArgStore) Values {

	ec.Exprs[0].Out(ArgStore{
		from: "echo",
	})
	return nil
}

func vuln_reporter(a *VulnReport) {
	fmt.Print("[!]Vulnerability Found on line ", a.position.StartLine, "\n")
	fmt.Println("Type :", a.name)
	fmt.Println("Description :", a.message)

	for _, i := range VulnTracker.taintvar {
		for k, v := range i {
			if k == a.some.(TaintSpec).alias {
				fmt.Println("Vulnerable Source :", v.spec)
			}
		}
	}
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

func (a *ArrayDimFetch) Out(argstore ArgStore) Values {
	x := a.Variable.Out(ArgStore{})
	y := a.Dim.Out(ArgStore{})

	if reflect.TypeOf(x).String() == "parse.IdentifierNew" {
		return ArrayDimFetchNew{Variable: x.(IdentifierNew).Value, Value: y}
	} else if reflect.TypeOf(x).String() == "string" {
		return ArrayDimFetchNew{Variable: x, Value: y}
	}

	return nil
}

type Variable struct {
	Position *Position

	VarName Node
}

func (v *Variable) Out(argstore ArgStore) Values {

	return v.VarName.Out(ArgStore{})
}

type Name struct {
	Parts []Node
}

func (n *Name) Out(argstore ArgStore) Values {
	for _, i := range n.Parts {
		return i.Out(ArgStore{})
	}

	return nil
}

type ConstFetch struct {
	Constant Node
}

func (c *ConstFetch) Out(argstore ArgStore) Values {
	return c.Constant.Out(ArgStore{})
}
