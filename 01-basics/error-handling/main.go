package main

import "fmt"

// func divide(a, b float64) (float64, error) {
func divide(a, b float64) (float64, string) {
	if b == 0 {
		// return 0, fmt.Errorf("b is 0")
		return 0, "denominator can be zero"
	}
	return a / b, "nill"
}

func main() {
	fmt.Println("started Error handling")
	ans, err := divide(4, 0)

	if err != "denominator can be zero" {
		fmt.Println("Errof handking")
	}
	println("answer for division is", ans)
}
