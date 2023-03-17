package parser

type VulnReport struct {
	name     string
	message  string
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
							Value: "p",
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
			&Echo{
				Exprs: []Node{
					&Concat{
						Left: &String{
							Value: "'my name is'",
						},
						Right: &ConstFetch{
							Constant: &Name{
								Parts: []Node{
									&NamePart{
										Value: "p",
									},
								},
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
