package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(("Whats your name?"))

	var name string

	// fmt.Scan(&name)

	reader := bufio.NewReader(os.Stdin)

	name, _ = reader.ReadString('\n')
	fmt.Println("Hello, Mr.", name)

}
