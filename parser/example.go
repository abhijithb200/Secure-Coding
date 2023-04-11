package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
		&Function{
			ReturnsRef: false,
			PhpDocComment: "",
			FunctionName: &Identifier{
				Value: "writeMsg",
			},
			Params: []Node{
				&Parameter{
					ByRef: false,
					Variadic: false,
					Variable: &Variable{
						VarName: &Identifier{
							Value: "c",
						},
					},
				},
			},
			Stmts: []Node{
				&Echo{
					Exprs: []Node{
						&Concat{
							Left: &String{
								Value: "\"Hello world!\"",
							},
							Right: &Variable{
								VarName: &Identifier{
									Value: "c",
								},
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
						Value: "\"name\"",
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
				ArgumentList: &ArgumentList{
					Arguments: []Node{
						&Argument{
							Variadic: false,
							IsReference: false,
							Expr: &Variable{
								VarName: &Identifier{
									Value: "a",
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
