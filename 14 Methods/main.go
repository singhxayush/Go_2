package main

import "fmt"

func main() {
	fmt.Printf("-----------\n| Methods |\n-----------\n\n")
	ayush := user{"Ayush", "ayush@go.dev", true, 22}
	ayush.GetStatus()
	ayush.ChangeEmail()
	fmt.Printf(ayush.email)
}

// NOTE
type user struct {
	name   string
	email  string
	status bool
	age    int
}

func (u user) GetStatus() {
	fmt.Printf("Status of %v: %v\n", u.name, u.status)
}

func (u user) ChangeEmail() { //-> This won't change the email, until the user is passed by reference.
	u.email = "new_ayush@go.dev"
	fmt.Println(u.email)
}