package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Route 1")
}

func helloRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}

func aboutRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func handleHelloQuery(w http.ResponseWriter, r *http.Request) {
	// name := r.URL.Query().Get("name")
	// name := r.URL.Query()["name"][0]
	slice_of_query_key_val := r.URL.Query()["name"]
	name := slice_of_query_key_val[0]
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func handlePostReq(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Recieved Post Request with data: %s", string(body))
}

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func handleSendJsonResponse(w http.ResponseWriter, r *http.Request) {
	// checking if the request from the client to the server(us) is valid or not
	// if not send a resonse with http.Error method and status code StatusMethodNotAllowed
	// https://pkg.go.dev/net/http@go1.23.1#Error
	if r.Method != http.MethodPost {
		http.Error(w, "Only post method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the body from the request
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	// Once the request is read and it's content is stored in the body, defer it
	defer r.Body.Close()

	// time to parse the json request from the body
	var msg Message

	// Unmarshel is the way to do that
	err = json.Unmarshal(body, &msg)

	if err != nil {
		http.Error(w, "Invalid json format, parsing failed", http.StatusBadRequest)
		return
	}

	fmt.Printf("Recieved -> Name: %s, Message: %s\n", msg.Name, msg.Message)

	// now writing the respose back to the client
	// Send a response
	response := fmt.Sprintf("Hello %s, your message '%s' was received!", msg.Name, msg.Message)
	w.Write([]byte(response))
}

func Run() {
	// route handling
	http.HandleFunc("/", homeRouteHandler)
	http.HandleFunc("/hello", helloRouteHandler)
	http.HandleFunc("/about", aboutRouteHandler)
	http.HandleFunc("/greet", handleHelloQuery)
	http.HandleFunc("/post", handlePostReq)
	http.HandleFunc("/post2", handleSendJsonResponse)

	fmt.Println("Server listening on port :: 8888")
	// The nil parameter indicates that the default multiplexer (router) is used

	log.Fatal(http.ListenAndServe(":8888", nil))
}
