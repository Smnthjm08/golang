package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
	Fax   string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details  Person
	Person_Contacts Contact
	Person_Address  Address
}

func main() {
	// Create a new Person
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Print the full details of the person
	fmt.Println("Person details:", person1)

	// new memory alllocated
	var person2 = new(Person)
	person2.Age = 22
	person2.FirstName = "Sumanth"
	person2.LastName = "JM"

	fmt.Println("Person 2", person2)

	fmt.Println("Age of john is", person1.Age)

	employee1 := new(Employee)
	employee1.Person_Details = Person{
		FirstName: "Sumanth",
		LastName:  "JM",
		Age:       26,
	}

	employee1.Person_Contacts.Email = "sumanth@gmail.com"
	employee1.Person_Contacts.Phone = "89238293332"
	employee1.Person_Contacts.Fax = "hdh"
	employee1.Person_Address = Address{
		House: 24,
		Area:  "Ranchi",
		State: "Jharkhand",
	}

	fmt.Println("Employee 1:", employee1)
}
