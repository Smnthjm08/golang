package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning URL")
	myNewUrl := "https://dummyjson.com/todos?id=23"
	parsedUrl, err := url.Parse(myNewUrl)
	if err != nil {
		fmt.Println("Cant parse url", err)
	}
	fmt.Println("type of url", parsedUrl.Host)
	fmt.Println("type of url", parsedUrl.Scheme)
	fmt.Println("type of url", parsedUrl.Path)
	fmt.Println("type of url", parsedUrl.RawQuery)

	// modifying url
	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=smnthjm08"

	newUrl := parsedUrl.String()
	fmt.Println(">>>>>>", newUrl)
}
