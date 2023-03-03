package main

import (
	"fmt"
)

func main() {
	var a, b int = 3, 4
	var c float32 = 3.13
	d := int(c) // var d int = int(c)

	fmt.Println(a, b, c, d)
}
