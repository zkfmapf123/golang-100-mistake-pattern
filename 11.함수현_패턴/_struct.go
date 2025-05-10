package main

import (
	"net/http"
)

/*
요구사항
- 포트를 지정했는가?    -> Default Port 사용
  - 음수인가?           -> 음수 사용 불가 (에러 리턴)
  - 0 인가?            -> 0 사용 불가 (에러 리턴)
  - 지정한 포트 사용
*/
// func NewServer(addr string, port int) (*http.Server, error) {
// 	// ...
// }

/*							첫번째 방법 Struct 활용하는 방법 ✅ 								*/
type Config struct {
	port int // default 0
}

/*
장점 :
1. 포트 지정 여부를 명확하게 표현할 수 있다.
2. Default Port가 0 이다 (int)

단점 :
코드 길이가 늘어날 수 있음 (사용하는 부분)
*/
func NewServerUseConfig(addr string, config Config) (*http.Server, error) {
	// ...
	return nil, nil
}
