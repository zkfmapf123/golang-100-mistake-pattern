package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	- Mutex 활용
	- Channel 활용
	- atomic 활용 -> 오버헤드가 클 수 있음
*/

/*
데이터 경쟁
- 두개이상의 고루틴이 동시에 같은 메모리 위치에 접근
- 그 중 하나 이상이 쓰기 작업을 할 경우 발생
*/
func state_data_race() {
	var counter int

	go func() {
		counter++ // 쓰기 작업
	}()

	go func() {
		counter++ // 쓰기 작업
	}()

	time.Sleep(time.Second * 2)
	fmt.Println(counter)
}

/*
경쟁상태
- 프로그램의 실행결과가 고루틴들의 실행순서의 의존하는 상황
*/
func state_race() {
	var value int
	var ready bool

	go func() {
		value = 100
		ready = true
	}()

	go func() {
		for !ready {
			// ready 가 true 가 될때까지 대기
		}

		fmt.Println(value)
	}()

	time.Sleep(time.Second * 2)
}

// ❌ 잘못된 뮤텍스 사용
func bad_mutex() {
	var mu sync.Mutex
	go func() {
		mu.Lock()
		// 작업 수행
		// mu.Unlock() 호출 누락
	}()
}

// ❌ 잘못된 채널 사용
func bad_channel() {
	ch := make(chan struct{}, 1)
	go func() {
		ch <- struct{}{} // 버퍼드 채널로 인한 경쟁 상태
	}()
}
