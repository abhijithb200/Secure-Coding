package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // the buffer can store 50 integers
	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch // receiving 42 and prints
		fmt.Println(i)

		i = <-ch // receiving 27 and prints
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	wg.Wait()

}
