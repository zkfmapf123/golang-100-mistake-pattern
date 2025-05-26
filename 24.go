package main

import "context"

/*
❌ 반복문 내에서 잘못된 break 구문
*/
func _badLoop1() {
	for i := 0; i < 10; i++ {
		switch i {
		default:

		case 2:
			break // switch 문만 빠져나옴
		}
	}

	ch := make(chan int)
	for {
		select {
		case <-ch:

		case <-context.Background().Done():
			break // switch 문만 빠져나옴
		}
	}
}

/*
✅ 반복문내에서는 레이블을 사용해서 빠져나올 수 있음
*/
func _goodLoop1() {
loop: // label
	for i := 0; i < 10; i++ {
		switch i {
		default:
		case 2:
			break loop
		}
	}

	ch := make(chan int)

loop_2:
	for {
		select {
		case <-ch:
		case <-context.Background().Done():
			break loop_2
		}
	}
}
