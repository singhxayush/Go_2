package main

import (
	"fmt"

	utils "learning_go/myutil"
)

func main() {
	fmt.Println("Hello Go!")
	// myutil.printMessage("123") //-> Wont work because the function is a part of an external package and is not exported [not exported because it doesnt start with a capital letter] - unexported func are also called private func
	utils.PrintMessage1("This will work.")
}