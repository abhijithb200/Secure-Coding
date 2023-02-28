package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John",
		companions: []string{
			"Abhi",
			"Akhil",
		},
	}

	fmt.Println(aDoctor)
}
