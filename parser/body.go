package parser

import (
	"fmt"
	"reflect"
)

type Values interface {
}

type Demo interface {
	Write(p []byte) (n int, err error)
}

func Write(p []byte) (n int, err error) {
	fmt.Println("HI")

	return 1, nil
}

type Node interface {
	Out() Values
}

type Root struct {
	Stmts []Node `json:"Stmts"`
}

type Echo struct {
	Exprs []Node
}

func (ec *Echo) Out() Values {
	ec.Exprs[0].Out()

	return nil
}

type Concat struct {
	Left  Node
	Right Node
}

type AttackReport struct {
	name    string
	message string
}

func vuln_reporter(a *AttackReport) {
	fmt.Println(*a)
}

func (c *Concat) Out() Values {

	// fmt.Println("right", reflect.TypeOf(c.Right.Out()).String(), c.Right.Out())
	// fmt.Println("left", reflect.TypeOf(c.Left.Out()).String(), c.Left.Out())

	if reflect.TypeOf(c.Left.Out()).String() == "parser.ArrayDimFetchNew" {
		if a := c.Left.Out().(ArrayDimFetchNew).Variable; a == "_GET" {

			vuln_reporter(&AttackReport{
				name:    "XSS Echo",
				message: "Found _GET[] with the parameter :" + c.Left.Out().(ArrayDimFetchNew).Value.(string),
			})

		}
	} else if reflect.TypeOf(c.Right.Out()).String() == "parser.ArrayDimFetchNew" {
		if a := c.Right.Out().(ArrayDimFetchNew).Variable; a == "_GET" {
			vuln_reporter(&AttackReport{
				name:    "XSS Echo",
				message: "Found _GET[] with the parameter" + c.Right.Out().(ArrayDimFetchNew).Value.(string),
			})
		}
	}

	return 1

}

type ArrayDimFetch struct {
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
	VarName Node
}

func (v *Variable) Out() Values {

	return v.VarName.Out()
}

// out values

type String struct {
	Value string
}

func (s *String) Out() Values {

	return s.Value
}

type Identifier struct {
	Value string
}

func (i *Identifier) Out() Values {

	return i.Value
}
