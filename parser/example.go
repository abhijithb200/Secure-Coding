package parser

type program *Root

func Test() program {
	p := &Root{
		Stmts: []Node{
			&Expression{
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "host",
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
						Value: "\"db_user\"",
					},
				},
			},
			&Expression{
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "passwd",
						},
					},
					Expression: &String{
						Value: "\".mypwd\"",
					},
				},
			},
			&Expression{
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "dbname",
						},
					},
					Expression: &String{
						Value: "\"my_db\"",
					},
				},
			},
			&Expression{
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "con",
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
									Variadic:    false,
									IsReference: false,
									Expr: &Variable{
										VarName: &Identifier{
											Value: "host",
										},
									},
								},
								&Argument{
									IsReference: false,
									Variadic:    false,
									Expr: &Variable{
										VarName: &Identifier{
											Value: "username",
										},
									},
								},
								&Argument{
									Variadic:    false,
									IsReference: false,
									Expr: &Variable{
										VarName: &Identifier{
											Value: "passwd",
										},
									},
								},
								&Argument{
									Variadic:    false,
									IsReference: false,
									Expr: &Variable{
										VarName: &Identifier{
											Value: "dbname",
										},
									},
								},
							},
						},
					},
				},
			},
			&If{
				Cond: &Variable{
					VarName: &Identifier{
						Value: "con",
					},
				},
				Stmt: &StmtList{
					Stmts: []Node{
						&Expression{
							Expr: &Print{
								Expr: &String{
									Value: "\"Connection Established Successfully\"",
								},
							},
						},
					},
				},
				Else: &Else{
					Stmt: &StmtList{
						Stmts: []Node{
							&Expression{
								Expr: &Print{
									Expr: &String{
										Value: "\"Connection Failed \"",
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
							Value: "sql",
						},
					},
					Expression: &String{
						Value: "\"SELECT name FROM user\"",
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
									Variadic:    false,
									IsReference: false,
									Expr: &Variable{
										VarName: &Identifier{
											Value: "con",
										},
									},
								},
								&Argument{
									Variadic:    false,
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
			&If{
				Cond: &Greater{
					Left: &PropertyFetch{
						Variable: &Variable{
							VarName: &Identifier{
								Value: "result",
							},
						},
						Property: &Identifier{
							Value: "num_rows",
						},
					},
					Right: &Lnumber{
						Value: "0",
					},
				},
				Stmt: &StmtList{
					Stmts: []Node{
						&While{
							Cond: &Assign{
								Variable: &Variable{
									VarName: &Identifier{
										Value: "row",
									},
								},
								Expression: &MethodCall{
									Variable: &Variable{
										VarName: &Identifier{
											Value: "result",
										},
									},
									Method: &Identifier{
										Value: "fetch_assoc",
									},
									ArgumentList: &ArgumentList{},
								},
							},
							Stmt: &StmtList{
								Stmts: []Node{
									&Echo{
										Exprs: []Node{
											&Concat{
												Left: &ArrayDimFetch{
													Variable: &Variable{
														VarName: &Identifier{
															Value: "row",
														},
													},
													Dim: &String{
														Value: "'name'",
													},
												},
												Right: &String{
													Value: "\"<br>\"",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Else: &Else{
					Stmt: &StmtList{
						Stmts: []Node{
							&Echo{
								Exprs: []Node{
									&String{
										Value: "\"0 results\"",
									},
								},
							},
						},
					},
				},
			},
			&Expression{
				Expr: &MethodCall{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "con",
						},
					},
					Method: &Identifier{
						Value: "close",
					},
					ArgumentList: &ArgumentList{},
				},
			},
			&InlineHtml{
				Value: "\t",
			},
		},
	}

	return p
}
