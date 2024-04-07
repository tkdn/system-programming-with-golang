package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at port 8080.")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			fmt.Printf("Accept: %v\n", conn.RemoteAddr())
			// Accept 後のソケットで何度も応答を返せるようループ
			for {
				// タイムアウト設定
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				req, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時はループを終了
					// それ以外はエラー
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout, 5 seconds has been reached.")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				dump, err := httputil.DumpRequest(req, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				res := http.Response{
					StatusCode: http.StatusOK,
					ProtoMajor: 1,
					ProtoMinor: 1,
					Body:       io.NopCloser(strings.NewReader("Hello World\n")),
				}
				res.Write(conn)
			}
		}()
	}
}
