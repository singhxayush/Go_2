package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const urlGet string = "https://api.genderize.io/?name=peter"

func main() {
	fmt.Printf("----------------\n| Web Requests |\n----------------\n\n")

	makeGetReq(urlGet)
	handleUrl(urlGet)
}

func makeGetReq(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // It is the caller's responsibilty to close the connection

	fmt.Printf("Type of res: %T\n", res)

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Error Reading from the request")
	}
	content := string(dataBytes)

	fmt.Println(content)
}

func handleUrl(urlH string) {
	resUrl, _ := url.Parse(urlH)
	fmt.Println(resUrl)

	qparams := resUrl.Query() // Note: Returns a MAP of query params key and value
	// fmt.Printf("Type of Querey params: %T\n", qparams)
	// fmt.Println("name:", qparams["name"])
	for k, v := range qparams {
		fmt.Println(k, ":", v)
	}

	// Constructing a URL from parts
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "api.genderize.io",
		Path:    "/dev",
		RawPath: "user=peter",
	}
	constructedUrl := partsOfUrl.String()
	fmt.Println("Final URL:", constructedUrl)

	// -> Fields:
	// fmt.Println(resUrl.Fragment)
	// *Returns the URL fragment (part after "#").
	// ?Return - type: string (the fragment, possibly empty)

	// fmt.Println(resUrl.Host)
	// *Returns the hostname and port (if any).
	// ?Return - type: string (e.g., "example.com:8080")

	// fmt.Println(resUrl.OmitHost)
	// *Returns a URL with the hostname and port removed.
	// ?Return - type: string (URL without host)

	// fmt.Println(resUrl.Opaque)
	// *Returns the opaque part of the URL (for non-standard schemes).
	// ?Return - type: string (opaque part)

	// fmt.Println(resUrl.Path)
	// *Returns the URL path.
	// ?Return - type: string (e.g., "/path/to/resource")

	// fmt.Println(resUrl.RawFragment)
	// *Returns the raw fragment (without escaping).
	// ?Return - type: string (the fragment as it was in the original URL)

	// fmt.Println(resUrl.RawPath)
	// *eturns the raw path (without escaping).
	// ?Return - type: string (the path as it was in the original URL)

	// fmt.Println(resUrl.RawQuery)
	// *eturns the raw query (without escaping).
	// ?Return - type: string (the query as it was in the original URL)

	// fmt.Println(resUrl.Scheme)
	// *Returns the URL scheme (protocol).
	// ?Return - type: string (e.g., "http", "https", "ftp")

	// fmt.Println(resUrl.User)
	// *Returns the user information (username:password).
	// ?Return - type: *url.Userinfo (if present)

	// -> Functions:
	// fmt.Println(resUrl.EscapedFragment())
	// *Returns the URL fragment, properly escaped.
	// ?Return - type: string (the escaped fragment)

	// fmt.Println(resUrl.EscapedPath())
	// *Returns the URL path, properly escaped.
	// ?Return - type: string (the escaped path)

	// fmt.Println(resUrl.Hostname())
	// *Returns the hostname without the port.
	// ?Return - type: string (e.g., "example.com")

	// fmt.Println(resUrl.IsAbs())
	// *Checks if the URL is absolute (has a scheme).
	// ?Return - type: bool (true if absolute, false otherwise)

	// fmt.Println(resUrl.MarshalBinary())
	// *Encodes the URL into a byte slice.
	// ?Return - type: []byte (encoded URL)

	// fmt.Println(resUrl.Parse())
	// *Parses a URL relative to the current URL.
	// ?Return - type: *url.URL (the parsed URL, or nil if an error)

	// fmt.Println(resUrl.Port())
	// *Returns the port number, if any.
	// ?Return - type: string (e.g., "8080")

	// fmt.Println(resUrl.Query())
	// *Returns the query parameters as a `url.Values` map.
	// ?Return - type: url.Values (map-like structure)

	// fmt.Println(resUrl.Redacted())
	// *Returns a redacted URL, hiding sensitive information (usernames/passwords).
	// ?Return - type: string (redacted URL)

	// fmt.Println(resUrl.RequestURI())
	// *eturns the request URI, suitable for HTTP requests.
	// ?Return - type: string (request URI)

	// fmt.Println(resUrl.ResolveReference())
	// *Resolves a relative reference against the current URL.
	// ?Return - type: *url.URL (the resolved URL, or nil if an error)

	// fmt.Println(resUrl.String())
	// *Returns the string representation of the URL.
	// ?Return - type: string (the URL)

	// fmt.Println(resUrl.User.Password())
	// *Returns the password from the user information.
	// ?Return - type: string (password, if present)

	// fmt.Println(resUrl.User.String())
	// *eturns the string representation of user information.
	// ?Return - type: string (e.g., "username:password")

	// fmt.Println(resUrl.User.Username())
	// *eturns the username from the user information.
	// ?Return - type: string (username, if present)
}

