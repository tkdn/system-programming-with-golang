package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// サイズが1以上の、Signalを受け付けるチャネルを作成
	signals := make(chan os.Signal, 1)
	// 最初のチャネル以降は可変長引数で任意のチャネルを渡せる
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	s := <-signals
	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}
