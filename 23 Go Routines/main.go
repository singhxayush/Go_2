package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Note: Go Routines are lightweight, concurrent functions that run independently.

func main() {
	// greeter_caller()
	fmt.Println("Staring Normal getStatusCode")
	start := time.Now()
	getStatusCode_caller()
	elapsed := time.Since(start)
	fmt.Printf("Took %v\n\n", elapsed)

	time.Sleep(2 * time.Second)

	fmt.Println("Staring Go Routine getStatusCode")
	start = time.Now()
	getStatusCode_GoRoutine_caller()
	elapsed = time.Since(start)
	fmt.Printf("Took %v\n\n", elapsed)

}

func greeter_caller() {
	go greeter("Go Routine")
	greeter("Main")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(s)
	}
}

/// Get Status Code (without go routines)

func getStatusCode_caller() {
	endpointlist := []string{
		"https://youtube.com",
		"https://google.com",
		"https://fb.com",
		"https://chess.com",
	}

	for _, endpoint := range endpointlist {
		getStatusCode(endpoint)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		// fmt.Println(res, "Status Code of", endpoint)
		fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
	}
}

/// Get Status Code (go routines)

var wait_group sync.WaitGroup

func getStatusCode_GoRoutine_caller() {
	endpointlist := []string{
		"https://youtube.com",
		"https://google.com",
		"https://fb.com",
		"https://chess.com",
	}

	for _, endpoint := range endpointlist {
		go getStatusCode_GoRoutine(endpoint)
		wait_group.Add(1)
	}

	wait_group.Wait()
}

func getStatusCode_GoRoutine(endpoint string) {
	defer wait_group.Done() // Passes a signal to the wait group that task is done

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		// fmt.Println(res, "Status Code of", endpoint)
		fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
	}
}
