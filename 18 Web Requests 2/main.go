// Note: This section is all about "CRUD" operations of a go Web Server
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Printf("------------------\n| Web Requests 2 |\n------------------\n\n")

	PerformGetReq()
	fmt.Printf("\n\n")
	PerformPostJsonReq()
	fmt.Printf("\n\n")
	PerformPostFormReq()
}

func PerformGetReq() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)

	if err != nil {
		fmt.Println("Error getting the Req")
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	content, _ := io.ReadAll(response.Body)
	var responseString strings.Builder
	responseString.Write(content)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonReq() {
	const myUrl = "http://localhost:8000/post"

	// Fake Json payload
	requestBody := strings.NewReader(`
		{
			"name":"Hayes",
			"age":24,
			"profession":"singer"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error getting the Req")
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormReq() {
	const myUrl = "http://localhost:8000/postform"

	// Form data
	data := url.Values{}
	data.Add("firstname", "Ayush")
	data.Add("lastname", "Kumar")
	data.Add("email", "ayush@go.dev")
	resonse, err := http.PostForm(myUrl, data)
	if err != nil {
		fmt.Println("Error getting the Req")
		panic(err)
	}
	defer resonse.Body.Close()

	content, _ := io.ReadAll(resonse.Body)
	fmt.Println(string(content))
}
