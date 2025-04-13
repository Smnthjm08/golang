package main

import (
	"fmt"
	// "hello-world/myutils"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("=============")

	var name string = "String"
	var version float32 = 1.01
	var users = 100000
	var stable bool = false
	stable = true

	const PI = 3.14
	url := "http://url.example.com"

	// var Public = "Public"
	var Private = "Private"

	fmt.Println(name)
	fmt.Println(version)
	fmt.Println(users)
	fmt.Println(stable)
	fmt.Println(PI)

	fmt.Println(url)

	fmt.Println(Private)

	// myutils.PrintMessage("hello world 2")
}
