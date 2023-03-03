package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
