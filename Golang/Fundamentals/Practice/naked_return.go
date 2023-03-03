package main

import "fmt"

func split() (x, y int) {
	x = 100 + 1
	y = 100 - 1

	//naked return
	return
}

func main() {
	fmt.Println(split())
}
