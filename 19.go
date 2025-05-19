package main

import "github.com/inancgumus/prettyslice"

func receiveMessage() []byte {
	var b []byte
	return b
}

func getMessageType(msg []byte) []byte {
	return msg[:5]
}

func storageMessage(b []byte) {}

/*
메모리 누수 ❌
getMesageType 에서 메모리를 100 을 쓴다는 가정
stoargeMessage b 보 메모리를 100으로 쓴다 (msg:5를 하더라도...)
*/
func consumeMessages() {
	for {
		msg := receiveMessage()

		storageMessage(getMessageType(msg))
	}
}

/*
메모리 누수 없음 ✅
num 으로 길이를 지정해서
길이를 지정한 부분까지만 copy ...
*/
func consumeMessageGood(src []byte, num int) {
	dst := make([]byte, num) // num 으로 지정해서 구성된 길이까지 copy

	copy(dst, src)
}

func main() {
	num := 5
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	b := make([]int, num)
	copy(b, a)

	prettyslice.Show("aa", a)
	prettyslice.Show("bb", b)
}
