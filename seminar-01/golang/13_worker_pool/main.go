package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Воркер %d начал %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Воркер %d закончил %d\n", id, j)

		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		fmt.Println("Результат:", <-results)
	}
}
