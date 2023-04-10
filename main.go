package main

import "codeguardian/parser"

func main() {
	// src := []byte(`<?php
	// 	function writeMsg($zee,$zoo) {
	// 	  echo "Hello world!".$zee;
	// 	}
	// 	$a = $_GET["name"];
	// 	$b = "Abhijith";
	// 	writeMsg($a,$b);
	// 	`)

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
