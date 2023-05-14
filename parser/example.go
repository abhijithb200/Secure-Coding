package parser

type program *Root

func Test() program {
    p:= &Root{
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
		&Echo{
			Exprs: []Node{
				&Concat{
					Left: &String{
						Value: "\"Name is\"",
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

 return p 
}
