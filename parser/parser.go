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
}

var VulnTracker *VulnReport = &VulnReport{}
var StmtsNew = map[string]Node{}

var CurrFuncStatus = map[string][]struct {
	pos  int
	evil string
}{}

func Parser() {
	VulnTracker.taintvar = make(map[string]TaintSpec)

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
}
