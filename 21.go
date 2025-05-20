package main

import "fmt"

type Account struct {
	Balance int
}

var a = []Account{
	{Balance: 1000},
	{Balance: 2000},
	{Balance: 3000},
}

/*
❌ 실제 값이 바뀌지 않는다
golang에서의 range는 값을 복제해서 사용한다. (실제값을 바꾸지 않음)
*/
func balanceA() {
	for _, v := range a {
		v.Balance += 1
	}
}

/*
✅ 실제 값을 바꾸려면, index 값을 직접 조정해야 한다.
*/
func balanceB() {
	for i := range a {
		a[i].Balance += 1
	}
}

/*
✅ 굳이, 첫번째 방법을 사용하려면 포인터 형태로 접근하면 가능하다...
하지만 슬라이스의 포인터를 반복한다면 -> CPU 연산 효율이 떨어진다...
*/
func balanceC() {
	// pointer...
}

/*
range는 복제라는 예시 1 (배열은 다름)
*/
func range_ex1() {
	sliceA := []int{1, 2, 3}
	for i, v := range sliceA {
		sliceA[2] = 10

		if i == 2 {
			fmt.Println(v) // 10 slice는 참조방식 -> 값 자체가 바뀐다
		}
	}

	arrayA := [3]int{1, 2, 3}
	for i, v := range arrayA {
		arrayA[2] = 10

		if i == 2 {
			fmt.Println(v) // 2 배열은 복사값의 대해서 변경되지 않는다.
		}
	}
}

/*
✅ 만약, 배열의 대한 값을 바꿔야 한다면? -> 배열 포인터 활용
배열크기가 크다면 아래 방법을 사용하는 것을 추천
배열전체를 굳이 copy 하지 않아도 되기때문에 ...
*/
func arrayBetterCode() {
	a := [3]int{1, 2, 3}

	for i, v := range &a {
		a[2] = 10

		if i == 2 {
			fmt.Println(v)
		}
	}
}
