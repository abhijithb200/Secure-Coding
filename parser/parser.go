package parser

import (
	"fmt"
)

var a string = "helo"

// var b *string = &a

func Parser() {
	// var strBuffer bytes.Buffer

	program := Test()

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
