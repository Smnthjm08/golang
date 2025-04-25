package main

import (
	"fmt"
)

func add(a, b int) int {
	fmt.Println("add")
	return a + b
}

// defer is used to delay the execution of a function until the surrounding function returns
// The deferred function's arguments are evaluated when the defer statement is executed
// this follows LIFO order
func main() {
	fmt.Println("1")
	data := add(5, 6)
	defer fmt.Println("2")
	defer fmt.Println("data", data)
	fmt.Println("3")
	// fmt.Println("1")
	// fmt.Println("1")
}
