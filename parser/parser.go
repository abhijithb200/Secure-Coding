package parser

type VulnReport struct {
	name     string
	message  string
	position Position
	taintvar []string
}

var VulnTracker *VulnReport = &VulnReport{}

func Parser() {
	program := &Root{
		Stmts: []Node{
			&Expression{
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "a",
						},
					},
					Expression: &ArrayDimFetch{
						Variable: &Variable{
							VarName: &Identifier{
								Value: "_GET",
							},
						},
						Dim: &String{
							Value: "'name'",
						},
					},
				},
			},
			&Expression{
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "b",
						},
					},
					Expression: &Variable{
						VarName: &Identifier{
							Value: "a",
						},
					},
				},
			},
			&Echo{
				Exprs: []Node{
					&Concat{
						Left: &String{
							Value: "'my address is'",
						},
						Right: &Variable{
							VarName: &Identifier{
								Value: "b",
							},
						},
					},
				},
			},
		},
	}
	for _, r := range program.Stmts {

		r.Out()

		// switch r.(type) Println()
		// case *Echo:
		// 	// s := r.(*Echo).Exprs[0]
		// 	// m := s.(*String)
		// 	// fmt.Println(m.Value)
		// 	r.Out()

		// case *Expression:
		// 	r.Out()

		// default:
		// 	fmt.Println("no")
		// }

	}
}
