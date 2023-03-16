package parser

import "fmt"

var a string = "helo"

// var b *string = &a

func Parser() {
	program := &Root{
		Stmts: []Node{
			&Echo{
				Exprs: []Node{
					&Concat{
						Left: &String{
							Value: "'my name is'",
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

	for _, r := range program.Stmts {

		r.Out()

		fmt.Println(a)
		// switch r.(type) Println()
		// case *Echo:
		// 	// s := r.(*Echo).Exprs[0]
		// 	// m := s.(*String)
		// 	// fmt.Println(m.Value)
		// 	r.Out()

		// case *Expression:
		// 	r.Out()

		// default:
		// 	fmt.Println("no")
		// }

	}
}
