package parser

import (
	"fmt"
)

type AllVulns interface{}

type XSSReport struct {
	Value    string `json:"value"`
	Variable string `json:"variable"`
}

type Report struct {
	Type   string   `json:"type"`
	Source AllVulns `json:"source"`
}

var VulnStore []Report
var VulnOutput string

// global variable to handle the recursive functions data
var z Values

func VulnSourceResolve(a TaintSpec) {
	// recursively finding the inner variable with spec
	for k, v := range VulnTracker.taintvar {
		if k == a.alias {
			if v.spec != nil {
				VulnOutput += fmt.Sprintln("Vulnerable Source :", v.spec)
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
				VulnOutput += fmt.Sprintln("Vulnerable Source :", v.spec)
				return
			} else {
				VulnSourceResolve(v)
			}
		}
	}

}

func vuln_reporter(a *VulnReport) {
	VulnOutput += "----------------------------------------------------------\n"
	VulnOutput += fmt.Sprintln("[!]Vulnerability Found on line ", a.position.StartLine)
	VulnOutput += fmt.Sprintln("Type :", a.name)
	VulnOutput += fmt.Sprintln("Description :", a.message)

	switch a.some.(type) {

	//if it is inside the taintvar map
	case TaintSpec:
		VulnSourceResolve(a.some.(TaintSpec))
		VarSourceResolve(a.some.(TaintSpec))

	// if the vuln source is directly in the echo
	case ArrayDimFetchNew:
		z = a.some
		VulnOutput += fmt.Sprintln("Vulnerable Source :", a.some)
	}

	if a.name == "Reflected XSS" || a.name == "SQL Injection" {
		v := Report{
			Type: a.name,
			Source: XSSReport{
				Value:    z.(ArrayDimFetchNew).Value.(string),
				Variable: z.(ArrayDimFetchNew).Variable.(string),
			},
		}
		VulnStore = append(VulnStore, v)
	}

	// nullify the global variable
	z = 0

}
