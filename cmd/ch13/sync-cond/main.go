package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()

			// BroadCast が呼ばれるまで待つ
			cond.Wait()
			fmt.Println(name)
		}()

		fmt.Println("よーい")
		time.Sleep(time.Second)
		fmt.Println("どん")

		cond.Broadcast()
		time.Sleep(time.Second)
	}
}
