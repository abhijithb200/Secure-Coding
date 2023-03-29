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

func (c *Concat) Out(argstore ArgStore) Values {

	// fmt.Println("right", reflect.TypeOf(c.Right.Out()).String())
	// fmt.Println("left", reflect.TypeOf(c.Left.Out()).String())

	a := c.Left.Out(argstore)
	b := c.Right.Out(argstore)

	// finding Vulneravle array used in expression in assign
	if argstore.from == "assign" {

		// if it contain the vulnerable arrray
		if reflect.TypeOf(a).String() == "parser.ArrayDimFetchNew" {

			if x := a.(ArrayDimFetchNew).Variable; x == "_GET" || x == "_POST" {
				p := make(map[string]TaintSpec)
				p[argstore.variable] = TaintSpec{
					alias: argstore.variable,
					spec:  a.(ArrayDimFetchNew),
				}
				VulnTracker.taintvar = append(VulnTracker.taintvar, p)
			}
		} else if reflect.TypeOf(b).String() == "parser.ArrayDimFetchNew" {

			if x := b.(ArrayDimFetchNew).Variable; x == "_GET" || x == "_POST" {
				p := make(map[string]TaintSpec)
				p[argstore.variable] = TaintSpec{
					alias: argstore.variable,
					spec:  b.(ArrayDimFetchNew),
				}
				VulnTracker.taintvar = append(VulnTracker.taintvar, p)
			}
		}

		// if it contain the tainted variable
		if reflect.TypeOf(a).String() == "parser.IdentifierNew" {

			for _, i := range VulnTracker.taintvar {
				for k, _ := range i {
					if a.(IdentifierNew).Value == k {

						p := make(map[string]TaintSpec)

						p[argstore.variable] = TaintSpec{
							alias: k,
							// spec:  v,
						}
						VulnTracker.taintvar = append(VulnTracker.taintvar, p)
					}
				}
			}

		} else if reflect.TypeOf(b).String() == "parser.IdentifierNew" {

			for _, i := range VulnTracker.taintvar {
				for k, _ := range i {

					// if expression contain tainted value, set the spec of current variable to it
					if b.(IdentifierNew).Value == k {

						p := make(map[string]TaintSpec)

						p[argstore.variable] = TaintSpec{
							alias: k,
							// spec:  v,
						}
						VulnTracker.taintvar = append(VulnTracker.taintvar, p)
					}
				}
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

	if argstore.from == "echo" {

		if reflect.TypeOf(b).String() == "parser.IdentifierNew" {

			x := b.(IdentifierNew).Value

			for _, i := range VulnTracker.taintvar {
				for k, v := range i {
					if x == k {
						vuln_reporter(&VulnReport{
							name:    "Reflected XSS ",
							message: "Found Tainted value " + x.(string) + " inside echo",
							some:    v,
							// position: *c.Position,
						})
					}
				}
			}
		} else if reflect.TypeOf(a).String() == "parser.IdentifierNew" {

			y := a.(IdentifierNew).Value

			for _, i := range VulnTracker.taintvar {

				for k, v := range i {
					if y == k {
						vuln_reporter(&VulnReport{
							name:    "Reflected XSS ",
							message: "Found Tainted value  inside echo",
							some:    v,
							// position: *c.Position,
						})
					}
				}
			}
		}
	}

	return 1

}
