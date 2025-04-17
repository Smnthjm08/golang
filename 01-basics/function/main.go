package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	return a + b
}
func multiply(a, b int) (result int) {
	result = a * b
	return result
}

func main() {
	simpleFunction()

	ans1 := add(3, 4)
	ans2 := multiply(3, 4)
	fmt.Println("add", ans1)
	fmt.Println("multiply", ans2)

	fmt.Println(("We are learning functions"))
}
