package main

import (
	"fmt"
)

func main() {
	statePopulation := make(map[string]int){
		"California": 200,
		"Texas":      300,
		"New York":   400,
	}
	fmt.Println(statePopulation)
}
