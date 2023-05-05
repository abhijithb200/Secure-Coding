package parser

import (
	"encoding/json"
	"fmt"
)

type AllVulns interface{}

type XSSReport struct {
	Value    string `json:"value"`
	Variable string	`json:"variable"`
}

type Report struct {
	Type   string `json:"type"`
	Source AllVulns `json:"source"`
}

type VulnStore []Report

// global variable to handle the recursive functions data
var z Values 

func VulnSourceResolve(a TaintSpec)  {
	// recursively finding the inner variable with spec
	for k, v := range VulnTracker.taintvar {
		if k == a.alias {
			if v.spec != nil {
				fmt.Println("Vulnerable Source :", v.spec)
				z = v.spec
				return 
			} else {
				VulnSourceResolve(v)
			}
		}

		
	}
}
func VarSourceResolve(a TaintSpec) {
	// recursively finding the inner variable with spec
	for k, v := range VulnTracker.allvar {
		if k == a.alias {
			if v.spec != nil {
				fmt.Println("Vulnerable Source :", v.spec)
				return
			} else {
				VulnSourceResolve(v)
			}
		}
	}

}

func vuln_reporter(a *VulnReport) {
	fmt.Print("[!]Vulnerability Found on line ", a.position.StartLine, "\n")
	fmt.Println("Type :", a.name)
	fmt.Println("Description :", a.message)
	


	switch a.some.(type) {

	//if it is inside the taintvar map
	case TaintSpec:
		VulnSourceResolve(a.some.(TaintSpec))
		VarSourceResolve(a.some.(TaintSpec))

	// if the vuln source is directly in the echo
	case ArrayDimFetchNew:
		z = a.some
		fmt.Println("Vulnerable Source :", a.some)
	}


	if a.name == "Reflected XSS " {
		v := Report{
			Type: a.name,
			Source: XSSReport{
				Value: z.(ArrayDimFetchNew).Value.(string),
				Variable: z.(ArrayDimFetchNew).Variable.(string),
			},
		}
		p,_ := json.Marshal(v)
		
		fmt.Println(string(p))
	}
	

	// nullify the global variable
	z = 0
	fmt.Println("----------------------------------------------------")
	fmt.Println()
}
