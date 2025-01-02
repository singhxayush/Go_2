package main

import "fmt"

func main() {

	charChannel := make(chan string, 3) // buffered channel (fixed in size)
	// communication between go routines is synchronous when we are using unbuffered channels
	// to make it asynchronous buffered channels are required
	// and that is because an unbuffered channel provides a guarantee that an exchange
	// between 2 goroutines is performed at the instant the send and receive takes place

	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for res := range charChannel {
		fmt.Println(res)
	}
}
