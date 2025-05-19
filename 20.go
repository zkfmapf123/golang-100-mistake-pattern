package main

import (
	"reflect"
)

type customer struct {
	id string
	op []float64
}

/*
	❌ 값 비교 불가 함
*/
// func BadEquals() {
// 	c1 := customer{id: "x", op: []float64{1.}}
// 	c2 := customer{id: "y", op: []float64{1.}}

// 	return c1 == c2
// }

/*
✅ 값 비교 가능
relfect 이 함수는 기존 == 연산자보다 100배 성능이 느려짐 (배포용 보다는 테스트용으로 해야 함...)
*/
func GoodEquals() bool {
	c1 := customer{id: "x", op: []float64{1.}}
	c2 := customer{id: "x", op: []float64{1.}}

	return reflect.DeepEqual(c1, c2)
}

/*
✅ 성능 좋은 equlas 함수
*/
func (a customer) equals(b customer) bool {
	if a.id != b.id {
		return false
	}

	if len(a.op) != len(b.op) {
		return false
	}

	for i := 0; i < len(a.op); i++ {
		if a.op[i] != b.op[i] {
			return false
		}
	}

	return true
}
