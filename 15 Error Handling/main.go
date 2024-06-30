package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("------------------\n| Error Handling |\n------------------\n\n")

	//-> In Go, we use defer, panic and recover statements to handle errors.

	// Defer
	// DeferExample1()
	// DeferExample2()

	// Panic
	// PanicExample()

	// Recover
	RecoverExample()
}

func DeferExample1() {
	// Note: We use defer to delay the execution of functions that might cause an error.
	// defer the execution of Println() function
	defer fmt.Println("Three-1")

	fmt.Println("One-1")
	fmt.Println("Two-1")

	// Multiple defer Statements in Go
	// Note: When we use multiple defer in a program, the order of execution of the defer statements will be LIFO (Last In First Out).
	defer fmt.Println("One-2")
	defer fmt.Println("Two-2")
	defer fmt.Println("Three-2")
}

func DeferExample2() {
	f := createFile("/tmp/defer.txt")
	// defer the closing of the file
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	// if there's any error during creation, panic
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// if there's any error during closing, print the error and exit
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func PanicExample() {
	// Note: We use the panic statement to immediately end the execution of the program. If our program reaches a point where it cannot be recovered due to some major errors, it's best to use panic.
	fmt.Println("Help! Something bad is happening.")
	panic("Ending the program")
	fmt.Println("Waiting to execute") // will never get though
}

/*
	-> Output of PanicExample():
	Help! Something bad is happening.
	panic: Ending the program

	goroutine 1 [running]:
	main.PanicExample()
	/Users/ayushkumar/programming/Go_2/15 Error Handling/main.go:70 +0x68
	main.main()
	/Users/ayushkumar/programming/Go_2/15 Error Handling/main.go:18 +0x4c
	exit status 2

	-> Explanation
	You're looking at a stack trace, a crucial debugging tool in Go (and many other programming languages).
	Here's a breakdown of what it means:

		1. goroutine 1 [running]: This indicates that the panic occurred within goroutine number 1. Goroutines are lightweight, concurrent execution units in Go.
		2. main.PanicExample(): This shows the function where the panic originated: your PanicExample function.
		3. :70: The line number where the panic statement is located.
		4. +0x68: A hexadecimal offset indicating the exact byte position within the line.
		5. main.main(): This shows the function that called PanicExample. In this case, it's your main function, the entry point of your program.
		6. /Users/ayushkumar/programming/Go_2/15 Error Handling/main.go:18 +0x4c: Similar to the previous entry, this shows the location of the call to PanicExample within your main function.
		7. exit status 2: As mentioned before, this indicates that the program terminated abnormally due to the panic.
*/

func RecoverExample() {
	// Note: The panic statement immediately terminates the program. However, sometimes it might be important for a program to complete its execution and get some required results.
	// In such cases, we use the recover statement to handle panic in Go. The recover statement prevents the termination of the program and recovers the program from panic.
	// Recover is a built-in function that regains control of a panicking goroutine

	division(4, 2)
	division(8, 0)
	division(2, 8)
}

func division(num1, num2 int) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

// -> recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}
}

/*
	? Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. So if a defered function is called before panicking, will it still be called once panicking starts? give example with outputs to show.

	-> Deferred functions in Go will still be executed even if ⁠panic is called before the deferred function is scheduled to run. Here's why and an example to illustrate it: Deferred Functions and Panics.

	*	Deferred functions are stack-based: When you use ⁠defer, the function you specify is added to a stack associated with the current function. This stack is processed in reverse order when the function returns (either normally or due to a panic).
	*	Panic unwinds the stack: When ⁠panic is called, the program unwinds the call stack, running deferred functions in reverse order as it goes.

	-> Example:
	package main

	import "fmt"

	func myFunc() {
		defer fmt.Println("Deferred 1")
		defer fmt.Println("Deferred 2")

		fmt.Println("Inside myFunc")
		panic("Oops!") // Trigger a panic
		fmt.Println("This line won't execute") // This code is unreachable
	}

	func main() {
		defer fmt.Println("Deferred in main")
		myFunc()
		fmt.Println("This line won't execute either") 
	}
	
	-> Output:
		Inside myFunc
		Deferred 2
		Deferred 1
		Deferred in main
		panic: Oops!
		goroutine 1 [running]:
		main.myFunc(...)
				/home/user/go/src/main.go:11 +0x60
		main.main()
				/home/user/go/src/main.go:16 +0x38

	-> Explanation:
		1.	⁠myFunc is called: The code inside ⁠myFunc executes.
		2.	Deferred functions added: The ⁠defer statements add two functions to the stack.
		3.	⁠panic is called: Execution stops in ⁠myFunc, but the deferred functions are not skipped.
		4.	Stack unwinding: The program unwinds the call stack, running the deferred functions in reverse order: ⁠Deferred 2, ⁠Deferred 1.
		5.	⁠panic message: The panic message is printed (along with the stack trace).
		6.	Unreachable code: The lines after ⁠panic are never executed.
		7.	Deferred in main: The deferred function in ⁠main executes after the panic.
*/