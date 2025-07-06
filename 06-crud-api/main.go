package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        *int   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Println("Error getting data")
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error getting response", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error reading:", err)
	// 	return
	// }

	// fmt.Println("data", string(data))

	var todo Todo
	// var todo []Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("todo: ", todo)
	fmt.Println("todo.Completed: ", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "prince kumar",
		Completed: true,
	}
	// convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)
	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/"

	// send post request
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer res.Body.Close()
	fmt.Println("res.Body", res.Body)

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("status: ", res.StatusCode)
	fmt.Println("json data:", string(data))
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23,
		Title:     " wjdcbejw wjenfjewnef eh jekn chjdencdbnv d dh deh ",
		Completed: true,
	}

	// convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error updating the data: ", err)
	}

	req.Header.Set("Content-type", "application/json")

	// send the request
	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending the request: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("status: ", res.StatusCode)
	fmt.Println("json data:", string(data))

}

func performDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error updating the data: ", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	// send the request
	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending the request: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("status: ", res.StatusCode)
	fmt.Println("json data:", string(data))

}

func main() {
	fmt.Println("learning CRUD...")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
