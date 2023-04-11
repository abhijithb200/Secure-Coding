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

	FunctionName Node

	Params []Node
	Stmts  []Node
}

type FunctionNew struct {
	FunctionName string
}

func (f *Function) Out(argstore ArgStore) Values {
	for _, r := range CurrFuncStatus["po"] {

		VulnTracker.taintvar[f.Params[r.pos].Out(ArgStore{}).(string)] = TaintSpec{
			alias: r.evil,
			scope: f.FunctionName.Out(ArgStore{}).(IdentifierNew).Value.(string),
		}
	}
	if argstore.from != "Disable" {
		for _, r := range f.Stmts {
			r.Out(ArgStore{})
		}
	}

	return FunctionNew{
		FunctionName: f.FunctionName.Out(ArgStore{}).(IdentifierNew).Value.(string),
	}
}

type Parameter struct {
	Variadic bool
	ByRef    bool
	Variable Node
}

func (p *Parameter) Out(argstore ArgStore) Values {
	return p.Variable.Out(ArgStore{}).(IdentifierNew).Value
}

type FunctionCall struct {
	Function     *Name
	ArgumentList Node
}

func (f *FunctionCall) Out(argstore ArgStore) Values {
	f.ArgumentList.Out(ArgStore{})
	return nil
}

type ArgumentList struct {
	Arguments []Node
}

func (f *ArgumentList) Out(argstore ArgStore) Values {
	for i, r := range f.Arguments {

		if _, ok := VulnTracker.taintvar[r.(*Argument).Expr.Out(ArgStore{}).(IdentifierNew).Value.(string)]; ok {
			CurrFuncStatus["po"] = append(CurrFuncStatus["po"], struct {
				pos  int
				evil string
			}{
				pos:  i,
				evil: r.(*Argument).Expr.Out(ArgStore{}).(IdentifierNew).Value.(string),
			})
		}
	}
	return nil
}

type Expression struct {
	Position *Position

	Expr Node
}

func (e *Expression) Out(argstore ArgStore) Values {
	switch e.Expr.(type) {
	case *FunctionCall:
		e.Expr.Out(ArgStore{})
		StmtsNew["writeMsg"].(*Function).Out(ArgStore{})
	default:
		e.Expr.Out(ArgStore{})
	}
	return nil
}

type Echo struct {
	Position *Position

	Exprs []Node
}

func (ec *Echo) Out(argstore ArgStore) Values {

	a := ec.Exprs[0].Out(ArgStore{
		from: "echo",
	})

	switch a.(type) {

	// adding tainted variable directly on the echo statement
	case IdentifierNew:
		for k, v := range VulnTracker.taintvar {
			if k == a.(IdentifierNew).Value {
				vuln_reporter(
					&VulnReport{
						name:    "Reflected XSS",
						message: "Found tainted variable" + a.(IdentifierNew).Value.(string) + " directly on the echo statement",
						some:    v,
					},
				)
			}
		}
	}
	return nil
}

func VulnSourceResolve(a TaintSpec) {
	// recursively finding the inner variable with spec
	for k, v := range VulnTracker.taintvar {
		if k == a.alias {
			if v.spec != nil {
				fmt.Println("Vulnerable Source :", v.spec)
				return
			} else {
				VulnSourceResolve(v)
			}
		}
	}

}

func vuln_reporter(a *VulnReport) {
	fmt.Print("[!]Vulnerability Found on line ", a.position.StartLine, "\n")
	fmt.Println("Type :", a.name)
	fmt.Println("Description :", a.message)

	switch a.some.(type) {

	//if it is inside the taintvar map
	case TaintSpec:
		VulnSourceResolve(a.some.(TaintSpec))

	// if the source is directly in the echo
	case ArrayDimFetchNew:
		fmt.Println("Vulnerable Source :", a.some)
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

	if reflect.TypeOf(x).String() == "parser.IdentifierNew" {
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
