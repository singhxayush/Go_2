package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)
	// fmt.Println(*ptr) //-> Error [cannot deference a nil pointer]

	// myNum := 23
	// var ptr *int = &myNum
	// var ptr = &myNum
	// fmt.Println(ptr)
	// fmt.Println(*ptr)
	// fmt.Printf("Type of ptr : %T", ptr)
}