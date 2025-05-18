package main

type (
	Foo struct{}
)

/*
   15_test.go
   benchmark 결과

   성능 3 > 2 > 1

   즉, 결론

   기본적으로 길이 ,용량 모두 같이 지정하는것이 좋다
   3번처럼 길이를 직접지정하면 성능이 더 빠르긴하나,
   가독성 측면에서는 index 자체가 아닌 append 함수를 사용하는 것이 좋다.
*/

/*
잘못된 슬라이스 지정 방법 1) 초기 할당을 하지 않은 경우
계속 늘어날 여지 존재 ❌
*/
func convert1(foos []Foo) []Foo {
	bars := make([]Foo, 0)

	for _, foo := range foos {
		bars = append(bars, foo) // 0 -> 1 -> 2 -> 4 -> 8 -> 12 -> 16 ...
	}

	return bars
}

/*
슬라이스 지정방법 2) capa 에다 넣기 (append)
성능이 중간이지만 golang에서는 이 방법을 선호한다. ✅

이유는...
capa가 2배씩 증가한다고 해서, 성능에 크기 영향을 미치지않고,
index 자체를 사용해서 controlle 하는것이 가독성 및 사용성면에서 크게 효율성이 떨어진다.
*/
func convert2(foos []Foo) []Foo {
	l := len(foos)
	bars := make([]Foo, 0, l)

	for _, foo := range foos {
		bars = append(bars, foo)
	}

	return bars
}

/*
슬라이스 지정방법 3) len 에다 직접 넣기 [i] 인덱스 직접 지정
성능이 제일 빠르나 golang에서는 이와같은 방법을 사용하지는 않는다. ❌
*/
func convert3(foos []Foo) []Foo {
	l := len(foos)
	bars := make([]Foo, l) // len을 직접 지정

	for i, foo := range foos {
		bars[i] = foo
	}

	return bars
}
