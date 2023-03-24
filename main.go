package main

import "codeguardian/parser"

func main() {
	// src := []byte(`<?php
	// $a = $_GET['name']."abhi";
	// echo 'my address is'.$a;
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
