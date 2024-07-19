# [SYNC](https://pkg.go.dev/sync)

The `sync` package in Go provides essential tools for managing concurrency and synchronizing access to shared resources in multi-threaded programs.

Here's a summary of key components and important things to learn:

## **1. Mutual Exclusion:**

* **Mutex (`sync.Mutex`):** A basic mutex (mutual exclusion lock) ensures that only one goroutine can access a critical section of code at a time. Use `Lock()` to acquire the lock, `Unlock()` to release it, and `TryLock()` to attempt locking without blocking.

* **RWMutex (`sync.RWMutex`):** Allows multiple goroutines to read concurrently but only one goroutine to write at a time. Use `RLock()` and `RUnlock()` for reading, and `Lock()` and `Unlock()` for writing.

## **2. Condition Variables:**

* **Cond (`sync.Cond`):** Used to signal goroutines waiting on a specific condition. A `Cond` is associated with a `Locker` (like a `Mutex`). Goroutines can `Wait()` on a condition, `Signal()` to wake up one waiting goroutine, or `Broadcast()` to wake up all waiting goroutines.

## **3. Synchronization Primitives:**

* **WaitGroup (`sync.WaitGroup`):** Tracks the number of active goroutines. Useful for ensuring that all goroutines complete before proceeding. Use `Add()`, `Done()`, and `Wait()`.

* **Once (`sync.Once`):** Guarantees that a function is executed only once, even if called concurrently from multiple goroutines.

## **4. Concurrent Data Structures:**

* **Map (`sync.Map`):** A concurrent-safe map implementation. Provides methods like `Load()`, `Store()`, `Delete()`, and atomic operations like `CompareAndSwap()`.

* **Pool (`sync.Pool`):** A pool of reusable objects. Useful for reducing object allocation overhead in scenarios where objects are frequently created and destroyed.

**Important Things to Learn:**

* **Concurrency Safety:** Understand the importance of using synchronization primitives to prevent data races and ensure thread safety when accessing shared resources.

* **Deadlocks:** Be aware of potential deadlocks when using locks and condition variables.

* **Performance Considerations:** Choose the appropriate synchronization mechanism based on your specific needs. Overuse of locks can lead to performance bottlenecks.

* **Alternatives:** Explore other concurrency patterns and libraries (e.g., channels) for more complex synchronization scenarios.



Remember that concurrency can be complex. Start with simple examples, gradually increase complexity, and thoroughly test your code to ensure correctness and efficiency.

---

## The `sync` package: A Guide to Concurrency in Go

The `sync` package is the cornerstone of concurrent programming in Go, providing primitives for synchronizing access to shared resources and coordinating goroutines. It offers tools for:

* **Mutual exclusion:** Preventing data races by ensuring only one goroutine can access a critical section at a time (e.g., `Mutex`, `RWMutex`).
* **Synchronization:** Managing the execution order of goroutines and signaling events (e.g., `WaitGroup`, `Cond`).
* **Data structures:** Providing thread-safe data structures (e.g., `Map`).

**Important Concepts:**

* **Data Races:** Occur when multiple goroutines access shared memory without proper synchronization, leading to unpredictable and incorrect behavior.
* **Critical Sections:** Sections of code that access shared resources, requiring mutual exclusion to avoid data races.
* **Synchronization Primitives:** Tools like `Mutex`, `WaitGroup`, and `Cond` that ensure proper synchronization and coordination among goroutines.

### Detailed Breakdown of Types in `sync`

**1. Mutex:**

* **Purpose:** Provides mutual exclusion for a critical section.
* **Methods:**
    * `Lock()`: Acquires the lock, blocking the goroutine if another goroutine holds it.
    * `Unlock()`: Releases the lock, allowing other goroutines to acquire it.
    * `TryLock()`: Attempts to acquire the lock without blocking. Returns `true` if successful, `false` otherwise.

**Example:**

```go
import "sync"

var mu sync.Mutex
var counter int

func increment() {
	mu.Lock() // Acquire the lock
	defer mu.Unlock() // Release the lock when the function exits

	counter++
}

func main() {
	for i := 0; i < 10; i++ {
		go increment()
	}
	// ... Wait for goroutines to finish
	fmt.Println(counter)
}
```

**2. RWMutex:**

