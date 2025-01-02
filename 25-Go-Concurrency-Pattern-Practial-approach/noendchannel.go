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