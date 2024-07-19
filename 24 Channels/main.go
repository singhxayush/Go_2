package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)

	// Recieve Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		if !isChannelOpen {
			fmt.Println("Value Received from closed channel")
		}

		val, isChannelOpen = <-myCh
		fmt.Println(val)

		if !isChannelOpen {
			fmt.Println("Value Received from closed channel")
		}

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// Send Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer close(myCh)
		myCh <- 2
		myCh <- 3
		wg.Done()
	}(myCh, wg)
}
