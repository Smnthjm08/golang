package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	nums = append(nums, 5, 6, 7, 23, 43, 34, 223)

	fmt.Println("Slice:", nums)

	fmt.Println("length of the slice:", len(nums))

	// name := []string{}
	// make Func - used to create a slice with a speciify length and capacity
	name := make([]int, 3, 5)
	// stock := make([]int, 0)
	// if capacity exceeds then the4 value is doubled
	name = append(name, 2)
	name = append(name, 6)
	name = append(name, 7)
	fmt.Println("name:", name)
	fmt.Println("len:", len(name))
	fmt.Println("cap:", cap(name))

}
