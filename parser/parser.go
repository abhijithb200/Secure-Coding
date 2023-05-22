package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	allvar   map[string]TaintSpec
}

var VulnTracker *VulnReport = &VulnReport{}
var StmtsNew = map[string]Node{}

var CurrFuncStatus = map[string][]struct {
	pos  int
	evil string
}{}

// add all the ArrayDim to the slice
var CSRFlist []string

type FinalReport struct {
	Hash       string   `json:"hash"`
	Everything string   `json:"everything"`
	Vulns      []Report `json:"vulns"`
}

var DatabaseDetails = map[string][]string{}

func Parser(hash string) {
	VulnTracker.taintvar = make(map[string]TaintSpec)
	VulnTracker.allvar = make(map[string]TaintSpec)

	DatabaseDetails = make(map[string][]string)

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

	// f, err := os.Create("../source/Codeguardian.json")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // Write the package declaration to the file
	// _, err = fmt.Fprintf(f, string(p))
	// if err != nil {
	// 	panic(err)
	// }

	var CSRFStatus bool = false
	for _, r := range CSRFlist {
		if r == "csrf_token" {
			CSRFStatus = true
		}
	}
	if !CSRFStatus {
		vuln_reporter(
			&VulnReport{
				name: "CSRF token missing",
			},
		)
	}

	fmt.Println(VulnTracker.allvar)
	fmt.Println(DatabaseDetails)

	b := FinalReport{
		Hash:       hash,
		Everything: VulnOutput,
		Vulns:      VulnStore,
	}
	p, _ := json.Marshal(b)
	os.Stdout.Write(p)
	// fmt.Println(b.Everything)
	VulnOutput = ""
}
