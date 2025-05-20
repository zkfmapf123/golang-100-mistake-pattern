package main

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
