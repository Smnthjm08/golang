package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d ended\n", i)
}

func main() {
	fmt.Println("=====")
	var wg sync.WaitGroup

	// start 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1) // increment the wg counter
		go worker(i, &wg)
	}

	// wait for all the workers to finish
	wg.Wait()
	fmt.Println("worker task completed!")

}
