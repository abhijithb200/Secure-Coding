package parser

import (
	"fmt"
)

type AllVulns interface{}

type XSSReport struct {
	Value    string `json:"value"`
	Variable string `json:"variable"`
}

type SQLReport struct {
	Value     string `json:"value"`
	Variable  string `json:"variable"`
	Dbdetails map[string][]string `json:"dbdetails"`
}

type Report struct {
	Type   string   `json:"type"`
	Description   string   `json:"discription"`
	Position int `json:"position"`
	Source AllVulns `json:"source"`
	Severity int `json:"severity"`
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

	if a.name == "Reflected XSS" {
		v := Report{
			Type: a.name,
			Description: a.message,
			Position: a.position.StartLine,
			Severity: 1,
			Source: XSSReport{
				Value:    z.(ArrayDimFetchNew).Value.(string),
				Variable: z.(ArrayDimFetchNew).Variable.(string),
			},
		}
		VulnStore = append(VulnStore, v)
	}else if a.name == "SQL Injection"{
		v := Report{
			Type: a.name,
			Description: a.message,
			Severity: 1,
			Position: a.position.StartLine,
			Source: SQLReport{
				Value:    z.(ArrayDimFetchNew).Value.(string),
				Variable: z.(ArrayDimFetchNew).Variable.(string),
				Dbdetails: DatabaseDetails,
			},
		}
		VulnStore = append(VulnStore, v)
	}else if a.name == "CSRF token missing"{
		v := Report{
			Type: a.name,
			Severity: 3,
			Description: "",
			Position: 1,
		}
		VulnStore = append(VulnStore, v)
	}

	// nullify the global variable
	z = 0

}
