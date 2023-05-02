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
				Expression: &String{
					Value: "\"localhost\"",
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
				Expression: &String{
					Value: "\"username\"",
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
				Expression: &String{
					Value: "\"password\"",
				},
			},
		},
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "conn",
					},
				},
				Expression: &FunctionCall{
					Function: &Name{
						Parts: []Node{
							&NamePart{
								Value: "mysqli_connect",
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
										Value: "servername",
									},
								},
							},
							&Argument{
								Variadic: false,
								IsReference: false,
								Expr: &Variable{
									VarName: &Identifier{
										Value: "username",
									},
								},
							},
							&Argument{
								Variadic: false,
								IsReference: false,
								Expr: &Variable{
									VarName: &Identifier{
										Value: "password",
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
