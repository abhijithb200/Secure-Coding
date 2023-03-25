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
					Expression: &Concat{
						Left: &Concat{
							Left: &Concat{
								Left: &Concat{
									Left: &ArrayDimFetch{
										Variable: &Variable{
											VarName: &Identifier{
												Value: "_GET",
											},
										},
										Dim: &String{
											Value: "'name'",
										},
									},
									Right: &String{
										Value: "\"abhi\"",
									},
								},
								Right: &ArrayDimFetch{
									Variable: &Variable{
										VarName: &Identifier{
											Value: "_GET",
										},
									},
									Dim: &String{
										Value: "'age'",
									},
								},
							},
							Right: &String{
								Value: "\"is goin to be a super star in\"",
							},
						},
						Right: &ArrayDimFetch{
							Variable: &Variable{
								VarName: &Identifier{
									Value: "_GET",
								},
							},
							Dim: &String{
								Value: "'nation'",
							},
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
								Value: "a",
							},
						},
					},
				},
			},
		},
	}
	for _, r := range program.Stmts {

		r.Out("")

		// fmt.Println(VulnTracker.taintvar)

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
