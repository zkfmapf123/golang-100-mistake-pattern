package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
1. 경쟁상태 (atomic 활용)
*/
func _27_atomic() {
	var i int64
	go func() {
		atomic.AddInt64(&i, 1)
	}()

	go func() {
		atomic.AddInt64(&i, 1)
	}()

	fmt.Println(i)
}

/*
2. 경쟁상태 Critical Section 활용
*/
func _27_criticalSection() {
	var i int64

	var mu sync.Mutex

	go func() {
		mu.Lock()
		defer mu.Unlock()
		i++
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		i++
	}()

	fmt.Println(i)
}

/*
3. 경쟁상태 (채널 활용)
*/
func _27_channel() {
	var i int64
	ch := make(chan int64)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch

	fmt.Println(i)
}
