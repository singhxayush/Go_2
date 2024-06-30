package main

import "fmt"

//syntax: func $func-name ($parameters) $return-type { code }
//-> example: func sum(num1 int, num2 int) int64 {

func main() {
	fmt.Printf("-------------\n| Functions |\n-------------\n\n")
	greeter()
	fmt.Println(sum(1, 2))
	fmt.Println(proAdder(1, 2, 11, -1))
}

func greeter() {
	fmt.Println("Namaste!")
}

func sum(num1 int, num2 int) int64 {
	return int64(num1)+int64(num2)
}

func proAdder(val...int) int64 { //-> Here the Values are now a slice
	var sum int64 = 0;
	for _, x := range val {
		sum += int64(x)
	}
	return sum
}
