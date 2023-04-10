package main

import (
	"encoding/json"
	"fmt"
)

type Node interface {
}

type Root struct {
	Stmts []Node
}

type Exprs struct {
	Exprs []Node
}

type Concat struct {
	Left  Node
	Right Node
}

type String struct {
	Value string
}

func main() {

	// parser.Parser()

	p := `
	{
	"Stmts":[
		{
			"Exprs":[
				{
					"Left":{
						"Value":"\"abhi\""
					},
					"Right":{
						"Value":"\"ypooo\""
					}
				}
					]
		}
			]
	}
	`
	var r Root
	json.Unmarshal([]byte(p), &r)

	fmt.Println(r.Stmts["Right"])
}
