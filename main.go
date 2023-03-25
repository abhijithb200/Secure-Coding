package main

import (
	"os"

	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	src := []byte(`<?php
	function writeMsg() {
		echo "Hello world!";
	  }
	  
	  writeMsg(); // call the function
	`)

	parser := php5.NewParser(src, "example.php")
	parser.Parse()

	visitor := visitor.GoDumper{
		Writer: os.Stdout,
	}
	//
	rootNode := parser.GetRootNode()
	rootNode.Walk(&visitor)
	// parser.Parser()
}
