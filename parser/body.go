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

	Stmts []Node
}

type Expression struct {
	Position *Position

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
	fmt.Print("[!]Vulnerability Found on line ", a.position.StartLine, "\n")
	fmt.Println("Type :", a.name)
	fmt.Println("Description :", a.message)
	fmt.Println("----------------------------------------------------")
	fmt.Println()
}

func (c *Concat) Out() Values {

	// fmt.Println("right", reflect.TypeOf(c.Right.Out()).String())
	// fmt.Println("left", reflect.TypeOf(c.Left.Out()).String())

	a := c.Left.Out()
	b := c.Right.Out()

	// Finding $_GET[] or $_POST[] used directly in the echo statement

	if reflect.TypeOf(a).String() == "parser.ArrayDimFetchNew" {
		if x := a.(ArrayDimFetchNew).Variable; x == "_GET" || x == "_POST" {
			vuln_reporter(&VulnReport{
				name:    "XSS Echo",
				message: "Found " + x.(string) + " inside echo with the parameter : " + a.(ArrayDimFetchNew).Value.(string),
				// position: *c.Position,
			})

		}
	} else if reflect.TypeOf(b).String() == "parser.ArrayDimFetchNew" {
		if y := b.(ArrayDimFetchNew).Variable; y == "_GET" || y == "_POST" {
			vuln_reporter(&VulnReport{
				name:    "XSS Echo",
				message: "Found " + y.(string) + " inside echo with the parameter : " + b.(ArrayDimFetchNew).Value.(string),
				// position: *c.Position,
			})

		}
	}

	// Finding is there any tainted variable used on echo statement

	if reflect.TypeOf(b).String() == "parser.IdentifierNew" {

		x := b.(IdentifierNew).Value

		for _, i := range VulnTracker.taintvar {
			if x == i {
				vuln_reporter(&VulnReport{
					name:    "XSS Echo Parameter ",
					message: "Found Tainted value " + i,
					// position: *c.Position,
				})
			}
		}
	} else if reflect.TypeOf(a).String() == "parser.IdentifierNew" {

		y := a.(IdentifierNew).Value

		for _, i := range VulnTracker.taintvar {

			if y == i {
				vuln_reporter(&VulnReport{
					name:    "XSS Echo Parameter",
					message: "Found Tainted value " + i,
					// position: *c.Position,
				})
			}
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
	return ArrayDimFetchNew{Variable: a.Variable.Out().(IdentifierNew).Value, Value: a.Dim.Out()}
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
	Position   *Position
	Variable   Node
	Expression Node
}

func (a *Assign) Out() Values {
	// a.Expression.Out() // right
	// a.Variable.Out()   // left

	x := a.Expression.Out()
	y := a.Variable.Out()

	// if any variable accept value from _GET or _POST add to the tainted array list
	if reflect.TypeOf(x).String() == "parser.ArrayDimFetchNew" {
		if b := x.(ArrayDimFetchNew).Variable; b == "_GET" || b == "_POST" {
			VulnTracker.taintvar = append(VulnTracker.taintvar, y.(IdentifierNew).Value.(string))
		}
	}

	// if any assignment have right varible tainted, add left variable to the tainted list
	if reflect.TypeOf(x).String() == "parser.IdentifierNew" {
		// fmt.Println(a.Variable.Out().(IdentifierNew).Value) // b
		// fmt.Println(x.(IdentifierNew).Value)                // a

		for _, i := range VulnTracker.taintvar {
			if x.(IdentifierNew).Value == i {
				VulnTracker.taintvar = append(VulnTracker.taintvar, y.(IdentifierNew).Value.(string))
			}
		}

		fmt.Println(VulnTracker.taintvar)
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

	Value Values
}

// created identifiernew to distinguish between string an actual identifier
type IdentifierNew struct {
	Value Values
}

func (i *Identifier) Out() Values {

	return IdentifierNew{Value: i.Value}
}

type NamePart struct {
	Value string
}

func (n *NamePart) Out() Values {
	return n.Value
}
