# how to use a single channel in Go?

In the previous example, both channels are used. But what if you want one of them? In Go we have two ways: using `select` and using a
channel with only one channel.

## Using select

With `select`, you can control what channel should receive messages. Let's say that you have three channels in which you want to receive
data, but not all at the same time. You could do this by using the `select` statement:

```go
	ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    go func() {
        for i := 0; ; i++ {
            select {
                case <-ch1:
                    println("Received on channel #1")
                case <-ch2:
                    println("Received on channel #2")
                case <-ch3:
                    println("Received on channel #3")
            }
        }
    }()
```

The `select` statement makes Go wait for the first channel that has data, if all of them have data then it blocks until one of them
receives a message. The value that was sent will be received by the channel that received it.

You don't need to create a channel with every one you want, just call `select` and then each of your channels inside the `select`. If
all channels are empty, Go waits indefinitely until something is received on any of them (this might cause a deadlock if there's no
receiver for that channel).

## Using a channel with only one channel.

The second way to use multiple channels is by using a channel with two elements. The first element will be used as a signal and the
second element will be used as data. This way, we can avoid creating a channel with every single one of them that are not going to
receive any messages.

```go
	ch := make(chan struct{})

    go func() {
        for i := 0; ; i++ {
            select {
                case <-ch:
                    println("received")
                default:
					// send a message
                    ch <- struct{}{}
            }
        }
    }()
```

In the example above, we have two channels. The first is a signal channel that will be used for sending messages. We create it using
`make(chan struct{})`. When we want to send something on this channel, then we call `<-ch`, which sends a message onto the channel. We
can also use it as a wait group. Let's say that we have 10 channels and each one of them should wait for a message to send back. We
create an infinite loop and when there is no data in any of these channels, then we signal those 10 channels. This way they will all
receive messages from the channel which was used as a signal.

---

## There are 2 Types of channels

1. Unbuffered
2. Buffered

## Buffered vs unbuffered channels. which to use and when?

a channel is a communication mechanism. buffered channels allow a limited number of messages to be stored in the channel. a buffered
channel with 5 message slots, can hold 3 messages, if 4th message comes while one slot is already used. buffered channel waits until
some slot becomes available for sending another message to the channel and then sends it. unbuffered channels do not wait but fail if
there are no available slots.
buffered channels:

- they allow us to limit number of messages in a channel.
- when you have large amount of messages, buffered channel makes sense.
- buffered channel is useful for situations where you want to control how many messages you send through the channel at one time and it
  also ensures that you don't wait until some slot becomes available.
  unbuffered channels:
- they are useful in a situation where you are sending message to the channel from different goroutines but you want to ensure that a
  certain number of messages is sent before any go routines block on recieving from unbuffered channel.

## Channel with no receiving end

```go
package main

import "fmt"

func main() {
	// This code has 2 cases, one will cause dead lock in case 1 where the channel has no receiving end. It is because channels donot hold data. They only act as the medium to transfer data between goroutines simultaneously. Now in this case when data is sent through the channel, we are under same groutine and there is no other goroutine to receive the data. So, the main goroutine will be blocked and will cause deadlock. On the other hand in case 2 we have a goroutine to receive the data from the channel. So, the main goroutine will not be blocked and the program will run successfully.

	dataChan := make(chan int)

	// Case 1: This will block the main goroutine
	// OP: fatal error: all goroutines are asleep - deadlock!
	dataChan <- 42

	// Case 2: This will block the main goroutine
	// OP: 42
	go func() {
		dataChan <- 42
	}()

	n := <-dataChan

	fmt.Println("n =", n)
}

```
