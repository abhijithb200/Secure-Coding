package parser

type program *Root

func Test() program {
    p:= &Root{
	Stmts: []Node{
		&Echo{
			Exprs: []Node{
				&String{
					Value: "''",
				},
			},
		},
	},
}

 return p 
}
