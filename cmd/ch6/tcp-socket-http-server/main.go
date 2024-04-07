package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
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
		go processSession(conn)
	}
}

// 1セッションの処理
func processSession(conn net.Conn) {
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
			Header:     make(http.Header),
		}
		if isGzipAcceptable(req) {
			content := "Hello World (gzipped)\n"
			var buff bytes.Buffer
			writer := gzip.NewWriter(&buff)
			io.WriteString(writer, content)
			writer.Close()
			res.Body = io.NopCloser(&buff)
			res.ContentLength = int64(buff.Len())
			res.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello World\n"
			res.Body = io.NopCloser(strings.NewReader(content))
			res.ContentLength = int64(len(content))
		}
		res.Write(conn)
	}
}

// リクエストヘッダAccept-Encodingにgzipが含まれるか
func isGzipAcceptable(req *http.Request) bool {
	return strings.Contains(strings.Join(req.Header["Accept-Encoding"], ","), "gzip")
}
