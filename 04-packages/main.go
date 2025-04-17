package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, orange, banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count: ", count)

	str = "  Hello,         Go!     "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed:", trimmed)

	str1 := "sumanth"
	str2 := "user"
	result := strings.Join([]string{str1, "is", str2}, " space ")
	fmt.Println("result", result)
}
