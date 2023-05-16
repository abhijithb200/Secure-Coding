package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
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
							Value: "_POST",
						},
					},
					Dim: &String{
						Value: "'username'",
					},
				},
			},
		},
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "sql",
					},
				},
				Expression: &Encapsed{
					Parts: []Node{
						&EncapsedStringPart{
							Value: "SELECT * FROM users WHERE username = ",
						},
						&Variable{
							VarName: &Identifier{
								Value: "username",
							},
						},
					},
				},
			},
		},
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "result",
					},
				},
				Expression: &FunctionCall{
					Function: &Name{
						Parts: []Node{
							&NamePart{
								Value: "mysqli_query",
							},
						},
					},
					ArgumentList: &ArgumentList{
						Arguments: []Node{
							&Argument{
								Variadic: false,
								IsReference: false,
								Expr: &Variable{
									VarName: &Identifier{
										Value: "conn",
									},
								},
							},
							&Argument{
								Variadic: false,
								IsReference: false,
								Expr: &Variable{
									VarName: &Identifier{
										Value: "sql",
									},
								},
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
