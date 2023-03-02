package main

import (
	"fmt"
)

func main() {
	var w Writer = age{}
	w.Write([]byte("Hello"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type age struct {
	name string
}

func (a age) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	fmt.Println(a.name)
	return n, err
}
