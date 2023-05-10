package main

import (
	"codeguardian/parser"
	"os"
)

var hash string = ""

func main() {
	args := os.Args
	
	if len(args)>1 {
		hash = args[1]
	}
	parser.Parser(hash)
}
