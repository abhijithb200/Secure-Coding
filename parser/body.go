package parser

import (
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
	r := f.ArgumentList.Out(ArgStore{
		variable: f.Function.Parts[0].Out(ArgStore{}).(string),
	})

	// if the function name is exec and the argument is in the tainted list
	if r != nil {
		if f.Function.Parts[0].Out(ArgStore{}).(string) == "exec" {

			vuln_reporter(&VulnReport{
				name:    "OS Command Injection",
				message: "Found " + r.(*Argument).Expr.Out(ArgStore{}).(IdentifierNew).Value.(string) + " inside exec",
				some:    VulnTracker.taintvar[r.(*Argument).Expr.Out(ArgStore{}).(IdentifierNew).Value.(string)],
				// position: *c.Position,
			})

		}
	}

	return 0
}

type ArgumentList struct {
	Arguments []Node
}

func (f *ArgumentList) Out(argstore ArgStore) Values {
	for i, r := range f.Arguments {

		// if tainted value is in the argument
		if _, ok := VulnTracker.taintvar[r.(*Argument).Expr.Out(ArgStore{}).(IdentifierNew).Value.(string)]; ok {
			CurrFuncStatus["po"] = append(CurrFuncStatus["po"], struct {
				pos  int
				evil string
			}{
				pos:  i,
				evil: r.(*Argument).Expr.Out(ArgStore{}).(IdentifierNew).Value.(string),
			})

			return r
		}
	}

	if argstore.variable == "mysqli_connect"{

		// check for third argument(password for connection) is in allvar[] or not
		if _,ok := VulnTracker.allvar[f.Arguments[2].Out(ArgStore{}).(IdentifierNew).Value.(string)]; ok{
			vuln_reporter(&VulnReport{
				name:    "Hardcoded Credentials",
				message: "Found " + f.Arguments[2].Out(ArgStore{}).(IdentifierNew).Value.(string) + " inside mysqli_connect",
				some:    VulnTracker.allvar[f.Arguments[2].Out(ArgStore{}).(IdentifierNew).Value.(string)],
				// position: *c.Position,
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

	CSRFlist = append(CSRFlist,y.(string) )


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

// Local file inclusion

type Include struct {
	Expr Node
}

func (i *Include) Out(argstore ArgStore) Values {
	a := i.Expr.Out(ArgStore{})

	// encapsed find the tainted variable and send back the variable
	if reflect.TypeOf(a).String() == "parser.IdentifierNew" {
		vuln_reporter(&VulnReport{
			name:    "Local File Inclusion",
			message: "Found " + a.(IdentifierNew).Value.(string) + " inside includes",
			some:    VulnTracker.taintvar[a.(IdentifierNew).Value.(string)],
			// position: *c.Position,
		})
	}

	return ""
}

type Encapsed struct {
	Parts []Node
}

func (e *Encapsed) Out(argstore ArgStore) Values {

	for _, i := range e.Parts {

		a := i.Out(ArgStore{})
		//if tainted variable is seen inside Encapsed return the value
		if reflect.TypeOf(a).String() == "parser.IdentifierNew" {
			if _, ok := VulnTracker.taintvar[a.(IdentifierNew).Value.(string)]; ok {
				return a
			}
		}
	}
	return ""
}

type EncapsedStringPart struct {
	Value string
}

func (e *EncapsedStringPart) Out(argstore ArgStore) Values {
	return e.Value
}

// CSRF Protection

type InlineHtml struct {
	Value string
}

func (i *InlineHtml) Out(argstore ArgStore) Values {
	return i.Value
}
