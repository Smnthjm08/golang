package main

import (
	"fmt"
)

// if else conditions
func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is smaller than 5")
	}

	z := -1.0
	fmt.Printf("z is of type %T\n", z)
	if z > -20 && (x <= 7 || z == -1) {
		fmt.Println("z > 10 && x < 10")
	} else if z > 5 {
		fmt.Println("z is greater than 10")
	} else {
		fmt.Println("z is smaller than 5")
	}
}

// func main() {
// 	y := 4
// 	if y > 4 {
// 		fmt.Println("y is greater than 4")
// 	} else if y == 4 {
// 		fmt.Println("y is equal to 4")
// 	} else {
// 		fmt.Println("y is lesser than 4")
// 	}
// }