* **Purpose:** Allows multiple goroutines to read a shared resource concurrently but restricts writes to a single goroutine at a time.
* **Methods:**
    * `RLock()`: Acquires a read lock, allowing multiple goroutines to read concurrently.
    * `RUnlock()`: Releases a read lock.
    * `Lock()`: Acquires a write lock, blocking read and write operations until released.
    * `Unlock()`: Releases a write lock.
    * `TryLock()`: Attempts to acquire a write lock without blocking.
    * `TryRLock()`: Attempts to acquire a read lock without blocking.
    * `RLocker()`: Returns a `Locker` interface that only allows read locks.

**Example:**

```go
import "sync"

var rwMutex sync.RWMutex
var data map[string]int

func readData(key string) int {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return data[key]
}

func writeData(key string, value int) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	data[key] = value
}

func main() {
	data = make(map[string]int)

	go readData("key1")
	go readData("key2")
	go writeData("key3", 10)
}
```

**3. WaitGroup:**

* **Purpose:** Coordinates the execution of multiple goroutines. Allows a goroutine to wait until a set of other goroutines have completed their work.
* **Methods:**
    * `Add(delta int)`: Increments the internal counter by `delta`. Typically used with a positive `delta` to indicate the number of goroutines that will be added to the group.
    * `Done()`: Decrements the counter by 1. Each goroutine that is part of the group should call `Done()` when it finishes its work.
    * `Wait()`: Blocks the calling goroutine until the counter reaches 0, indicating all goroutines in the group have finished.

**Example:**

```go
import "sync"
import "time"

var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done() // Signal completion
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	wg.Add(5) // Indicate 5 goroutines will be added
	for i := 0; i < 5; i++ {
		go worker(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers finished")
}
```

**4. Once:**

* **Purpose:** Ensures a function is executed only once, even if multiple goroutines attempt to call it concurrently.
* **Methods:**
    * `Do(f func())`: Executes the function `f` only once. Subsequent calls to `Do()` with the same function will have no effect.

**Example:**

```go
import "sync"
import "fmt"

var once sync.Once
var data int

func initDatabase() {
	fmt.Println("Initializing database...")
	data = 10
}

func getData() int {
	once.Do(initDatabase) // Initialize database only once
	return data
}

func main() {
	for i := 0; i < 10; i++ {
		go getData()
	}
	// ...
}
```

**5. Cond:**

* **Purpose:** Provides a mechanism for goroutines to wait for a condition to be met and for other goroutines to signal when the condition has changed.
* **Methods:**
    * `NewCond(l Locker)`: Creates a new `Cond` associated with a given `Locker`. The `Locker` is used to protect the shared condition variable.
    * `Wait()`: Blocks the calling goroutine until it is signaled by another goroutine using `Signal()` or `Broadcast()`.
    * `Signal()`: Wakes up a single goroutine waiting on the condition.
    * `Broadcast()`: Wakes up all goroutines waiting on the condition.

**Example:**

```go
import "sync"
import "fmt"

var mu sync.Mutex
var cond *sync.Cond
var buffer []int

func init() {
	mu.Lock()
	cond = sync.NewCond(&mu)
	mu.Unlock()
}

func producer(value int) {
	mu.Lock()
	buffer = append(buffer, value)
	fmt.Printf("Producer: %d\n", value)
	cond.Signal() // Signal consumer
	mu.Unlock()
}

func consumer() {
	mu.Lock()
	for len(buffer) == 0 {
		fmt.Println("Consumer: waiting...")
		cond.Wait() // Wait for a signal
	}
	value := buffer[0]
	buffer = buffer[1:]
	fmt.Printf("Consumer: %d\n", value)
	mu.Unlock()
}

func main() {
	go producer(1)
	go consumer()
	// ...
}
```

**6. Pool:**

* **Purpose:** Manages a pool of objects to reduce the overhead of creating and destroying objects frequently.
* **Methods:**
    * `Get()`: Returns an object from the pool. If the pool is empty, it creates a new object.
    * `Put(x any)`: Returns an object back to the pool.

**Example:**

```go
import "sync"

var pool sync.Pool

func init() {
	pool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func worker() {
	data := pool.Get().([]byte)
	// ... use data ...
	pool.Put(data)
}

func main() {
	for i := 0; i < 10; i++ {
		go worker()
	}
	// ...
}
```

**7. Map:**

