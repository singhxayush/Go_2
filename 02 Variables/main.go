package main

import (
	"fmt"
)

func main() {

	fmt.Print("Learning Variables\n\n")
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("Type of the variable is: %T\n", username)

	var isTrue bool = true
	fmt.Println(isTrue)

	var smallVal uint8 = 255 //-> will give error for 256
	fmt.Println(smallVal)

	var num int; // 0 by default
	fmt.Println(num)
	fmt.Printf("Type of the variable is: %T\n", num)


	var website = "google.com" //-> no variables assigend, lexer does that job
	fmt.Println(website)

	// -- No Var Style --
	// short variable declaration or assignment operator -> :=
	someNum := 2000 // doesnt work globally	
	fmt.Println(someNum)
}

// press -> hl for this
// package main 

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

// func main() {
// 	http.HandleFunc("/", greet)
// 	http.ListenAndServe(":8080", nil)
// }
