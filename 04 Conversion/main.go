package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Prompt user to input a rating
	fmt.Println("Rate b/n 0 - 10")
	// Create a new buffered reader to read from stdin
	reader := bufio.NewReader(os.Stdin)
	// Read a line from stdin until a newline character is encountered
	ip, _ := reader.ReadString('\n')
	// Print a thank you message along with the input rating
	fmt.Println("Thanks for Rating: ", ip)
	// Parse the input string as a float64, removing any leading or trailing whitespace
	numRating, err := strconv.ParseFloat(strings.TrimSpace(ip), 64)
	// Handle any error that occurred during parsing
	if err != nil {
		fmt.Println(err)
	} else {
		// Print the new rating after adding 1 to the parsed rating
		fmt.Println("New rating = ", numRating+1)
	}
}