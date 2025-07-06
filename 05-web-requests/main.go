package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning web services")
	res, err := http.Get("https://dummyjson.com/todos")
	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("status\n", res.Status)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("response\n", data)

	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}

	fmt.Println("response", string(data))
}
