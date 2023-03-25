package parser

import (
	"reflect"
)

/*
	Concatenation
*/

type Concat struct {
	Position *Position

	Left  Node
	Right Node
}

type ConcatNew struct {
	Left  Values
	Right Values
}

func (c *Concat) Out(from string) Values {

	// fmt.Println("right", reflect.TypeOf(c.Right.Out()).String())
	// fmt.Println("left", reflect.TypeOf(c.Left.Out()).String())

	a := c.Left.Out(from)
	b := c.Right.Out(from)

	// exit from recurssion when concatenation contain vuln array
	if a == CONCAT_WITH_VULN_ARRAY || b == CONCAT_WITH_VULN_ARRAY {
		return CONCAT_WITH_VULN_ARRAY
	}

	// finding Vulneravle array used in expression in assign
	if from == "assign" {
		if reflect.TypeOf(a).String() == "parser.ArrayDimFetchNew" {
			if x := a.(ArrayDimFetchNew).Variable; x == "_GET" || x == "_POST" {
				return CONCAT_WITH_VULN_ARRAY
			}
		} else if reflect.TypeOf(b).String() == "parser.ArrayDimFetchNew" {
			if x := b.(ArrayDimFetchNew).Variable; x == "_GET" || x == "_POST" {
				return CONCAT_WITH_VULN_ARRAY
			}
		}
	}

	// Finding $_GET[] or $_POST[] used directly in the echo statement

	if reflect.TypeOf(a).String() == "parser.ArrayDimFetchNew" {

		if x := a.(ArrayDimFetchNew).Variable; x == "_GET" || x == "_POST" {
			vuln_reporter(&VulnReport{
				name:    "Reflected XSS",
				message: "Found " + x.(string) + " inside echo with the parameter : " + a.(ArrayDimFetchNew).Value.(string),
				// position: *c.Position,
			})

		}
	} else if reflect.TypeOf(b).String() == "parser.ArrayDimFetchNew" {
		if y := b.(ArrayDimFetchNew).Variable; y == "_GET" || y == "_POST" {
			vuln_reporter(&VulnReport{
				name:    "Reflected XSS",
				message: "Found " + y.(string) + " inside echo with the parameter : " + b.(ArrayDimFetchNew).Value.(string),
				// position: *c.Position,
			})

		}
	}

	// Finding is there any tainted variable used on echo statement

	if from == "echo" {

		if reflect.TypeOf(b).String() == "parser.IdentifierNew" {

			x := b.(IdentifierNew).Value

			for _, i := range VulnTracker.taintvar {
				if x == i {
					vuln_reporter(&VulnReport{
						name:    "Reflected XSS ",
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
						name:    "Reflected XSS",
						message: "Found Tainted value " + i,
						// position: *c.Position,
					})
				}
			}
		}
	}

	return 1

}