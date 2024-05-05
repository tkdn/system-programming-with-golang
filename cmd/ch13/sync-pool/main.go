package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d\n", count)
		},
	}

	pool.Put("manually added: 1")
	pool.Put("manually added: 2")
	// sync.Pool は WeakRef のような弱いキャッシュなので
	// GC によってキャシュは消える
	runtime.GC()
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
