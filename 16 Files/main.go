package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("---------\n| Files |\n---------\n\n")
	file1operations("file1.txt")
}
func file1operations(filename string) {
	fmt.Printf("File Operations on: %v\n", filename)
	dataForFile1 := "This needs to go in a file1"
	file := OpenFile(filename)
	defer file.Close()
	WritingToFile(file, dataForFile1)
	ReadingFromFile_usingIOutils(filename)
	ReadingFromFile_usingIO(file) //-> Needs refrence of the opened file.
}

func OpenFile(filename string) *os.File {
	fmt.Printf("\nOpening the File\n")
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func WritingToFile(file *os.File, data string) int {
	fmt.Printf("\nWriting to File\n")

	length, err := io.WriteString(file, data)
	if err != nil {
		panic(err)
	}
	return length
}

func ReadingFromFile_usingIOutils(filename string) {
	fmt.Printf("\nReading the file\n")
	fileContent, err := ioutil.ReadFile(filename) // -> this is depricated as of Go 1.16

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string
	content := string(fileContent)

	fmt.Println("File contents:")
	fmt.Println(content)
}

func ReadingFromFile_usingIO(file *os.File) {
	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the file contents
	fmt.Println(string(data))
}
