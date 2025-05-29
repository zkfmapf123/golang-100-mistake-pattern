package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch chan<- int, value int) {
	ch <- value
}

func receiver(ch <-chan int) {
	for c := range ch {
		fmt.Println("cc >> ", c)
	}
}

/*
close(ch)
데이터 전송을 모두 끝내면 close 해줘야 함
*/
func channel() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		sender(ch, 10)
		sender(ch, 20)
		sender(ch, 30)
		sender(ch, 40)
		sender(ch, 50)
	}()

	receiver(ch)
}

func syncWaitGroup() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)

		sender(ch, 10)
		sender(ch, 20)
		sender(ch, 30)
		sender(ch, 40)
		sender(ch, 50)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		receiver(ch)
	}()

	wg.Wait()
}

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		time.Sleep(time.Second)
		results <- fmt.Sprintf("worker %d processed job %d", id, j)
	}
}

func workerpool() {
	workerCount := 3
	jobs, results := make(chan int, 100), make(chan string, 100)

	// worker -> 3개씩 돌림
	for w := 0; w < workerCount; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 9; a++ {
		fmt.Println(<-results)
	}
}
