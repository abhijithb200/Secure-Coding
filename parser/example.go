package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
		&Echo{
			Exprs: []Node{
				&Concat{
					Left: &String{
						Value: "\"abhi\"",
					},
					Right: &ArrayDimFetch{
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
		},
	},
}

 return p 
}
