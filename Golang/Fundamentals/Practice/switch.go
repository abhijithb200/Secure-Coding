package main

import (
	"fmt"
)

func main() {
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
