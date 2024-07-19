## "Go Routine" and "Main"

```go
package main

import "fmt"

func main() {
	go greeter("Go Routine")
	greeter("Main")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
```
```text
Output:

Main
Main
Main
Main
Main
```

### Layman's Explanation

Imagine you have two friends, let's call them "Go Routine" and "Main".

You tell "Go Routine" to say "Main" five times, but you also tell them to start saying it *later*.

Meanwhile, you tell "Main" to say "Go Routine" five times right away.

Since "Main" starts talking first and doesn't wait for "Go Routine", you only hear "Go Routine" repeated five times.

"Go Routine" is still saying "Main" in the background, but you can't hear it because "Main" is louder.

### In-Depth Explanation

The code demonstrates the concept of **goroutines** in Go.

* **Goroutines:** These are lightweight, concurrent functions that run independently.

* **`go greeter("Main")`:** This line starts a new goroutine that executes the `greeter` function with the argument "Main".

* **`greeter("Go Routine")`:** This line calls the `greeter` function directly in the main goroutine, passing "Go Routine" as the argument.

Because the `greeter` function is called directly in the main goroutine, its output ("Go Routine" five times) is printed immediately.

The goroutine started with `go greeter("Main")` runs concurrently, but its output ("Main" five times) is not printed until the main goroutine finishes.

**Why doesn't the output interleave?**

Go's default scheduling behavior doesn't guarantee interleaving of goroutine output. The runtime scheduler might prioritize the main goroutine, leading to the observed output.

**To see interleaved output, you would need to use synchronization mechanisms like channels or mutexes to control the execution flow between goroutines.**

### Solution:
To see both `Main` and `Go Routine` printing, you need to synchronize the execution.  Here are a few ways to do that:

* **time.Sleep:**  You could add a short delay in the main goroutine using ⁠time.Sleep before calling ⁠greeter("Go Routine").  This would give the ⁠greeter("Main") goroutine some time to print its messages.

* **Channels:**  Channels allow goroutines to communicate with each other. You could use a channel to signal the main goroutine that ⁠greeter("Main") has finished printing.

Visit the Go website: [Ayush](https://go.dev/)