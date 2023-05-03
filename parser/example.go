package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "servername",
					},
				},
				Expression: &ArrayDimFetch{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "_GET",
						},
					},
					Dim: &String{
						Value: "\"servername\"",
					},
				},
			},
		},
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "username",
					},
				},
				Expression: &ArrayDimFetch{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "_GET",
						},
					},
					Dim: &String{
						Value: "\"username\"",
					},
				},
			},
		},
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "password",
					},
				},
				Expression: &ArrayDimFetch{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "_POST",
						},
					},
					Dim: &String{
						Value: "\"password\"",
					},
				},
			},
		},
	},
}

 return p 
}
