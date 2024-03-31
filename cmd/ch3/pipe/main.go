package main

import (
	"fmt"
	"io"
	"sync"
)

func main() {
	// パイプの作成
	reader, writer := io.Pipe()

	// ゴルーチン間の同期のためのWaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	// ゴルーチン1: パイプから読み取り、データを処理する
	go func() {
		defer wg.Done()
		defer reader.Close()
		data := make([]byte, 100) // データを読み取るためのバッファを作成
		n, err := reader.Read(data)
		if err != nil {
			fmt.Println("Error reading from pipe:", err)
			return
		}
		fmt.Println("Read data:", string(data[:n]))
	}()

	// ゴルーチン2: データをパイプに書き込む
	go func() {
		defer wg.Done()
		defer writer.Close()
		_, err := writer.Write([]byte("Hello, Pipe!"))
		if err != nil {
			fmt.Println("Error writing to pipe:", err)
			return
		}
	}()

	// ゴルーチンの完了を待機
	wg.Wait()
}
