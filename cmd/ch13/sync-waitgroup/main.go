package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// ジョブ数を先に追加
	wg.Add(2)

	go func() {
		fmt.Println("Job 1")
		//　完了を WaitGroup に通知
		wg.Done()
	}()

	go func() {
		fmt.Println("Job 2")
		// 完了を WaitGroup に通知
		wg.Done()
	}()

	// 完了の待ち合わせ
	wg.Wait()
	fmt.Println("終了")
}
