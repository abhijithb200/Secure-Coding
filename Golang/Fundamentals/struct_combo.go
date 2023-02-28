package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	Canfly bool
}

func main() {
	b := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		Speed:  30,
		Canfly: false,
	}

	fmt.Println(b.Name)
}
