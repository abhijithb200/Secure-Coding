package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan() // scan the input

		if !scanned {
			return
		}

		line := scanner.Text() // the text user typed
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { //for k := 0; k < 4; k++
			fmt.Printf("%+v\n", tok) // print the fields in the struct
		}
	}

}
