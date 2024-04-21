package main

import "os"

func main() {
	// os.Mkdir("setting", 0755)
	os.MkdirAll(".temp/setting", 0755)

	os.Truncate(".temp/server.log", 100)
	// os.File 構造体を使って file.Tuncate する方法もある
	// file, _ := os.Open(".temp/server.log")
	// file.Truncate(100)
	// file.Close()

	file, _ := os.Create(".temp/org.txt")
	file.WriteString("origin")
	file.Close()
	os.Rename(".temp/org.txt", ".temp/copy.txt")
	os.Mkdir(".temp/foo", 0755)
	os.Rename(".temp/copy.txt", ".temp/foo/copy.txt")

	// 違うデバイスに移動
	// 書籍ではエラーがでる想定のようだがエラーは発生しなかった
	if err := os.Rename(".temp/foo/copy.txt", "/tmp/copy.txt"); err != nil {
		panic(err)
	}
}
