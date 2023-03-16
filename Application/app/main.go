package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	src := bytes.NewBufferString(`<? echo "Hello world";`)

	parser := php5.NewParser(src, "example.php")
	parser.Parse()

	for _, e := range parser.GetErrors() {
		fmt.Println(e)
	}

	visitor := visitor.Dumper{
		Writer:    os.Stdout,
		Indent:    "",
		Comments:  parser.GetComments(),
		Positions: parser.GetPositions(),
	}

	rootNode := parser.GetRootNode()
	rootNode.Walk(visitor)
}
