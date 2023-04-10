package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	src := []byte(`<?php
	$c = $_GET['name'];
 	echo "Name is".$c;
	`)

	var b bytes.Buffer

	par := php5.NewParser(src, "example.php")
	par.Parse()

	visitor := visitor.GoDumper{
		Writer: &b,
	}

	//
	rootNode := par.GetRootNode()
	rootNode.Walk(&visitor)

	f, err := os.Create("./parser/example.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Write the package declaration to the file
	_, err = fmt.Fprintf(f, "package parser\n\n")
	if err != nil {
		panic(err)
	}

	// Write the import statements to the file
	_, err = fmt.Fprintf(f, "type program *Root\n\n")
	if err != nil {
		panic(err)
	}

	// Write the code to the file
	_, err = fmt.Fprintf(f, "func Test() program {\n    p:= "+b.String()+"\n return p \n}\n")
	if err != nil {
		panic(err)
	}

}
