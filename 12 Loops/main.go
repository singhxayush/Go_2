package main

import (
	"fmt"
)

func main() {
	fmt.Printf("---------\n| Loops |\n---------\n\n")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	for i := 0; i < len(days); i++ {
		fmt.Printf("Day %v : %v\n", i+1, days[i])
	}

	fmt.Println()

	for i := range days {
		fmt.Printf("Day %v : %v\n", i+1, days[i])
	}

	fmt.Println()

	for idx, days := range days {
		fmt.Printf("Day %v : %v\n", idx+1, days)
	}
}
