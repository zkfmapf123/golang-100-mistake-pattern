package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Context : Deadline, 취소 시그널, API 경계의 대한 여러값을 가지고 있음

	context.Background
	- 빈 컨텍스트

	context.Todo (왠만하면 얘를 쓰자...)
	- 아직 정의되지 않은 컨텍스트
*/

/*
deadline
몇초 , 몇일 동안만 유지되는 고루틴을 생성함 (WithTimeout)
*/
func tempGenerateDeadline() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second) // 4초후 고루틴 없어짐
	defer cancel()

	fmt.Println(ctx.Err()) // nil

	time.Sleep(5 * time.Second)
	fmt.Println(ctx.Err()) // context canceled
	return ctx
}

/*
context cancel
cancel 함수를 호출하면 context 가 취소됨
*/
func contextCancelEx() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(ctx.Err())
	cancel()
	fmt.Println(ctx.Err())
}

/*
취소 감지 예제
*/
func handler(ctx context.Context, ch chan string) error {
	for {
		select {
		case msg, ok := <-ch:
			// channel close
			if !ok {
				fmt.Println("채널이 닫혔습니다")
				return nil
			}

			fmt.Println("채널 수신 >> ", msg)
		case <-ctx.Done():
			fmt.Println("Done...", ctx.Err())
			return ctx.Err()
		}
	}
}
