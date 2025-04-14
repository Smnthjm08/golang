package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["User1"] = 34
	studentGrades["User2"] = 78
	studentGrades["User3"] = 56
	studentGrades["User4"] = 89
	studentGrades["User5"] = 92

	fmt.Println("User2>", studentGrades["User2"])
	studentGrades["User2"] = 20
	fmt.Println("User2>", studentGrades["User2"])

	delete(studentGrades, "User2")
	fmt.Println("User2>", studentGrades["user3"])

	grades, exists := studentGrades["User6"]
	fmt.Println("User3>", grades)
	fmt.Println("User3 boolean>", exists)

	fmt.Println("User3>", studentGrades["User3"])

}
