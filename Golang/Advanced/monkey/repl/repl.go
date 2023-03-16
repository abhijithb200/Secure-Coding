package repl

import (
	"bufio"
	"fmt"

	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
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
		p := parser.New(l)

		program := p.ParseProgram()

		evaluated := evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}

}
