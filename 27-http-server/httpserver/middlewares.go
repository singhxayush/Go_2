package main

import (
	"log"
	"net/http"
)

/*
  Middleware in Go can be visualized as a layer that wraps around HTTP handlers.
  Each middleware layer can execute logic before and after calling the next handler in the chain.

  Creating Middleware in Go
  In Go, middleware functions generally follow this pattern:

  1. They take in an http.Handler as an argument.
  2. They return a new http.Handler that wraps around the original handler,
  adding the desired functionality.

*/

// Explicitly defining a type for middleware functions to handle bigger chainings
type Middleware func(http.Handler) http.Handler

// Chaining the functions via loop that takes a slice of type Middleware
func chainMiddleWares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

// Middleware function that takes an http.Handler and returns another http.Handler
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		// next handler
		next.ServeHTTP(w, r)
	})
}

// Authentication middleware to check for a valid token
func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for a token in the Authorization header
		token := r.Header.Get("Authorization")
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Handler function for the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Fucking World!"))
}

// Handler function for the /goodbye route
func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye, World!"))
}

// Middleware func showing registering of Middleware functions to a specific path and chaning multiplexers, theres a better version down below
func Middleware1() {
	// Create a new HTTP multiplexer (router)
	mux := http.NewServeMux()

	// Register the router
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/goodbye", goodbyeHandler)

	// Wrap the mux with logging middleware
	// This line wraps the multiplexer (mux) with the logging middleware.
	// In simpler terms, it means that every time the server receives a request,
	// the logging middleware will be executed before the request reaches the actual routes (like /hello or /goodbye).

	// Chained wrapping of mux wrappers
	wrappedMux := loggingMiddleware(mux)
	authMux := authenticationMiddleware(wrappedMux)

	log.Println("Server is running on PORT::8888")
	log.Fatal(http.ListenAndServe(":8888", authMux))
}

func Middleware2() {
	mux := http.NewServeMux()

	// Register the routes
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/goodbye", goodbyeHandler)

	// Store middlewares in a slice
	middlewares := []Middleware{
		loggingMiddleware,
		authenticationMiddleware,
		// Add more middlewares here
	}

	// Chain them together
	chain := chainMiddleWares(mux, middlewares...)

	// Start the server
	log.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", chain))
}
