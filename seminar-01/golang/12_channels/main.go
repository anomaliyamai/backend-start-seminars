package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Отправлено: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("Получено: %d\n", num)
	}

	/*
		for {
		    num, ok := <-ch
		    if !ok {
		        break
		    }
		    fmt.Printf("Получено: %d\n", num)
		}
	*/
}

func main() {
	ch := make(chan int)

	// var wg sync.WaitGroup

	// wg.Add(1)

	// go counter(ch)
	// go receiver(ch, &wg)

	// wg.Wait()

	ch <- 42


 }
