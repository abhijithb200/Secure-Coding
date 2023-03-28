package main

import "codeguardian/parser"

func main() {
	// src := []byte(`<?php
	// $a = $_GET['name'];
	// $b = "Name is ".$a;

	// echo "I am ".$b;
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
