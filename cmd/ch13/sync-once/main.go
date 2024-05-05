package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("初期化処理")
}

var once sync.Once

func main() {
	// 呼び出しが複数回あっても一度しかよばれない
	// ただし init 関数が利用されることがほとんど
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
