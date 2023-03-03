package main

import (
	"fmt"
)

var (
	a bool   = false
	b uint64 = 1<<64 - 1
	c string = "abhi"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %q\n", c, c)
}

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128
