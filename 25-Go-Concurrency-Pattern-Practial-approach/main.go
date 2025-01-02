package main

import (
	"fmt"
	"time"
)

func f(c chan int) {
	fmt.Println("Inside go routne function")

	fmt.Printf("Value received from the channel is %v", <- c)
}

func main() {
	var channel chan int = make(chan int, 1)

	go f(channel)

	channel <- 100

	time.Sleep(time.Second * 3)

}

