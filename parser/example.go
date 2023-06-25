package parser

type program *Root

func Test() program {
    p:= &Root{
	Position: &Position{
		StartLine: 2,
		EndLine: 4,
		StartPos: 7,
		EndPos: 75,
	},
	Stmts: []Node{
		&Echo{
			Position: &Position{
				StartLine: 2,
				EndLine: 2,
				StartPos: 7,
				EndPos: 38,
			},
			Exprs: []Node{
				&Concat{
					Position: &Position{
						StartLine: 2,
						EndLine: 2,
						StartPos: 12,
						EndPos: 37,
					},
					Left: &String{
						Position: &Position{
							StartLine: 2,
							EndLine: 2,
							StartPos: 12,
							EndPos: 23,
						},
						Value: "\"Name is :\"",
					},
					Right: &ArrayDimFetch{
						Position: &Position{
							StartLine: 2,
							EndLine: 2,
							StartPos: 24,
							EndPos: 37,
						},
						Variable: &Variable{
							Position: &Position{
								StartLine: 2,
								EndLine: 2,
								StartPos: 24,
								EndPos: 29,
							},
							VarName: &Identifier{
								Position: &Position{
									StartLine: 2,
									EndLine: 2,
									StartPos: 24,
									EndPos: 29,
								},
								Value: "_GET",
							},
						},
						Dim: &String{
							Position: &Position{
								StartLine: 2,
								EndLine: 2,
								StartPos: 30,
								EndPos: 36,
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
				EndLine: 4,
				StartPos: 40,
				EndPos: 75,
			},
			Exprs: []Node{
				&Concat{
					Position: &Position{
						StartLine: 3,
						EndLine: 3,
						StartPos: 45,
						EndPos: 70,
					},
					Left: &String{
						Position: &Position{
							StartLine: 3,
							EndLine: 3,
							StartPos: 45,
							EndPos: 56,
						},
						Value: "\"Name is :\"",
					},
					Right: &ArrayDimFetch{
						Position: &Position{
							StartLine: 3,
							EndLine: 3,
							StartPos: 57,
							EndPos: 70,
						},
						Variable: &Variable{
							Position: &Position{
								StartLine: 3,
								EndLine: 3,
								StartPos: 57,
								EndPos: 62,
							},
							VarName: &Identifier{
								Position: &Position{
									StartLine: 3,
									EndLine: 3,
									StartPos: 57,
									EndPos: 62,
								},
								Value: "_GET",
							},
						},
						Dim: &String{
							Position: &Position{
								StartLine: 3,
								EndLine: 3,
								StartPos: 63,
								EndPos: 69,
							},
							Value: "'name'",
						},
					},
				},
			},
		},
	},
}

 return p 
}
