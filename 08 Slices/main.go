package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Watermelon", "Cherry", "Peach"}
	fmt.Printf("Type of fruitList: %T\n", fruitList)
	fmt.Println("fruitList:", fruitList)

	fruitList = append(fruitList, "Mango", "Pineapple")
	fmt.Println("fruitList:", fruitList)

	fruitList = append(fruitList[1:], "")
	fmt.Println("fruitList:", fruitList)

	highScores := make([]int, 5)
	highScores[0] = 1
	highScores[1] = 2
	highScores[2] = 9
	highScores[3] = 0
	highScores[3] = -1

	fmt.Println(highScores)
	// sort.Slice(highScores, func(i, j int) bool {
	// 	return highScores[i] > highScores[j] // Sort in descending order
	// })

	sort.Ints(highScores)
	fmt.Println(highScores)

	// removing elements from slice
	var idx int = 2
	highScores = append(highScores[:idx], highScores[idx+1:]...)
	fmt.Println(highScores)

	highScores2 := highScores[1:3]
	fmt.Println("highScores2:", highScores2)

	t := []string{"g", "h", "i"}
    t2 := []string{"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d:", twoD)
}
