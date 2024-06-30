package main

import "fmt"

func main() {
	var array1 [4]string
	array1[0] = "a"
	array1[2] = "aa"
	array1[3] = "za"
	fmt.Println(len(array1))
	fmt.Printf("array1: %v\n", array1)

	var array2 = [3] string {"nn", "aa", "qq"}
	fmt.Println(len(array2))
	fmt.Printf("array2: %v\n", array2)

	array3 := [...] int {1, 2, 3, 1, 10} //-> Compiler count the number of elements for you with ...
	fmt.Println("array3:", array3)

	array4 := [...]int{100, 3: 400, 500} //-> If you specify the index with :, the elements in between will be zeroed.
	fmt.Println("array4:", array4)

	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}
