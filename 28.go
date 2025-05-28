package main

import (
	"fmt"
)

/*
❌ 이러면 경쟁상태가 발생함
*/
func _28_BadRaceCondition() {
	i := 0

	// 경쟁 1 > i ++
	go func() {
		i++ // 1
	}()

	// 경쟁 1 > i 출력
	fmt.Println(i) // 1
}

/*
❌ Buffered Channel은 경쟁상태가 발생함 -> 블로킹이 적어 성능이 더 좋을 수 있음 -> 퍼포먼스가 중요하다면 활용
✅ UnBuffered Channel은 경쟁상태가 발생하지 않는다. -> 매번 블로킹되어 성능이 상대적으로 떨어짐 -> 동기화가 중요할 때, 경쟁상태 방지용
*/
func _28_BadRaceCondition2() {
	i := 0
	ch := make(chan struct{}, 1) // ❌ 이러면 경쟁상태가 발생함
	// ch := make(chan struct{}) // ✅ 이러면 경쟁상태가 발생하지 않음

	go func() {
		i = 1 // 2
		<-ch  // 3
	}()

	ch <- struct{}{} // 1
	fmt.Println(i)   // 2
}

/*
✅ 경쟁상태가 발생하지 않음
*/
func _28_GoodRaceCondition() {
	i := 0
	ch := make(chan struct{})

	go func() {
		<-ch           // 3
		fmt.Println(i) // 4
	}()

	i++              // 1
	ch <- struct{}{} // 2
}
