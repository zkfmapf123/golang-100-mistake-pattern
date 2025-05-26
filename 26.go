package main

import "strings"

/*
❌ 간단한 문자열 연결
간단하디만, 메모리를 다시할당하고 하는만큼 긴 문자열은 문제가 될 수 있음

✅ 대신 간단한 문자열은 아래와같이 구성하는것이 좋음
*/
func _26_simpleConcat() string {
	a := "helloworld"

	s := ""

	for _, v := range a {
		s += string(v)
	}

	return s
}

/*
✅ 문자열 빌러를 사용하면 메모리 할당을 줄일 수 있음
문자열같은 경우 strings.Builder를 활용하자
*/
func _26_longConcat() string {
	a := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}

	s := strings.Builder{}

	for _, v := range a {
		s.WriteString(string(v))
	}

	return s.String()
}
