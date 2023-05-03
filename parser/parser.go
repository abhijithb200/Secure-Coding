package parser

import "fmt"

type TaintSpec struct {
	alias string // another name for the variable
	spec  Values // the vulnerable source

	scope string // which function name it is applicable
}

type VulnReport struct {
	name     string
	message  string
	position Position

	some Values

	taintvar map[string]TaintSpec
	allvar map[string]TaintSpec
}

var VulnTracker *VulnReport = &VulnReport{}
var StmtsNew = map[string]Node{}

var CurrFuncStatus = map[string][]struct {
	pos  int
	evil string
}{}

// add all the ArrayDim to the slice
var CSRFlist []string 



func Parser() {
	VulnTracker.taintvar = make(map[string]TaintSpec)
	VulnTracker.allvar = make(map[string]TaintSpec)

	program := Test()
	for _, r := range program.Stmts {
		switch r.(type) {
		case *Function:
			currFuncName := r.Out(ArgStore{
				from: "Disable",
			}).(FunctionNew).FunctionName

			StmtsNew[currFuncName] = r

		default:
			r.Out(ArgStore{})
		}

	}

	fmt.Println(VulnTracker.taintvar)
	fmt.Println(VulnTracker.allvar)
	fmt.Println(CSRFlist)

	var CSRFStatus bool = false;
	for _,r :=   range CSRFlist {
	if r == "csrf_token" {
		CSRFStatus = true
	}
	}
	if !CSRFStatus {
		vuln_reporter(
			&VulnReport{
				name:    "CSRF token missing",
			},
		)
	}
}
