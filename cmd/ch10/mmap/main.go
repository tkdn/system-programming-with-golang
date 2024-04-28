package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/edsrzf/mmap-go"
)

func main() {
	// テストデータを書き込んでおく
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := os.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}
	// memory mapping
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer m.Unmap()

	// メモリに展開したテストデータを修正して書き込み
	m[9] = 'X'
	m.Flush()

	// 読み込む
	fileData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmp: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}
