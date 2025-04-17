package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("num is", num)
	fmt.Printf("type of %T", num)

	// num = num + 2.3

	var data float64 = float64(num)

	data = data + 3.14
	fmt.Println("data is ", data)
	fmt.Printf("type of %T", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("type of str is %T\n", str)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("num_int is", number_int)
	fmt.Printf("type of number_int %T\n", number_int)

}
