package main

import "fmt"

func main() {
	var i, j int = 1, 2

	//declaration without var; only work inside a function
	k := 3
	a, b, c := true, false, "no"
	fmt.Println(i, j, k, a, b, c)
}