* **Purpose:** Provides a thread-safe map implementation.
* **Methods:**
    * `Store(key, value any)`: Stores a value associated with a key.
    * `Load(key any)`: Retrieves the value associated with a key.
    * `Delete(key any)`: Deletes the value associated with a key.
    * `LoadAndDelete(key any)`: Retrieves and deletes the value associated with a key.
    * `LoadOrStore(key, value any)`: Retrieves the value associated with a key, or stores a new value if it doesn't exist.
    * `Swap(key, value any)`: Replaces the value associated with a key and returns the previous value.
    * `CompareAndSwap(key, old, new any)`: Atomically replaces the value associated with a key if it matches `old`, and returns `true` if successful.
    * `CompareAndDelete(key, old any)`: Atomically deletes the value associated with a key if it matches `old`, and returns `true` if successful.

---

Let's break down the `sync` package in Go, exploring its key types and providing examples to illustrate their usage.

**Summary:**

The `sync` package is your toolbox for managing concurrency in Go. It provides tools to:

* **Protect shared data:** Prevent data races when multiple goroutines access the same data.
* **Coordinate goroutines:** Ensure they execute in a desired order or wait for each other.
* **Manage resources:** Efficiently share and reuse objects.

**Important Concepts:**

* **Concurrency Safety:**  When multiple goroutines access shared data, you need mechanisms to prevent unexpected behavior (data races). The `sync` package provides these mechanisms.
* **Deadlocks:** Be cautious when using locks. If goroutines hold locks that other goroutines need, you can create deadlocks (a standstill).

**Detailed Breakdown:**

1. **Mutex (`sync.Mutex`)**

   * **Purpose:** A basic lock. Only one goroutine can hold the lock at a time.

   * **Example:**

     ```go
     var counter int
     var mu sync.Mutex

     func increment() {
         mu.Lock() // Acquire the lock
         defer mu.Unlock() // Release the lock when done
         counter++
     }
     ```

2. **RWMutex (`sync.RWMutex`)**

   * **Purpose:** Allows multiple goroutines to read concurrently but only one to write at a time.

   * **Example:**

     ```go
     var data []int
     var rwmu sync.RWMutex

     func readData() int {
         rwmu.RLock() // Acquire read lock
         defer rwmu.RUnlock()
         return data[0]
     }

     func writeData(value int) {
         rwmu.Lock() // Acquire write lock
         defer rwmu.Unlock()
         data[0] = value
     }
     ```

3. **Cond (`sync.Cond`)**

   * **Purpose:** Allows goroutines to wait for a specific condition to become true.

   * **Example:**

     ```go
     var cond sync.Cond
     var ready bool

     func worker() {
         cond.L.Lock()
         for !ready {
             cond.Wait() // Wait until 'ready' is true
         }
         cond.L.Unlock()
         // ... do work ...
     }

     func signalWorker() {
         cond.L.Lock()
         ready = true
         cond.Signal() // Wake up one waiting goroutine
         cond.L.Unlock()
     }
     ```

4. **WaitGroup (`sync.WaitGroup`)**

   * **Purpose:** Tracks the number of active goroutines. Useful for ensuring all goroutines finish before proceeding.

   * **Example:**

     ```go
     var wg sync.WaitGroup
     for i := 0; i < 5; i++ {
         wg.Add(1) // Register a goroutine
         go func(id int) {
             defer wg.Done() // Signal completion
             // ... do work ...
         }(i)
     }
     wg.Wait() // Wait for all goroutines to finish
     ```

5. **Once (`sync.Once`)**

   * **Purpose:** Ensures a function is executed only once, even if called concurrently.

   * **Example:**

     ```go
     var once sync.Once
     func doSomething() {
         fmt.Println("Doing something once")
     }

     func main() {
         once.Do(doSomething) // Will execute 'doSomething' only once
         once.Do(doSomething) // Subsequent calls have no effect
     }
     ```

6. **Map (`sync.Map`)**

   * **Purpose:** A concurrent-safe map.

   * **Example:**

     ```go
     var data sync.Map

     func set(key, value string) {
         data.Store(key, value)
     }

     func get(key string) (string, bool) {
         return data.Load(key)
     }
     ```

7. **Pool (`sync.Pool`)**

   * **Purpose:** A pool of reusable objects.

   * **Example:**

     ```go
     var bufferPool sync.Pool

     func getBuffer() []byte {
         return bufferPool.Get().([]byte)
     }

     func putBuffer(buf []byte) {
         bufferPool.Put(buf)
     }
     ```

**Key Takeaways:**

* Choose the right tool for the job. Each type in the `sync` package serves a specific purpose.
* Understand the potential for deadlocks and take steps to avoid them.
* Test your concurrent code thoroughly to ensure correctness.