package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Received from the done channel, terminating the goroutine")
			return
		default:
			fmt.Println("Doing work...")
		}
	}
}

func main() {

	done := make(chan bool)
	/*
			  The done channel is commonly used for signaling a goroutine to stop, especially when:
		    1. You don't know when exactly you'll need to stop the goroutine.
		    2. You want to cancel a long-running or blocking operation.
		    3. You want to prevent resource leaks by ensuring that goroutines can be terminated when no longer needed.
	*/

	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done)

}
