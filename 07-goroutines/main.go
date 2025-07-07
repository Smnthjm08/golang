package main

import (
	"fmt"
	"time"
)

func sayOne() {
	fmt.Println("=========")
	// go sayTwo()
}

func sayTwo() {
	fmt.Println(">>>>>>>>>>")
}

func main() {
	fmt.Println("1")
	time.Sleep(100 * time.Millisecond)
	go sayOne()
	go sayTwo()
	fmt.Println("2")
}
