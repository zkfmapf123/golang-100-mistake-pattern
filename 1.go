package main

import (
	"fmt"
	"net/http"
)

// 아래와 같은 경우 clinet 변수가 대입이 되지 않는다.
// func HttpTracing(tracing bool) {
// 	var client *http.Client // Compile Error

// 	if tracing {
// 		client, _ := createClient()
// 		fmt.Println("tracing ... ", client)
// 	} else {

// 		client, _ := createClient()
// 		fmt.Println("None Tracing ... ", client)
// 	}
// }

// overwrite...
func GoodHttpTracingPattern(t bool) {
	var client *http.Client // Compile Error

	if t {
		c, _ := createClient()
		fmt.Println("tracing ... ", c)

		// overwrite
		client = c
	} else {

		c, _ := createClient()
		fmt.Println("None Tracing ... ", client)

		// overwrite
		client = c
	}
}

// better pattern
func BetterHttpTracingPattern(t bool) *http.Client {
	if t {
		client, _ := createClient()
		return client
	}

	client, _ := createClient()
	return client
}

func createClient() (*http.Client, error) {
	return &http.Client{}, nil
}
