package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func makeHttpGetReq() {
	// Response represents the response from an HTTP request.
	// func Get
	// func Get(url string) (resp *Response, err error)
	resp, err := http.Get("http://localhost:8888/greet?name=ayush")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading the response")
		return
	}

	fmt.Println("Response from the server: ", string(body))
}

func makeHttpPostReq() {
	data := []byte("This is a post request")

	resp, err := http.Post("http://localhost:8888/post", "text/plain", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading the response", err)
		return
	}

	fmt.Println("Response from the server: ", string(body))
}

func makeHttpPostReqWithJson() {
	jsonData := map[string]string{"name": "John", "message": "I am a Gopher"}

	jsonValue, err := json.Marshal(jsonData)

	if err != nil {
		log.Fatalf("Error marshling json data %v", err)
	}

	res, err := http.Post(
		"http://localhost:8888/post2",
		"application/json",
		bytes.NewBuffer(jsonValue),
	)

	if err != nil {
		log.Fatalf("Error making post request", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error reading response %v", res.Body)
	}

	fmt.Printf("Response: %s\n", body)
}

/*
  The Content-Type HTTP header specifies the media type of the resource being sent to the
  server in an HTTP request (like POST) or returned from the server in an HTTP response.
  It helps the server or client understand the format of the content, so they can properly
  parse and process it.
*/

func main() {
	// makeHttpGetReq()
	// makeHttpPostReq()
	makeHttpPostReqWithJson()

}
