package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Printf("--------------\n| Maps in Go |\n--------------\n\n")
	mp := make(map[int]int)
	mp[1]++
	mp[3] += 2
	mp[2] += 10
	mp[20] += 3
	fmt.Println(mp)
	mp[3] -= 2
	fmt.Println(mp)
	delete(mp, 3)
	fmt.Println(mp)

	for key, val := range mp {
		fmt.Println(key, ":", val)
	}

	_, prs := mp[7] //-> The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn't need the value itself, so we ignored it with the blank identifier _.
	fmt.Println("prs:", prs)

	
    clear(mp) //-> To remove all key/value pairs from a map, use the clear builtin.
    fmt.Println("map:", mp)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
