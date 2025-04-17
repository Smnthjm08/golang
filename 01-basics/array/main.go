package main

import "fmt"

func main() {
	fmt.Println("we are learning array")

	var name [5]string
	name[0] = "a"
	name[1] = "b"
	name[2] = "c"
	name[3] = "d"
	name[4] = "e"

	fmt.Println("name is:", name)

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is:", nums)

	fmt.Println("len:", len(nums))

	fmt.Println("values at index 2", name[2])

	var price [5]string

	price[0] = "bleh"
	price[3] = "blah"
	fmt.Println("price is:", price)
	fmt.Printf("price array is: %q\n", price)
}
