package main

import "fmt"

func main() {
	day := 0

	switch day {
	case 0:
		fmt.Println("Today is sunday")
	case 1:
		fmt.Println("Today is monday")
	case 2:
		fmt.Println("Today is tuesday")
	case 3:
		fmt.Println("Today is wednesday")
	case 4:
		fmt.Println("Today is thursday")
	case 5:
		fmt.Println("Today is friday")
	case 6:
		fmt.Println("Today is saturday")
	default:
		fmt.Println("none of the days of the week")
	}

}
