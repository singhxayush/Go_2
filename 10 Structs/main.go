package main

import "fmt"

func main() {
	fmt.Printf("-----------\n| Structs |\n-----------\n\n")

	ayush := user{"ayush", "ayush@go.dev", true, 22}
	fmt.Println(ayush)
	fmt.Printf("%+v\n", ayush)
	fmt.Printf("Age: %v\n", ayush.age)


}
// NOTE 
type user struct { //-> it's a good practice to declare attributes of a struct with Caps, it serves well in case of exporting it.
	name   string
	email  string
	status bool
	age    int
}

// func (u user) GetStatus() {
// 	fmt.Printf("Status of %v: %v\n", u.name, u.status)
// }