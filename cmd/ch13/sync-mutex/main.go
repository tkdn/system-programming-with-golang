package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	return result
}

func main() {
	// mutex := new(sync.Mutex) と同義
	var mutex sync.Mutex

	// ただしこのままでは100生成するというループ処理が完了する前にプログラムが完了してしまう
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
