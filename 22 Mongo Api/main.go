package main

import (
	"fmt"
	"log"
	"net/http"

	route "myMongoApi/router"
)

func main() {
	fmt.Printf("-------------\n| Mongo Api |\n-------------\n")
	fmt.Println("Server Running...")

	r := route.Router()
	fmt.Println("Liseting on ::8000 | http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
