package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
	-> Refer to this artcile also:
	https://www.kelche.co/blog/go/golang-file-handling/

	Note: When you open a file using os.OpenFile (or os.Open, os.Create, etc.), it returns a file pointer (of type *os.File). Initially, this pointer points to the starting byte of the file.
	Here's a quick summary of how file pointers work with respect to reading and writing:

	-> File Pointer Position: When a file is opened, the file pointer is initially set to the beginning of the file (byte 0).

	-> Reading/Writing Operations: Each read or write operation advances the file pointer by the number of bytes read or written. For example, if you write 10 bytes to the file, the file pointer moves 10 bytes forward.

	-> Seeking: You can change the position of the file pointer using the Seek method. This is useful when you need to read from or write to a specific position in the file.


*/

func main() {
	fmt.Printf("---------\n| Files |\n---------\n")
	file1Operations("file1.txt")
	file2Operations("file2.txt")
}

func file1Operations(filename string) {
	fmt.Printf("\n------\nFile Operations on: %v\n", filename)
	dataForFile1 := "This needs to go in a file1"
	file := OpenFile(filename)
	defer file.Close()
	WritingToFile(file, dataForFile1)
	// ReadingFromFile_usingIOutils(filename)
	ReadingFromFile_usingIO(file) //-> Needs reference of the opened file.
}

func file2Operations(filename string) {
	fmt.Printf("\n------\nFile Operations on: %v\n", filename)

	file := OpenFile(filename)
	defer file.Close()

	WritingToFile(file, "line1\n")
	AppendToFile(file, "line2 - Appended")
	AppendToFile(file, "line2 - Appended")
	ReadingFromFile_usingIO(file)
}

func OpenFile(filename string) *os.File {
	fmt.Printf("\n#Opening: %v\n", filename)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	/*
	*	0666 value is a file permission setting.
	* 	Flags for os.OpenFile
	*		os.O_RDWR: Open the file for both reading and writing.
	*		os.O_CREATE: Create the file if it does not exist.
	*		These flags can be combined using the bitwise OR operator (|)
	*		For example, os.O_RDWR | os.O_CREATE means "open the file for reading and writing, and create it if it doesn't exist."
	 */

	if err != nil {
		panic(err)
	}
	return file
}

func WritingToFile(file *os.File, data string) int {
	fmt.Printf("\n#Writing to: %v\n", file.Name())

	length, err := io.WriteString(file, data)
	if err != nil {
		panic(err)
	}
	return length
}

func ReadingFromFile_usingIOutils(filename string) {
	fmt.Printf("\n#Reading From: %v\n", filename)
	dataByte, err := ioutil.ReadFile(filename) // -> this is depricated as of Go 1.16

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string
	content := string(dataByte)

	fmt.Println("File contents:")
	fmt.Println(content)
}

func ReadingFromFile_usingIO(file *os.File) {
	// Reset the file pointer to the beginning
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	// Read the file contents
	fmt.Printf("\n#Reading From: %v\n", file.Name())
	dataByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the file contents
	fmt.Println("File contents:")
	fmt.Println(string(dataByte))
}

func AppendToFile(file *os.File, data string) {
	fmt.Printf("\n#Appending to: %v\n", file.Name())

	_, err := fmt.Fprintln(file, data)
	if err != nil {
		fmt.Println("Error Appending to file:", err)
	}
}

/*

	Note: File Permissions
	The 0666 is a file mode that sets the file's permissions.
	In Unix-like systems, file permissions determine who can read, write, or execute a file.
	The value 0666 is an octal number (base 8), which is a common way to represent Unix file permissions.
	Each digit represents different permissions:

	* The first digit is for the owner of the file.
	* The second digit is for the group associated with the file.
	* The third digit is for others (everyone else).
	* Each digit is a sum of the following values:

	-> 4 (read)
	-> 2 (write)
	-> 1 (execute)
	-> So, 6 means read (4) + write (2) = 6.

	* 0644: Owner can read and write, group and others can read.
		- Owner: 6 (read + write)
		- Group: 4 (read)
		- Others: 4 (read)

	* 0755: Owner can read, write, and execute, group and others can read and execute.
		- Owner: 7 (read + write + execute)
		- Group: 5 (read + execute)
		- Others: 5 (read + execute)

	* 0700: Owner can read, write, and execute; group and others have no permissions.
		- Owner: 7 (read + write + execute)
		- Group: 0 (no permissions)
		- Others: 0 (no permissions)

	Note: Common Flags

	-> os.O_RDONLY: Open the file for read-only access.
	-> os.O_WRONLY: Open the file for write-only access.
	-> os.O_RDWR: Open the file for both reading and writing.
	-> os.O_CREATE: Create the file if it does not exist.
	-> os.O_APPEND: Append data to the end of the file.
	-> os.O_EXCL: Ensure that the file is created only if it does not already exist.
	-> os.O_TRUNC: Truncate the file to zero length if it exists.

	* file, err := os.OpenFile("example.txt", os.O_RDONLY, 0444)
		- Opens the file for read-only access
		- Owner, group, and others can read the file.

	* file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE, 0644)
		- Opens the file for write-only access. If the file does not exist, it is created
		- Owner can read and write; group and others can read.

	* file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		- Opens the file for reading and writing. If the file exists, it is truncated to zero length (content is cleared). If it doesn't exist, it is created.
		- Owner, group, and others can read and write.

	* file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		- Opens the file for write-only access and appends data to the end of the file. If the file does not exist, it is created.
		- Owner, group, and others can read and write.

	* file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		- Creates a new file for write-only access. If the file already exists, the operation fails.
		- Owner, group, and others can read and write.
*/
