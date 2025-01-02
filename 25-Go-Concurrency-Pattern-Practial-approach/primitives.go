package main

import (
	"fmt"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	// go someFunc("123")
	// go someFunc("123")
	// go someFunc("123")

	// time.Sleep(time.Microsecond * 100)

	// fmt.Println("hi")

	// ----- Channels -----
	// Creating a Channel
	myChannel := make(chan string)
	myChannel2 := make(chan string)

	// Channels are used for communication and synchronization between goroutines.
	// An unbuffered channel (like make(chans string)) requires both a
	// sender and a receiver to be ready at the same time for a message to be sent or received.

	// Anonymous function which is also a go routine putting some message
	// in to the channel for other go routines to retrieve or whatever
	go func() {
		myChannel <- "data from channel: myChannel"
	}()

	// go func() {
	// 	myChannel <- "2nd data from channel myCHannel"
	// }()

	go func() {
		myChannel2 <- "data from channel: myChannel2"
	}()

	// main being one of the go routines accessing that message form the channel
	msg := <-myChannel // this is actually blocking in nature, main waits for the channel to close or receiving any message before moving forward
	fmt.Println(msg)

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromMyChannel2 := <-myChannel2:
		fmt.Println(msgFromMyChannel2)
	}
}
