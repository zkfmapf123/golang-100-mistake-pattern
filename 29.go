package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	time.Sleep(time.Second * 2)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("123123")
		}
	}()
}
