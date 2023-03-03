package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() { // the receiving goroutine
		i := <-ch // pulling data from the channel and assigning to variabl i
		fmt.Println(i)
		wg.Done()
	}()

	go func() { // the sending goroutine
		ch <- 42 // putting data into the channel
		wg.Done()
	}()

	wg.Wait()

}
