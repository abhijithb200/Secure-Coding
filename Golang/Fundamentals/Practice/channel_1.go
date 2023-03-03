package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch      // 2
		fmt.Println(i) // 3
		ch <- 27       // 4
		wg.Done()
	}()

	go func() {
		ch <- 42          // 1
		fmt.Println(<-ch) // 5
		wg.Done()
	}()

	wg.Wait()

}
