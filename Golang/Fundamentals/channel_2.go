package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) { // receive only - data flow out of channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // send only - data flow into the channel
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()

}
