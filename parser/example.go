package parser

type program *Root

func Test() program {
    p:= &Root{
	Position: &Position{
		StartLine: 2,
		EndLine: 5,
		StartPos: 7,
		EndPos: 62,
	},
	Stmts: []Node{
		&Expression{
			Position: &Position{
				StartLine: 2,
				EndLine: 2,
				StartPos: 7,
				EndPos: 24,
			},
			Expr: &Assign{
				Position: &Position{
					StartLine: 2,
					EndLine: 2,
					StartPos: 7,
					EndPos: 23,
				},
				Variable: &Variable{
					Position: &Position{
						StartLine: 2,
						EndLine: 2,
						StartPos: 7,
						EndPos: 9,
					},
					VarName: &Identifier{
						Position: &Position{
							StartLine: 2,
							EndLine: 2,
							StartPos: 7,
							EndPos: 9,
						},
						Value: "m",
					},
				},
				Expression: &ArrayDimFetch{
					Position: &Position{
						StartLine: 2,
						EndLine: 2,
						StartPos: 10,
						EndPos: 23,
					},
					Variable: &Variable{
						Position: &Position{
							StartLine: 2,
							EndLine: 2,
							StartPos: 10,
							EndPos: 15,
						},
						VarName: &Identifier{
							Position: &Position{
								StartLine: 2,
								EndLine: 2,
								StartPos: 10,
								EndPos: 15,
							},
							Value: "_GET",
						},
					},
					Dim: &String{
						Position: &Position{
							StartLine: 2,
							EndLine: 2,
							StartPos: 16,
							EndPos: 22,
						},
						Value: "'Nice'",
					},
				},
			},
		},
		&Expression{
			Position: &Position{
				StartLine: 3,
				EndLine: 3,
				StartPos: 26,
				EndPos: 43,
			},
			Expr: &Assign{
				Position: &Position{
					StartLine: 3,
					EndLine: 3,
					StartPos: 26,
					EndPos: 42,
				},
				Variable: &Variable{
					Position: &Position{
						StartLine: 3,
						EndLine: 3,
						StartPos: 26,
						EndPos: 28,
					},
					VarName: &Identifier{
						Position: &Position{
							StartLine: 3,
							EndLine: 3,
							StartPos: 26,
							EndPos: 28,
						},
						Value: "c",
					},
				},
				Expression: &ArrayDimFetch{
					Position: &Position{
						StartLine: 3,
						EndLine: 3,
						StartPos: 29,
						EndPos: 42,
					},
					Variable: &Variable{
						Position: &Position{
							StartLine: 3,
							EndLine: 3,
							StartPos: 29,
							EndPos: 34,
						},
						VarName: &Identifier{
							Position: &Position{
								StartLine: 3,
								EndLine: 3,
								StartPos: 29,
								EndPos: 34,
							},
							Value: "_GET",
						},
					},
					Dim: &String{
						Position: &Position{
							StartLine: 3,
							EndLine: 3,
							StartPos: 35,
							EndPos: 41,
						},
						Value: "'Nice'",
					},
				},
			},
		},
		&Echo{
			Position: &Position{
				StartLine: 4,
				EndLine: 5,
				StartPos: 45,
				EndPos: 62,
			},
			Exprs: []Node{
				&Concat{
					Position: &Position{
						StartLine: 4,
						EndLine: 4,
						StartPos: 50,
						EndPos: 57,
					},
					Left: &String{
						Position: &Position{
							StartLine: 4,
							EndLine: 4,
							StartPos: 50,
							EndPos: 54,
						},
						Value: "\"Hi\"",
					},
					Right: &Variable{
						Position: &Position{
							StartLine: 4,
							EndLine: 4,
							StartPos: 55,
							EndPos: 57,
						},
						VarName: &Identifier{
							Position: &Position{
								StartLine: 4,
								EndLine: 4,
								StartPos: 55,
								EndPos: 57,
							},
							Value: "m",
						},
					},
				},
			},
		},
	},
}

 return p 
}
