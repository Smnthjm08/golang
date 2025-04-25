package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println("hello world")
	// // create a new file
	// file, err := os.Create("example.txt")
	// // error handling of the file creation
	// if err != nil {
	// 	fmt.Println("error while creating files ", err)
	// 	return
	// }
	// defer file.Close()

	// // writing a file content
	// content := "This is the content in this file 2"
	// byte, errors := io.WriteString(file, content)
	// fmt.Println("byte", byte)

	// if errors != nil {
	// 	fmt.Println("errors while creating files ", errors)
	// 	return
	// }

	// // success message
	// fmt.Println("successfully created file")

	// opening a file
	file, err := os.Open("example.txt")
	fmt.Println("file", file)
	if err != nil {
		fmt.Println("Error while opening file", err)
	}

	defer file.Close()

	// create buffer
	buffer := make([]byte, 1024)

	// read file content into the buffer
	for {
		n, err := file.Read(buffer)

		if err == io.EOF {

			break
		}

		if err != nil {
			fmt.Println("error while reading the file", err)
			return
		}
		// process the read content
		fmt.Println(string(buffer[:n]))
	}

}
