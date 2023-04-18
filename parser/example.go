package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "file",
					},
				},
				Expression: &ArrayDimFetch{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "_GET",
						},
					},
					Dim: &String{
						Value: "'file'",
					},
				},
			},
		},
		&Expression{
			Expr: &Include{
				Expr: &Encapsed{
					Parts: []Node{
						&EncapsedStringPart{
							Value: "pages/",
						},
						&Variable{
							VarName: &Identifier{
								Value: "file",
							},
						},
					},
				},
			},
		},
	},
}

 return p 
}
