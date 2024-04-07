package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// see: https://github.com/tkdn/system-programming-with-golang/blob/c4cabe08f00ec3dc8f04b26e35524a171dd4ee99/ch3_io.reader/conn.go#L12
// User-Agent: Go-http-client/1.1
func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	for {
		var err error
		// 接続がまだなければ開始
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POST で文字列を送信
		req, err := http.NewRequest("POST", "http://localhost:8080", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Accept-Encoding", "gzip")
		err = req.Write(conn)
		if err != nil {
			panic(err)
		}
		// サーバから読み込み、タイムアウトはここでエラーとなるのでリトライする
		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer res.Body.Close()

		if res.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(res.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, res.Body)
		}
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
