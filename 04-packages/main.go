package main

import (
	"fmt"
	"strings"
	"time"
)

// strings package
func main() {
	// data := "apple, orange, banana"
	// parts := strings.Split(data, ",")
	// fmt.Println(parts)

	// str := "one two three four two two five"
	// count := strings.Count(str, "two")
	// fmt.Println("count: ", count)

	// str = "  Hello,         Go!     "
	// trimmed := strings.TrimSpace(str)
	// fmt.Println("trimmed:", trimmed)

	str1 := "sumanth"
	str2 := "user"
	result := strings.Join([]string{str1, "is", str2}, " space ")
	fmt.Println("result", result)

	// time package
	// 2005-01-02 15:04:05 -year-month-day hour:minute:second
	layout_str := "02/01/2006"
	dateStr := "02/01/2024"
	currentTime := time.Now()
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("formatted time:", formatted_time)
	new_date := currentTime.Add(time.Hour * 24)
	fmt.Println("new date:", new_date.Format("2006-01-02 15:04:05"))
	// formatted := currentTime.Format("01-02-2006, 03:04 Monday")
	// formatted := currentTime.Format("01-02-2006, 03:04 Monday")

	// AM should be in capital letters and also they are also should be in the 12 hr format if we want it to
	// be AM then give AM and if we want PM then PM
	// formatted := currentTime.Format("01-02-2006, 15:04:05")
	// fmt.Println("formatted time:", formatted)

}
