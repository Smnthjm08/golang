package main

import "fmt"

// for loops can be controlled by using break and
// continue if we want like the case of while of others

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0

	// for loop
	for {
		fmt.Println("infinite loop")
		counter++
		if counter == 2 {
			break
		}
	}

	// range keyword
	numbers := []int{1, 2, 3, 4, 5}
	for idx, value := range numbers {
		fmt.Printf("idx: %d, value: %d\n", idx, value)
	}

	text := "Hello World!"
	for idx, char := range text {
		fmt.Printf("index of data is %d, and values is %c\n", idx, char)
	}
}
