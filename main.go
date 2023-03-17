package main

import "codeguardian/parser"

func main() {
	// src := []byte(`<?php
	// $p = $_GET['name'];
	// echo 'my name is'.p;
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
