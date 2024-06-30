package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("-----------\n| If-Else |\n-----------\n\n")

	// rand.Seed(time.Now().UnixNano()) // Note: this is depricated

	// Create a new source for random numbers with a seed
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new random generator from the source
	r := rand.New(source)
	n := r.Intn(10)+1

	if n%2 == 0 {
		fmt.Printf("n=%v is even\n", n)
	} else {
		fmt.Printf("n=%v is odd\n", n)
	}
	//-> You can have an if statement without an else.

	n = r.Intn(10) + 1
	if n%2 == 0 || (n+1)%2 == 0 {
		fmt.Printf("either %v or %v are even \n", n, n+1)
	}
	//-> A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

