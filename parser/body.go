package parser

import (
	"fmt"
	"reflect"
)

type Values interface {
}

type Node interface {
	Out() Values
}

type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

type Root struct {
	Position *Position
	Stmts    []Node
}

type Expression struct {
	Expr Node
}

func (e *Expression) Out() Values {
	e.Expr.Out()
	return nil
}

type Echo struct {
	Position *Position

	Exprs []Node
}

func (ec *Echo) Out() Values {
	ec.Exprs[0].Out()
	return nil
}

/*
	Concatenation
*/

type Concat struct {
	Position *Position

	Left  Node
	Right Node
}

func vuln_reporter(a *VulnReport) {
	fmt.Print("[!]Vulnerability Found\n")
	fmt.Println("Name :", a.name)
	fmt.Println("Description :", a.message)
}

func (c *Concat) Out() Values {

	// fmt.Println("right", reflect.TypeOf(c.Right.Out()).String(), c.Right.Out())
	// fmt.Println("left", reflect.TypeOf(c.Left.Out()).String(), c.Left.Out())

	// Finding $_GET[] or $_POST[] used directly in the echo statement

	if reflect.TypeOf(c.Left.Out()).String() == "parser.ArrayDimFetchNew" {
		if a := c.Left.Out().(ArrayDimFetchNew).Variable; a == "_GET" {
			vuln_reporter(&VulnReport{
				name:     "XSS Echo",
				message:  "Found _GET[] with the parameter : " + c.Left.Out().(ArrayDimFetchNew).Value.(string),
				position: *c.Position,
			})

			fmt.Println(c.Position.EndPos)
		}
	} else if reflect.TypeOf(c.Right.Out()).String() == "parser.ArrayDimFetchNew" {
		if a := c.Right.Out().(ArrayDimFetchNew).Variable; a == "_GET" {
			vuln_reporter(&VulnReport{
				name:     "XSS Echo",
				message:  "Found _GET[] with the parameter : " + c.Right.Out().(ArrayDimFetchNew).Value.(string),
				position: *c.Position,
			})

		}
	}

	// Finding is there any tainted variable used on echo statement

	for _, i := range VulnTracker.taintvar {
		if c.Right.Out() == i || c.Left.Out() == i {
			vuln_reporter(&VulnReport{
				name:     "XSS Echo Parameter",
				message:  "Found Tainted value " + i,
				position: *c.Position,
			})
		}
	}

	return 1

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

func (a *ArrayDimFetch) Out() Values {
	return ArrayDimFetchNew{Variable: a.Variable.Out(), Value: a.Dim.Out()}
}

type Variable struct {
	Position *Position

	VarName Node
}

func (v *Variable) Out() Values {

	return v.VarName.Out()
}

type Name struct {
	Parts []Node
}

func (n *Name) Out() Values {
	for _, i := range n.Parts {
		return i.Out()
	}

	return nil
}

type ConstFetch struct {
	Constant Node
}

func (c *ConstFetch) Out() Values {
	return c.Constant.Out()
}

/*
 Assignment operator (=)
 Left - Variable
 Right - Expression
*/

type Assign struct {
	Variable   Node
	Expression Node
}

func (a *Assign) Out() Values {
	// a.Expression.Out() // right
	// a.Variable.Out()   // left

	if reflect.TypeOf(a.Expression.Out()).String() == "parser.ArrayDimFetchNew" {
		if b := a.Expression.Out().(ArrayDimFetchNew).Variable; b == "_GET" {
			VulnTracker.taintvar = append(VulnTracker.taintvar, a.Variable.Out().(string))
		}
	}

	return nil
}

/*
 The last struct that return something
*/

type String struct {
	Position *Position

	Value string
}

func (s *String) Out() Values {

	return s.Value
}

type Identifier struct {
	Position *Position

	Value string
}

func (i *Identifier) Out() Values {

	return i.Value
}

type NamePart struct {
	Value string
}

func (n *NamePart) Out() Values {
	return n.Value
}
