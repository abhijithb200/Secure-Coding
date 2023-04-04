package parser

type TaintSpec struct {
	alias string
	spec  Values
}

type VulnReport struct {
	name     string
	message  string
	position Position

	some Values

	taintvar []map[string]TaintSpec
}

var VulnTracker *VulnReport = &VulnReport{}

func Parser() {
	program := &Root{
		Stmts: []Node{
			&Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &Identifier{
					Value: "writeMsg",
				},
				Stmts: []Node{
					&Echo{
						Exprs: []Node{
							&Concat{
								Left: &String{
									Value: "\"Hello world!\"",
								},
								Right: &ArrayDimFetch{
									Variable: &ConstFetch{
										Constant: &Name{
											Parts: []Node{
												&NamePart{
													Value: "_GET",
												},
											},
										},
									},
									Dim: &String{
										Value: "\"name\"",
									},
								},
							},
						},
					},
				},
			},
			&Expression{
				Expr: &FunctionCall{
					Function: &Name{
						Parts: []Node{
							&NamePart{
								Value: "writeMsg",
							},
						},
					},
					ArgumentList: &ArgumentList{},
				},
			},
		},
	}
	for _, r := range program.Stmts {

		r.Out(ArgStore{})

		// for k, v := range VulnTracker.taintvar[0] {
		// 	fmt.Println(k, v)
		// }

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
