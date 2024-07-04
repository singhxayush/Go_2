// Note: Encoding of JSON
// DOCS: https://pkg.go.dev/encoding/json

package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //-> These are basically alias, a part of struct (read more ways in DOCS)
	Price    int
	Platform string
	Password string   `json:"-"`              //-> This alias means that this field won't be reflected in the JSON encoding
	Tags     []string `json:"tags,omitempty"` //-> This simply means that this section will be omitted from the encoding if empty
}

func main() {
	fmt.Printf("--------\n| JSON |\n--------\n")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{
			"ReactJs",
			200,
			"YT-Prem",
			"P@ssword",
			[]string{
				"Web-dev",
				"Frontend",
			},
		},
		{
			"NodeJs",
			250,
			"YT-Prem",
			"P@ssword",
			nil,
		},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs",
			"Price": 200,
			"Platform": "YT-Prem",
			"tags": [
				"Web-dev",
				"Frontend"
			]
        }
	`)

	var myCourse course

	isValid := json.Valid(jsonDataFromWeb)
	if isValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON data is invalid")
	}
	fmt.Printf("\n")

	// Add data to key-value pair
	var myOnlineData map[string]interface{}
	// interface{}: The ⁠interface{} type is Go's "empty interface."  This means that the values associated with your string keys can be any type.  It's a flexible but potentially unsafe way to store diverse data.
	// Performance: The dynamic nature of ⁠interface{} can sometimes result in performance overhead.
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	fmt.Printf("\n")

	for key, val := range myOnlineData {
		fmt.Printf("%v -> %v(%T)\n", key, val, val)
	}

}
