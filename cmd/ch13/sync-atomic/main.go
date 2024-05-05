package main

import (
	"fmt"
	"sync/atomic"
)

var id int64

func generateId() int {
	return int(atomic.AddInt64(&id, 1))
}

func main() {
	// ただしこのままでは100生成するというループ処理が完了する前にプログラムが完了してしまう
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId())
		}()
	}
}
