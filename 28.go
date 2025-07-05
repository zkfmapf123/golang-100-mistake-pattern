package main

import (
	"context"
	"fmt"
	"time"
)

/*
time.After : 해당 채널에 5초동안 아무런 메시지가 오지 않는다면 ... 한다

time.After 자체가 메모리를 많이 먹는건 아니지만 -> 결국 OOM 으로 발생할 여지가 존재한다

❌ 아래처럼 코드를 구성하면 메모리상의 문제가 발생할 여지가 존재한다. -> time.After 이후 채널을 닫아주지 않는다
1.15의 경우, time.After은 200Byte 가량 사용한다

	-> 2초 당 200Byte
	-> 1분 당, (2 * 30) = 200 * 30 = 6000KB = 6MB
	-> 1시간 당, (1 * 60) = 6MB * 60 = 360 MB
	-> 24시간 당, (1 * 24) = 360MB * 24 = ...8 GB
*/
func badTimeAfterTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		ch <- "hello world"
	}()

	defer close(ch)

	for {
		select {
		case v := <-ch:
			fmt.Println(v)

		// ❌ 2초 동안 없다면 -> 메모리 누수 발생
		case <-time.After(time.Second * 2):
			fmt.Println("timeout")
		}
	}
}

/*
✅ 이 방법은 좋지만...
for loop 를 돌면서, 컨텍스트를 다시 재생성한다
golang 언어에서는 Context를 다시 생성하는 건 꽤 크다
*/
func goodTimeAfterTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 10)
		ch <- "hello world"
	}()

	defer close(ch)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

		select {
		case v := <-ch:
			cancel()
			fmt.Println(v)

		// ✅  2초 동안 없다면...
		case <-ctx.Done():
			fmt.Println("Timeout...")
		}
	}
}

func betterTimeAfterTimeout() {
	ch := make(chan string)
	timeDuration := time.Second * 2
	timer := time.NewTimer(timeDuration)

	go func() {
		time.Sleep(time.Second * 10)
		ch <- "hello world"
	}()

	defer close(ch)

	for {
		timer.Reset(timeDuration) // loop 돌고 다시 timeDuration 만큼 리셋

		select {
		case v := <-ch:
			fmt.Println(v)

		// ✅  2초 동안 없다면...
		case <-timer.C:
			fmt.Println("Timeout...")
		}
	}
}
