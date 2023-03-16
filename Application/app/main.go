package main

import "app/parser"

func main() {
	// src := []byte(`<?php
	// echo 'my name is'.$_GET['name'];
	// `)

	// parser := php5.NewParser(src, "example.php")
	// parser.Parse()

	// visitor := visitor.GoDumper{
	// 	Writer: os.Stdout,
	// }
	// //
	// rootNode := parser.GetRootNode()
	// rootNode.Walk(&visitor)
	parser.Parser()
}
