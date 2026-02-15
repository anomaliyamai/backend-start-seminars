package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello!")

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Wow")
	}()

	wg.Wait()
}
