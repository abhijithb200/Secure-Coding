package parser

import (
	"reflect"
)

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

func (a *Assign) Out(argstore ArgStore) Values {

	y := a.Variable.Out(argstore)
	x := a.Expression.Out(
		ArgStore{
			from:     "assign",
			variable: y.(IdentifierNew).Value.(string),
		})

	// if any variable accept value from _GET or _POST add to the tainted array list
	if reflect.TypeOf(x).String() == "parser.ArrayDimFetchNew" {
		if b := x.(ArrayDimFetchNew).Variable; b == "_GET" || b == "_POST" {
			p := make(map[string]TaintSpec)
			p[y.(IdentifierNew).Value.(string)] = TaintSpec{
				alias: "",
				spec:  x.(ArrayDimFetchNew),
			}
			VulnTracker.taintvar = append(VulnTracker.taintvar, p)
		}
	}

	// if any assignment have right varible tainted, add left variable to the tainted list
	if reflect.TypeOf(x).String() == "parser.IdentifierNew" {
		// fmt.Println(a.Variable.Out().(IdentifierNew).Value) // b
		// fmt.Println(x.(IdentifierNew).Value)                // a

		for _, i := range VulnTracker.taintvar {
			for k, _ := range i {
				if x.(IdentifierNew).Value == k {

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

	return nil
}
