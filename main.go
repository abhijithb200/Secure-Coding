package main

import "codeguardian/parser"

func main() {
	// src := []byte(`<?php
	// echo 'my name is'.$_GET['name'];
	// echo 'my address is'.$_GET['address'];
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
