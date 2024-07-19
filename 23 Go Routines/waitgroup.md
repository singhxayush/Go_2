### [WAIT GROUP](https://pkg.go.dev/sync#WaitGroup)

Let's break down these three WaitGroup methods in simpler terms:

Imagine you have a group of workers (goroutines) who need to complete tasks before you can move on to the next step in your program. A `WaitGroup` helps you coordinate these workers and ensure they're all finished before proceeding.

**1. `wg.Add(delta int)`:**

* **Purpose:** Tells the `WaitGroup` that you're adding `delta` number of workers to the group.

* **Analogy:** Think of it like registering new workers with a foreman. You tell the foreman how many workers are joining the team.

* **Example:**

```go
var wg sync.WaitGroup
wg.Add(3) // Registering 3 workers
```

**2. `wg.Done()`:**

* **Purpose:** Called by a worker when it has finished its task. It decrements the internal counter of the `WaitGroup`, signaling that one worker is done.

* **Analogy:** A worker finishes their job and informs the foreman they're done.

* **Example:**

```go
go func() {
    defer wg.Done() // Worker signals completion
    // ... do some work ...
}()
```

**3. `wg.Wait()`:**

* **Purpose:** Blocks the main goroutine (the one that called `Wait()`) until all workers in the group have called `wg.Done()`.

* **Analogy:** The foreman waits until all registered workers have reported back as finished.

* **Example:**

```go
wg.Wait() // Main goroutine waits for all workers
fmt.Println("All workers finished!")
```

**In Summary:**

* `Add()` registers workers.
* `Done()` signals worker completion.
* `Wait()` blocks until all workers are done.

**Key Points:**

* `WaitGroup` is essential for synchronizing goroutines and ensuring they complete before moving on.
* `Done()` should be called *after* a worker's task is finished, typically using `defer wg.Done()`.
* `Wait()` is usually called in the main goroutine to wait for all worker goroutines.

---
Let's explore the use cases of `sync.WaitGroup` in Go, along with some advanced examples and in-depth explanations.

**Fundamental Use Case: Ensuring Goroutine Completion**

The most common use of `WaitGroup` is to guarantee that a set of goroutines completes their tasks before the main goroutine proceeds.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers finished!")
}
```

In this example, the main goroutine launches five worker goroutines. `wg.Add(1)` registers each worker, and `defer wg.Done()` ensures that the `WaitGroup` is decremented when each worker finishes. `wg.Wait()` blocks the main goroutine until all workers have signaled completion.

**Advanced Use Cases:**

1. **Parallel Data Processing:**

   Imagine processing a large dataset in parallel. You can use `WaitGroup` to coordinate worker goroutines that process chunks of data.

   ```go
   func processChunk(data []int, wg *sync.WaitGroup) {
       defer wg.Done()
       // ... process data chunk ...
   }

   func main() {
       var wg sync.WaitGroup
       data := generateData(10000) // Generate a large dataset
       chunkSize := 1000
       for i := 0; i < len(data); i += chunkSize {
           wg.Add(1)
           go processChunk(data[i:i+chunkSize], &wg)
       }
       wg.Wait()
       // ... combine results from processed chunks ...
   }
   ```

2. **Resource Management:**

   `WaitGroup` can help manage shared resources accessed by goroutines. For example, you might use it to ensure that all goroutines have finished using a database connection pool before closing it.

3. **Graceful Shutdown:**

   In a server application, you can use `WaitGroup` to gracefully shut down worker goroutines when a shutdown signal is received.

   ```go
   var wg sync.WaitGroup
   // ... start worker goroutines ...

   func shutdown() {
       // Signal workers to stop
       // ...

       wg.Wait() // Wait for all workers to exit
       // ... close resources ...
   }
   ```

4. **Pipeline Processing:**

   `WaitGroup` can be used to synchronize stages in a pipeline of goroutines, where the output of one stage becomes the input of the next.

**Important Considerations:**

* **Error Handling:**

   In real-world scenarios, you'll often need to handle errors within goroutines. Consider using channels or other mechanisms to propagate errors back to the main goroutine.

* **Deadlocks:**

   Be careful to avoid deadlocks when using `WaitGroup`. Ensure that goroutines don't wait indefinitely for each other.

* **Alternatives:**

   For more complex synchronization needs, consider using channels or other concurrency primitives.

Remember that `WaitGroup` is a powerful tool for coordinating goroutines, but it's important to use it thoughtfully and understand its limitations.