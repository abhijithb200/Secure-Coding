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
		Position: &Position{
			StartLine: 2,
			EndLine:   3,
			StartPos:  7,
			EndPos:    79,
		},
		Stmts: []Node{
			&Echo{
				Position: &Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  7,
					EndPos:    39,
				},
				Exprs: []Node{
					&Concat{
						Position: &Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  12,
							EndPos:    38,
						},
						Left: &String{
							Position: &Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  12,
								EndPos:    24,
							},
							Value: "'my name is'",
						},
						Right: &ArrayDimFetch{
							Position: &Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  25,
								EndPos:    38,
							},
							Variable: &Variable{
								Position: &Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  25,
									EndPos:    30,
								},
								VarName: &Identifier{
									Position: &Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  25,
										EndPos:    30,
									},
									Value: "_GET",
								},
							},
							Dim: &String{
								Position: &Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  31,
									EndPos:    37,
								},
								Value: "'name'",
							},
						},
					},
				},
			},
			&Echo{
				Position: &Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  41,
					EndPos:    79,
				},
				Exprs: []Node{
					&Concat{
						Position: &Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  46,
							EndPos:    78,
						},
						Left: &String{
							Position: &Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  46,
								EndPos:    61,
							},
							Value: "'my address is'",
						},
						Right: &ArrayDimFetch{
							Position: &Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  62,
								EndPos:    78,
							},
							Variable: &Variable{
								Position: &Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  62,
									EndPos:    67,
								},
								VarName: &Identifier{
									Position: &Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  62,
										EndPos:    67,
									},
									Value: "_GET",
								},
							},
							Dim: &String{
								Position: &Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  68,
									EndPos:    77,
								},
								Value: "'address'",
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
