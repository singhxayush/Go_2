package main

import (
	"fmt"
	"log"
	"net/url"
)

func UrlParserExample() {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
	/*
		Parse parses a raw url such as like this into a URL structure

		type URL struct {
			Scheme      string
			Opaque      string    // encoded opaque data
			User        *Userinfo // username and password information
			Host        string    // host or host:port (see Hostname and Port methods)
			Path        string    // path (relative paths may omit leading slash)
			RawPath     string    // encoded path hint (see EscapedPath method)
			OmitHost    bool      // do not emit empty host (authority)
			ForceQuery  bool      // append a query ('?') even if RawQuery is empty
			RawQuery    string    // encoded query values, without '?'
			Fragment    string    // fragment for references, without '#'
			RawFragment string    // encoded fragment hint (see EscapedFragment method)
		}
	*/

	if err != nil {
		fmt.Println("Got error")
		log.Fatal(err)
	}

	// fmt.Println(u["a"])
	// fmt.Println(u.Scheme)

	q := u.Query()
	/*

		  https: //pkg.go.dev/net/url@go1.23.1#URL.Query
			u.Query() in Go is a method on the URL struct (from the net/url package)
			that parses the query string part of the URL into a url.Values type.
			The query string is the portion of the URL that comes after the ?,
			containing key-value pairs that represent additional data passed to the server.

		  The Query() method helps you extract and manipulate these key-value pairs in a structured format.

	*/

	fmt.Println(q["a"])
	fmt.Println(q["b"])
	fmt.Println(q["c"])

}
