package main

import (
	"fmt"
	"time"
)

func f(c chan int) {
	if i, ok := <-c; ok {
		fmt.Println("Read number:", i)
	} else {
		fmt.Println("Channel closed")
	}
}

func main() {
	channel := make(chan int)
	go f(channel)

	close(channel)

	time.Sleep(time.Second*3)
}
