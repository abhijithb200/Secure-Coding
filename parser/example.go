package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "heading",
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
		&Expression{
			Expr: &Assign{
				Variable: &Variable{
					VarName: &Identifier{
						Value: "name",
					},
				},
				Expression: &ArrayDimFetch{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "_GET",
						},
					},
					Dim: &String{
						Value: "'heading'",
					},
				},
			},
		},
		&InlineHtml{
			Value: "\r\n<!DOCTYPE html>\r\n<html>\r\n<head>\r\n  <title>",
		},
		&Echo{
			Exprs: []Node{
				&Variable{
					VarName: &Identifier{
						Value: "heading",
					},
				},
			},
		},
		&InlineHtml{
			Value: "</title>\r\n  <meta charset=\"UTF-8\">\r\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n  <style>\r\n    /* CSS styles go here */\r\n    body {\r\n      font-family: Arial, sans-serif;\r\n      font-size: 16px;\r\n      line-height: 1.5;\r\n      background-color: #f0f0f0;\r\n    }\r\n    h1 {\r\n      color: #333;\r\n      text-align: center;\r\n      margin-top: 50px;\r\n    }\r\n  </style>\r\n</head>\r\n<body>\r\n  <h1>Welcome to ",
		},
		&Echo{
			Exprs: []Node{
				&Variable{
					VarName: &Identifier{
						Value: "heading",
					},
				},
			},
		},
		&InlineHtml{
			Value: ", ",
		},
		&Echo{
			Exprs: []Node{
				&Variable{
					VarName: &Identifier{
						Value: "name",
					},
				},
			},
		},
		&InlineHtml{
			Value: "!</h1>\r\n  <p>This is a sample webpage built using PHP and HTML.</p>\r\n</body>\r\n</html>\r\n",
		},
	},
}

 return p 
}
