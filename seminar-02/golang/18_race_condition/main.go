package main

import (
	"fmt"
	"sync"
)

func main() {
	text := ""
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)

	go func() {
		defer wg.Done()
		mu.Lock()
		text = "hello world"
		mu.Unlock()
	}()

	wg.Wait()
	wg.Add(1)

	go func() {
		defer wg.Done()
		mu.Lock()
		fmt.Println(text)
		mu.Unlock()
	}()

	wg.Wait()
}
