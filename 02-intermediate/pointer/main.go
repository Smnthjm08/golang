// pointer in go
// way to work w memory directly
// - A pointer is a variable that stores the memory address of another variable.

package main

import "fmt"

//   - The `&` - get the memory address of a variable.
//   - The `*` - dereference a pointer, i.e., to access the value stored at the memory
//     address the pointer is pointing to.
func main() {
	// var num int
	// num = 2
	// var ptr *int
	// ptr = &num
	num := 2
	ptr := &num

	// 1. `num` int var initialized with 2.
	// 2. `ptr` is a pointer to an integer, which is assigned the memory address of `num` using the `&` operator.

	fmt.Println("num has value: ", num)
	//   - memory address in `ptr`.
	fmt.Println("ptr has: ", ptr)
	//   - The value stored at the memory address pointed to by `ptr` using the `*` operator.
	fmt.Println("Data contains through Pointer: ", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("pointer is not assigned")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("value contains:", value)
}

func modifyValueByReference(num *int) {
	*num = *num + 20
}
