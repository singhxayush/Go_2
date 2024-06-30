package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userinput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	//-> Comma OK || err OK syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Rating = ", input);

	// The comma-ok idiom is a way to handle the return values of functions that can either return a successful value and an error, or just an error.  Heres how it looks:
	// • Error Handling: The comma-ok idiom is central to Gos approach to error handling. It forces you to explicitly check for errors and deal with them gracefully.
	// • Zero Values: If ⁠someFunction() doesnt return a value (due to an error), the ⁠value variable will be set to its zero value. For example, an integer variable would be set to ⁠0, a string would be set to ⁠"", and a pointer would be set to ⁠nil.
	// • Conciseness:  The comma-ok idiom provides a clean and concise way to handle errors without cluttering your code with unnecessary error-checking variables.
}
