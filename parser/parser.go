package parser

type TaintSpec struct {
	alias string
	spec  Values

	scope string
}

type VulnReport struct {
	name     string
	message  string
	position Position

	some Values

	taintvar map[string]TaintSpec
}

type FuncContainer struct {
	pos  int
	evil string
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
}
