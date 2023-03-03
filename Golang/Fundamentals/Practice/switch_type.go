package main

import (
	"fmt"
)

func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("another type")
	}
}
