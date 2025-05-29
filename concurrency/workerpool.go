package main

import (
	"fmt"
	"time"
)

func work(jobs <-chan int, results chan<- string) {
	for j := range jobs {
		time.Sleep(time.Second * 2)
		results <- fmt.Sprintf("processed job >>> %d", j)
	}
}

func main() {
	workerCount := 3
	jobs, results := make(chan int, 100), make(chan string, 100)

	for w := 0; w < workerCount; w++ {
		go work(jobs, results)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 0; a < 100; a++ {
		fmt.Println(<-results)
	}
}
