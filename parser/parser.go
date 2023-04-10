package parser

import "fmt"

type TaintSpec struct {
	alias string
	spec  Values

	scope string
}

type VulnReport struct {
	name     string
	message  string
	position Position

	some Values

	taintvar map[string]TaintSpec
}

type FuncContainer struct {
	pos  int
	evil string
}

var VulnTracker *VulnReport = &VulnReport{}
var StmtsNew = map[string]Node{}

var CurrFuncStatus = map[string][]struct {
	pos  int
	evil string
}{}

func Parser() {
	VulnTracker.taintvar = make(map[string]TaintSpec)

	program := &Root{
		Stmts: []Node{
			&Function{
				ReturnsRef:    false,
				PhpDocComment: "",
				FunctionName: &Identifier{
					Value: "writeMsg",
				},
				Params: []Node{
					&Parameter{
						Variadic: false,
						ByRef:    false,
						Variable: &Variable{
							VarName: &Identifier{
								Value: "zee",
							},
						},
					},
					&Parameter{
						Variadic: false,
						ByRef:    false,
						Variable: &Variable{
							VarName: &Identifier{
								Value: "zoo",
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
										Value: "zee",
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
				Expr: &Assign{
					Variable: &Variable{
						VarName: &Identifier{
							Value: "b",
						},
					},
					Expression: &String{
						Value: "\"Abhijith\"",
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
								Variadic:    false,
								IsReference: false,
								Expr: &Variable{
									VarName: &Identifier{
										Value: "a",
									},
								},
							},
							&Argument{
								Variadic:    false,
								IsReference: false,
								Expr: &Variable{
									VarName: &Identifier{
										Value: "b",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, r := range program.Stmts {

		switch r.(type) {
		case *Function:
			currFuncName := r.Out(ArgStore{
				from: "Disable",
			}).(FunctionNew).FunctionName

			StmtsNew[currFuncName] = r

		case *Expression:
			r.Out(ArgStore{})
		}
	}
	fmt.Println(VulnTracker.taintvar)
}
